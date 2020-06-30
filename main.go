package main

import (
	"fmt"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/gin-gonic/gin"
	"github/dreamlu/clipboard/clipboard"
	"github/dreamlu/clipboard/routers"
)

func main() {
	res, err := clipboard.Read()
	if err != nil {
		fmt.Println("read err:", err)
		return
	}
	//fmt.Println("clip content:", string(res))

	// websocket
	go http()
	// local read
	go clipboard.LocalRead()

	// window
	clip := app.New()

	w := clip.NewWindow("clip")
	w.SetContent(widget.NewVBox(
		widget.NewLabel(string(res)),
		//widget.NewButton("Quit", func() {
		//	//clip.Quit()
		//}),
	))

	w.ShowAndRun()
}

func http() {
	//log.Println(gt.Version)
	gin.SetMode(gin.DebugMode)
	//r := routers.SetRouter()
	// pprof.Register(r)
	// Listen and Server in 0.0.0.0:8080
	_ = routers.Router.Run(":9001")
}