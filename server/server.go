
package main

import (
	"myrepo/PufferBlock/server/action"
	//"myrepo/PufferBlock/server/websockets"
)

//主程序入口
func main() {
	//初始化网络
	//	action.Init()

	//建立连接，接受请求并回复
	//websockets.Websockets()
	action.InitCC("a")
	action.QueryCC("queryByOwner","a")
	action.InitCC("b")
	action.QueryCC("queryByOwner","b")
	//action.InitCC("q")
	//action.QueryCC("queryByOwner","o")
	//action.QueryCC("queryByOwner","p")
	action.InvokeCC("o","p",10)
	//action.QueryCC("queryByOwner","o")
	action.QueryCC("queryAllCarbonInfo","all")
}
