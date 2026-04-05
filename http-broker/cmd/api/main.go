package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpPort = "4455"

type Config struct{}

func main() {

	app := Config{}

	log.Printf("http-broker listening for HTTP on %s\n", httpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
