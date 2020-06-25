package server

import (
	"context"
	"net"
)

var server net.Listener

const (
	protocol = "tcp"
	bind     = ":6379"
)

func StartRedmeServer(ctx context.Context, chanConn chan<- net.Conn) {
	server, err := net.Listen(protocol, bind)
	if err != nil {
	}
	serve(server, chanConn)
}

func serve(server net.Listener, chanConn chan<- net.Conn) {
	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		chanConn <- conn
	}
}
