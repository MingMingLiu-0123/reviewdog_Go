package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

//调用服务
func main() {
	//链接远程RPC服务
	rp, err := rpc.DialHTTP("tpc", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	//调用与远程方法
	//定义接受服务端传回来的计算结果的变量

	ret := 0
	//求面积
	err2 := rp.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)

	//求周长
	err3 := rp.Call("Rect.Area", Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("面积：", ret)
}
