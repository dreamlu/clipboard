package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Clip(u *gin.Context) {

	go handleMessages()

	ws, err := upgrader.Upgrade(u.Writer, u.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	WsHander(ws)
}
