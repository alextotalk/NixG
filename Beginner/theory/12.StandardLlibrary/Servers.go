package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// слушать порт
	ln, err := net.Listen("tcp", ":9992")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// принятие соединения
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// обработка соединения
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// получение сообщения
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	// соединиться с сервером
	c, err := net.Dial("tcp", "127.0.0.1:9992")
	if err != nil {
		fmt.Println(err)
		return
	}

	// послать сообщение
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func testServer() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
