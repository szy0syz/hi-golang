package main

import "net/rpc"

type Server struct {

}

type Req struct {
	A int
	B int
}

type Res struct {
	Sum
}

// 给Server结构体上绑定一个方法
func (s *Server ) Add(req Req, res, Res)error {
	res.Sum = req.A + req.B
	return nil
}

func main() {
	rpc.Register(new(Server))
}
