package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	addr := ":6801"
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
			//go procText(num, conn)
			go procBin(num, conn)
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

func procBin(num int, conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096)
	var err error
	for n := 0; err == nil; {
		n, err = bufio.NewReader(conn).Read(buf)

		v := strings.Split(string(buf[:n]), ";")
		if len(v) < 10 {
			log.Println("incomplete", len(v))
		} else {
			var sc SensorCoords
			t := time.Now()
			for _, str := range v {
				if "" != str {
					err2 := sc.FromString(t, str)
					if nil == err2 {
						log.Println(num, t.Format("2006-01-02 15:04:05.000"), sc.X, sc.Y, sc.Z)
					} else {
						log.Println(num, "convert error")
					}
					t = t.Add(time.Millisecond * 100)
				}
			}
		}

		//log.Println(num, n, "bytes:", buf[:n])
		//log.Println(num, string(buf[:n]))
	}
	log.Println(num, "Connection closed")
}
