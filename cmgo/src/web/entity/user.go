package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//pg
type User_pg struct {
	Uid        int64     `xorm:"pk autoincr 'id'" json:"uid"`                                     //用户id
	NickName   string    `xorm:"varchar(40)" json:"nickname"`                                     //昵称
	Mobile     string    `xorm:"char(11)" json:"nickname"`                                        //手机
	Passwd     string    `xorm:"varchar(30)" json:"-"`                                            //密码
	CreateDate time.Time `xorm:"createdate" json:"createdate"  time_format:"2006-01-02 15:04:05"` //注册日期
}

//mg
type User_mg struct {
	Uid        bson.ObjectId `json:"uid"           bson:"uid"`        //用户id
	NickName   string        `json:"nickname"      bson:"nickname"`   //昵称
	Mobile     string        `json:"mobile"        bson:"mobile"`     //手机
	Passwd     string        `json:"passwd" 	   bson:"passwd"`     //密码
	CreateDate time.Time     `json:"createdate"    bson:"createdate"` //注册日期
}
