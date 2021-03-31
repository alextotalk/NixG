package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server2() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9488")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
		fmt.Println("conn")
	}
}
func client2() {
	c, err := rpc.Dial("tcp", ":9488")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}
func RPC() {
	go server2()
	go client2()

	var input string
	fmt.Scanln(&input)
}
