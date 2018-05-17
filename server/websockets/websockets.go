//Package websockets 服务器侧websocket服务
package websockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

//Request 请求结构
type Request struct {
	Type     string `json:"requesttype"`
	Name     string `json:"username"`
	Function string `json:"command"`
	OpName   string `json:"operatname"`
	OpAmount int    `json:"operatamount"`
}

//Websockets websocket实现
func Websockets() {

	http.Handle("/", websocket.Handler(echo))
	if err := http.ListenAndServe(":6001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

//接收请求并回复
func echo(ws *websocket.Conn) {
	for {
		requestAsBytes := []byte{}
		if err := websocket.JSON.Receive(ws, &requestAsBytes); err != nil {
			fmt.Println("Can't receive")
			break
		}

		request := &Request{}
		json.Unmarshal(requestAsBytes, request)
		fmt.Println("Received back from client: " + string(requestAsBytes))

		//分类解析请求，做出回复
		response, err := request.doSelect()
		if err != nil {
			fmt.Println("Can't understand requst")
			break
		}

		responseAsBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Can't understand response")
			break
		}
		fmt.Println("Sending to client: " + string(responseAsBytes))
		if err := websocket.JSON.Send(ws, responseAsBytes); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
