package main

import (
	"flag"
	"github.com/greboid/irc/irc"
	"github.com/greboid/irc/logger"
	"github.com/greboid/irc/rpc"
	"github.com/kouhin/envflag"
	"go.uber.org/zap"
)

//go:generate protoc -I ../../rpc plugin.proto --go_out=plugins=grpc:../../rpc

var (
	Server        = flag.String("server", "", "Which IRC server to connect to")
	Password      = flag.String("password", "", "The server password, if required")
	TLS           = flag.Bool("tls", true, "Connect with TLS?")
	Nickname      = flag.String("nick", "", "Nickname to use")
	Realname      = flag.String("realname", "", "'Real name' to use")
	Channel       = flag.String("channel", "", "Channel to join on connect")
	Debug         = flag.Bool("debug", false, "Enable IRC debug output")
	SASLAuth      = flag.Bool("sasl-auth", false, "Authenticate via SASL?")
	SASLUser      = flag.String("sasl-user", "", "SASL username")
	SASLPass      = flag.String("sasl-pass", "", "SASL password")
	RPCPort       = flag.Int("rpc-port", 8001, "gRPC server port")
	PluginsString = flag.String("plugins", "", "Comma separated list of plugins, name=token")
	FloodProfile  = flag.String("flood-profile", "restrictive", "Flood profile: restrictive, unlimited")
	WebPort       = flag.Int("web-port", 8000, "Web port for http server")
)

func main() {
	log := logger.CreateLogger(*Debug)
	defer func() {
		err := log.Sync()
		if err != nil {
			panic("Unable to sync logs")
		}
	}()
	log.Info("Starting bot")
	if err := envflag.Parse(); err != nil {
		log.Fatal("Unable to load config.", zap.String("error", err.Error()))
	}
	Plugins, err := rpc.ParsePluginString(*PluginsString)
	if err != nil {
		log.Fatal("Unable to load config.", zap.String("error", err.Error()))
	}
	eventManager := irc.NewEventManager()
	connection := irc.NewIRC(*Server, *Password, *Nickname, *Realname, *TLS, *SASLAuth, *SASLUser, *SASLPass, log,
		*FloodProfile, eventManager)
	rpcServer := rpc.NewGrpcServer(connection, eventManager, *RPCPort, Plugins, *WebPort, log)
	log.Info("Adding callbacks")
	addBotCallbacks(connection)
	go rpcServer.StartGRPC()
	err = connection.ConnectAndWaitWithRetry(5)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Exiting")
}
