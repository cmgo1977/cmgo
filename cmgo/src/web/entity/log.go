package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//log
type ErrorLog struct {
	Uid        bson.ObjectId `json:"uid"           bson:"uid"`        //用户id
	NickName   string        `json:"nickname"      bson:"nickname"`   //昵称
	Mobile     string        `json:"mobile"        bson:"mobile"`     //手机
	Passwd     string        `json:"passwd" 	   bson:"passwd"`     //密码
	CreateDate time.Time     `json:"createdate"    bson:"createdate"` //注册日期
}