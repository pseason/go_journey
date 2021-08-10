package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Hight int
}

type Rect struct {
}

// 矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Hight
	return nil
}

// 矩形周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Hight) * 2
	return nil
}

func main() {
	// 1.注册一个rect的服务
	rect := new(Rect)
	rpc.Register(rect)
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe(":8998", nil)
	if err != nil {
		log.Panicln(err)
	}
}
