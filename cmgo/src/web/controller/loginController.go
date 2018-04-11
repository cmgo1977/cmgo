package controller

import (
	"time"

	"web/entity"
	"web/handler"
	"web/model"
)

type Login struct{}

func (Login) UserLoginHandle() string {
	login := handler.UserLogin{
		entity.User_pg{
			Uid:        1,
			NickName:   "caimin",
			Mobile:     "13162578783",
			Passwd:     "123qwe",
			CreateDate: time.Now(),
		},
	}
	var user model.Ilogin
	user = login

	return user.Login()
}
