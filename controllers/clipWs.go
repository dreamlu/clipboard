package controllers

import (
	"github.com/gorilla/websocket"
	"github/dreamlu/clipboard/clipboard"
	"log"
	"net/http"
)

//客户端
type Client struct {
	//GroupID string // 标识客户端
	//UID     int64  // 唯一标识用户id
	Conn *websocket.Conn
}

var (
	clients []*Client //客户端队列,指针同步同一个client data
	//broadcast   = make(chan []byte) // 广播的数据
	//ClipContent = make(chan []byte) // clipboard content/广播的数据
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 消息读取
// 开启不同进程代表对应的客户端通信
func WsHander(ws *websocket.Conn) {

	defer ws.Close()

	// Register our new client
	//注册客户端连接
	var (
		ct Client
	)
	ct.Conn = ws
	//放入连接队列
	clients = append(clients, &ct)

	//消息读取,每个客户端数据
	for {
		var req []byte
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Printf("[错误-read]: %v", err)
			continue
		}

		// @ping: a ping
		if string(data) == "@ping" {
			log.Printf("[心跳检测]: %v", string(data))
			continue
		}
		//data = data[1 : len(data)-1]
		log.Printf("[消息内容]: %v", string(data))

		req = data
		// 同剪切板数据相同continue
		if string(req) == string(clipboard.Clip) {
			continue
		}

		// write to clipboard
		err = clipboard.Write(data)
		if err != nil {
			log.Println("ws write err:", err.Error())
		}

		//clipboard.ClipContent <- req
		if err != nil {
			log.Printf("[错误-read2]: %v", err)
			//delete(clients, ws) //删除对应连接
			for k, v := range clients { //删除对应连接,emm...暂时先遍历删除～
				v.Conn.Close()
				// 删除失效连接
				clients = append(clients[:k], clients[k+1:]...)
			}
			//break
		}
	}
}

// 消息写入
// 消息推送(不通进程代表各自客户端的写入进程)
func handleMessages() {
	for {
		msg := <-clipboard.ClipContent // 阻塞等待/广播
		//获得广播内容发送给所有端
		for k, client := range clients {
			// send message to every specified client, hehe~
			//if client.GroupID != msg.GroupId { // must same group_id
			//	continue
			//}
			// next have same group_id
			err := client.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil { //当连接挂了
				//fmt.Println("客户:",client,"聊天记录写入失败")
				log.Printf("[错误-write]: %v", err)
				client.Conn.Close()
				// 删除失效连接
				clients = append(clients[:k], clients[k+1:]...)
				////记录该用户最后读的消息id,广播中处理,待优化
				continue
			}
		}
	}
}
