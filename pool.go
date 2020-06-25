package main

import (
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
	conn.Write([]byte("+OK\r\n"))
	return struct{}{}
}

func init() {
	var numCPUs = runtime.NumCPU()
	pool = tunny.NewFunc(numCPUs, f)
}
