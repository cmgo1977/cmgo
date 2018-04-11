package controller

import (
	"log"

	"web/handler"
)

type Register struct{}

func (Register) UserRegisterHandle() (uid string) {

	uid, err := handler.Register{}.Register("滴滴打车", "13162578783", "123")
	if err != nil {
		log.Fatalf("Register时报错: %s\n", err)
		return ""
	}

	return uid
}
