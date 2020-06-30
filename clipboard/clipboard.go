package clipboard

import (
	"log"
	"time"
)

var (
	Clip        []byte              // 上一次剪切板内容
	ClipContent = make(chan []byte) // clipboard content/广播的数据
)

// load local clipboard
func LocalRead() {
	for {
		time.Sleep(1 * time.Second)
		b, err := Read()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if string(Clip) == string(b) {
			continue
		}
		Clip = b
		ClipContent <- b
	}
}
