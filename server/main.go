package main

import (
	"go-messenger/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":8080"
)

type mychatservice struct {
	chat.UnimplementedChatserviceServer
}

func main() {
	port, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Error listening on port %s", PORT)
	}
	serverRegisterd := grpc.NewServer()
	services := &mychatservice{}
	chat.RegisterChatserviceServer(serverRegisterd, services)
	err = serverRegisterd.Serve(port)
	if err != nil {
		log.Fatalf("Error serving grpc server %v", err)
	}

}
