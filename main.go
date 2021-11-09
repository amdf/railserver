package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	addrlocal = ":12321" //192.168.10.126
	addrvps   = ":6801"
)

func main() {

	addr := addrvps
	// Устанавливаем прослушивание порта

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Launching server at", addr)

	num := 0
	for {
		conn, err := ln.Accept()

		if err == nil {
			num++
			log.Println(num, "Connection opened")
			go procText(num, conn)
		}
	}
}

func procText(num int, conn net.Conn) {
	defer conn.Close()
	str := ""
	var err error
	var sc SensorCoords
	for err == nil {
		str, err = bufio.NewReader(conn).ReadString(';')
		err2 := sc.FromString(time.Now(), str)
		if nil == err2 {
			log.Println(num, sc.X, sc.Y, sc.Z)
		} else {
			log.Println(num, "convert error")
		}
	}
	log.Println(num, "Connection closed")
}

func procBinary(num int, conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	var err error
	for n := 0; err == nil; {
		n, err = bufio.NewReader(conn).Read(buf)

		log.Println(num, buf[:n])
		//log.Println(num, string(buf[:n]))

		// Отправить новую строку обратно клиенту
		//conn.Write([]byte(newmessage + "\n"))
	}
	log.Println(num, "Connection closed")
}
