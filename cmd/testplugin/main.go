package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/greboid/irc/rpc"
	"github.com/kouhin/envflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"time"
)

var (
	RPCHost  = flag.String("rpc-host", "localhost", "gRPC server to connect to")
	RPCPort  = flag.Int("rpc-port", 8001, "gRPC server port")
	RPCToken = flag.String("rpc-token", "", "gRPC authentication token")
	Channel  = flag.String("channel", "", "Channel to send messages to")
)

func main() {
	if err := envflag.Parse(); err != nil {
		log.Fatalf("Unable to load config: %s", err.Error())
	}
	time.Sleep(5 * time.Second)
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *RPCHost, *RPCPort), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to fake plugin: %v", err)
	}
	defer func() { _ = conn.Close() }()
	client := rpc.NewIRCPluginClient(conn)
	_, err = client.Ping(rpc.CtxWithToken(context.Background(), "bearer", *RPCToken), &rpc.Empty{})
	messages := []string{"Test message."}
	for message := range messages {
		_ = sendMessage(messages[message], client)
	}
	_ = getMessages(*Channel, client)
}

func sendMessage(message string, client rpc.IRCPluginClient) error {
	_, err := client.SendChannelMessage(rpc.CtxWithToken(context.Background(), "bearer", *RPCToken), &rpc.ChannelMessage{
		Channel: *Channel,
		Message: message,
	})
	if err != nil {
		log.Printf("Error sending: %s", err.Error())
	} else {
		log.Print("Sent message.")
	}
	return err
}

func getMessages(channel string, client rpc.IRCPluginClient) error {
	handler, err := client.GetMessages(
		rpc.CtxWithToken(context.Background(), "bearer", *RPCToken),
		&rpc.Channel{Name: channel},
	)
	if err != nil {
		log.Fatalf("Error getting messages: %s", err.Error())
	}
	for {
		msg, err := handler.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Err receiving message: %v", err)
		}
		log.Printf("%s: %s - %s", msg.Channel, msg.Source, msg.Message)
	}
	return err
}
