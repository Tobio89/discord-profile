package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const rpcPort = "5001"

type Config struct{}

func main() {

	app := Config{}

	err := rpc.Register(new(RPCServer))
	if err != nil {
		log.Panicln("Error registering RPC server: ", err)
	}

	go app.rpcListen()

	select {}

}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port: ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}

	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			log.Println("Error accepting RPC connection: ", err)
			continue
		}

		go rpc.ServeConn(rpcConn)
	}

}
