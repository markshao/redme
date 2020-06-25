package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"

	"github.com/Jeffail/tunny"
)

var pool *tunny.Pool
var f = func(input interface{}) interface{} {
	conn, err := input.(net.Conn)
	if !err {
		log.Println("%w", err)
	}
	defer conn.Close()
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		log.Println("got", n, "bytes.")
		buf = append(buf, tmp[:n]...)
		log.Println(string(buf))

	}
	log.Println(buf)
	conn.Write(buf)
	return struct{}{}
}

func init() {
	var numCPUs = runtime.NumCPU()
	pool = tunny.NewFunc(numCPUs, f)
}
