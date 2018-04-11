package handler

import (
	"bytes"

	"web/entity"
)

type UserLogin struct {
	User entity.User_pg
}

type AdminLogin struct {
	User entity.User_pg
}

type TeacherLogin struct {
	User entity.User_pg
}

func (u UserLogin) Login() string {

	var logininfo bytes.Buffer

	logininfo.WriteString(u.User.Mobile)
	logininfo.WriteString(u.User.Passwd)

	return logininfo.String()
}

func (u AdminLogin) Login() string {

	var logininfo bytes.Buffer

	logininfo.WriteString(u.User.Mobile)
	logininfo.WriteString(u.User.Passwd)

	return logininfo.String()
}

func (u TeacherLogin) Login() string {

	var logininfo bytes.Buffer

	logininfo.WriteString(u.User.Mobile)
	logininfo.WriteString(u.User.Passwd)

	return logininfo.String()
}
