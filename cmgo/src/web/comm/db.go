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

/*
	Redis操作相关
*/

//测试struct
type Myuser struct {
	Name  string
	Phone string
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

/*
	设置key的过期时间(秒)
	result1, err := redisConnect.Do("EXPIRE", key, 30)
	if err != nil {
	log.Fatalf("xxx 报错: %s\n", err)
	return
*/

//插入key，value
func RedisSetKey(key, value string) (result string) {
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	//插入，有过期时间
	result, err := redis.String(redisConnect.Do("SET", key, value, "EX", C.Db.Redis.R_expire))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return ""
	}
	//ok = v.(string) //将空接口转换为string

	/*插入，不过期
	result, err := redis.String(redisConnect.Do("SET", key, value))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return ""
	}
	*/

	return result //要判断是否返回ok
}

//根据key查询value
func RedisGetKey(key string) (result string) {
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	//redis操作
	result, err := redis.String(redisConnect.Do("GET", key))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return
	}

	return result
}

//累加（每执行一次加1,返回结果数）
func RedisAccumulation(key string) (result int64) {
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	//redis操作
	result, err := redis.Int64(redisConnect.Do("INCR", key))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return
	}

	return result
}

//判断某个key是否存在
func RedisExitKey(key string)(result bool){
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	//redis操作
	result, err := redis.Bool(redisConnect.Do("EXISTS", key))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return
	}

	return result
}

//删除key
func RedisDeleteKey(key string)(result bool){
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	//redis操作
	result, err := redis.Bool(redisConnect.Do("DEL", key))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return
	}

	return result
}

//插入json
//key := "profile"
//_map := map[string]string{"username": "666", "phonenumber": "888"}
//value, _ := json.Marshal(_map)
func RedisSetJson(key string,value string)(result int64){
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	//redis操作
	result, err := redis.Int64(redisConnect.Do("SETNX", key,value))
	if err != nil {
		log.Fatalf("Redis > SetKey 报错: %s\n", err)
		return
	}

	return result
}

//批量插入Map
func RedisSetMap() {
	redisPool = InitRedisPool()
	redisConnect := redisPool.Get() //从连接池获取连接
	defer redisConnect.Close()      //用完后放回连接池

	user := map[string]Myuser{
		"caimin": Myuser{Name: "caimin", Phone: "13162578783"},
		"lirui":  Myuser{Name: "lirui", Phone: "18234545454"},
	}

	//保存Map
	for sym, row := range user {
		if _, err := redisConnect.Do("HMSET", redis.Args{sym}.AddFlat(row)...); err != nil {
			log.Fatal(err)
		}
	}
}

/*
	Mongodb操作相关
*/
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
