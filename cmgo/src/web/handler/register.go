package handler

import (
	"log"
	"time"

	"web/comm"
	"web/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Register struct{}

//注册
func (Register) Register(nickName, mobile, passwd string) (uid string, err error) {

	user := entity.User_mg{
		Uid:        bson.NewObjectId(),
		NickName:   nickName,
		Mobile:     mobile,
		Passwd:     passwd,
		CreateDate: time.Now(),
	}

	sqlHandle := func(c *mgo.Collection) error {
		return c.Insert(user)
	}

	err = comm.MongodbGetCollection("user", sqlHandle)
	if err != nil {
		log.Fatalf("Register时报错: %s\n", err)
		return "", err
	}

	return user.Uid.Hex(), err
}
