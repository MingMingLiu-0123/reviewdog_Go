package main

import (
	"log"
	"net/http"
	"net/rpc"
)

//服务端。求矩形面积和周长

//声明矩形对象
type Rect struct {
}

//声明参数结构体，字段首字母大写
type Params struct {
	Width, Height int
}

//             接收的参数 ，返回给客户端的参数
func (r *Rect) Area(p Params, ret *int) error {

	*ret = p.Height * p.Width
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {

	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	//注册服务
	rect := new(Rect)
	rpc.Register(rect)
	//吧服务绑定到HTTP协议上
	rpc.HandleHTTP()
	//监听服务，等待客户端调用求面积和周长的方法
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		if err != nil {
			log.Fatal(err) //打印err信息
		}
	}
}
