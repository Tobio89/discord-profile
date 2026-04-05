package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const rpcPort = "5001"
const httpPort = "4455"

type Config struct{}

func main() {

	app := Config{}

	err := rpc.Register(new(RPCServer))
	if err != nil {
		log.Panicln("Error registering RPC server: ", err)
	}

	go app.rpcListen()

	log.Printf("profile-broker listening for HTTP on %s\n", httpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

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
