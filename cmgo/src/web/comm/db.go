// Copyright (c) 2018 数据库操作
package comm

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
)

var (
	mongodbSession *mgo.Session
	mongodbDb      *mgo.Database

	redisPool *redis.Pool
)

type RedisPool struct {
	Dial         func() (redis.Conn, error)            //Dial 是创建链接的方法
	TestOnBorrow func(c redis.Conn, t time.Time) error //TestOnBorrow 是一个测试链接可用性的方法
	MaxIdle      int                                   //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
	MaxActive    int                                   //最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
	IdleTimeout  time.Duration                         //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭,应该设置一个比redis服务端超时时间更短的时间
	Wait         bool                                  //当链接数达到最大后是否阻塞，如果不的话，达到最大后返回错误,如果Wait被设置成true，则Get()方法将会阻塞
}

//初始化redis连接池
func InitRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   1024,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", C.Db.Redis.R_host)
			if err != nil {
				log.Fatalf("redigo->RedigoPool->redis.Dial()初始化连接池时报错: %s\n", err)
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")

			return err
		},
	}
}

func RedisSetKey(key, value string) (ok string) {
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()   //用完后放回连接池

	//redis操作
	v, err := redisConnect.Do("SET", key, value)
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return ""
	}
	ok = v.(string) //将空接口转换为string

	return ok //要判断是否返回ok
}

func MongodbGetSession() *mgo.Session {
	if mongodbSession == nil {
		var err error
		//session不存在，就新创建
		mgoDialInfo := &mgo.DialInfo{
			Addrs:     []string{C.Db.Mongodb.M_host},
			Direct:    C.Db.Mongodb.M_direct,
			Timeout:   30 * time.Second,
			PoolLimit: C.Db.Mongodb.M_poollimit,
			//Username:  C.Db.Mongodb.M_user,
			//Password:  C.Db.Mongodb.M_passwd,
		}
		mongodbSession, err = mgo.DialWithInfo(mgoDialInfo) //创建一个维护套接字池的session
		if err != nil {
			log.Fatalf("getSession-mgo.DislwithInfo()时报错: %s\n", err)
		}
		/*
			Strong: session 的读写一直向主服务器发起并使用一个唯一的连接，因此所有的读写操作完全的一致。
			Monotonic: session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），只要出现了一次写操作，session 的连接就会切换至主服务器。由此可见此模式下，能够分散一些读操作到其他服务器，但是读操作不一定能够获得最新的数据。
			Eventual: session 的读操作会向任意的其他服务器发起，多次读操作并不一定使用相同的连接，也就是读操作不一定有序。session 的写操作总是向主服务器发起，但是可能使用不同的连接，也就是写操作也不一定有序。
		*/
		mongodbSession.SetMode(mgo.Monotonic, true)
		mongodbDb = mongodbSession.DB(C.Db.Mongodb.M_name) //使用指定数据库
	}

	return mongodbSession.Clone()
}

func MongodbGetCollection(c string, sqlHandle func(*mgo.Collection) error) error {
	session := MongodbGetSession()
	defer session.Close()
	collection := mongodbDb.C(c) //拿到指定集合

	return sqlHandle(collection)
}
