package main

import (
	"context"
	"go-messenger/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	PORT = ":8080"
)

func main() {
	conn, err := grpc.NewClient(PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := chat.NewChatserviceClient(conn)

	sendmessage(client)
}

func sendmessage(client chat.ChatserviceClient) {
	req := &chat.SendMessageRequest{
		Sender:    "Alice",
		Receiver:  "Bob",
		Message:   "Hello Bob!",
		Timestamp: time.Now().Unix(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := client.SenderMessage(ctx, req)
	if err != nil {
		log.Fatalf("unable to send message %v", err)
	}
	log.Println("Response from the server recorded %v\n", resp)
}
