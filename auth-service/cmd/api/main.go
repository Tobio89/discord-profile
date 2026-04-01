package main

import (
	"database/sql"
	"discord-profile/auth-service/data"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const rpcPort = "5001"
const connectionRetryDelay = 2 * time.Second
const connectionRetryLimit = 10

var connectionRetryCount int64

type Config struct {
	Repo data.Repository
}

var app Config

func main() {

	app = Config{}

	dbConnection := connectToDB()
	if dbConnection == nil {
		log.Panicln("Could not connect to database")
	}

	app.setupRepo(dbConnection)

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

// Opens connection to DB
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Handles retrying DB connection
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	// TODO: handle empty dsn env case

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready")
			connectionRetryCount++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		if connectionRetryCount > connectionRetryLimit {
			log.Println(err)
			return nil
		}

		log.Println("Waiting to retry...")
		time.Sleep(connectionRetryDelay)
		continue
	}

}

func (app *Config) setupRepo(conn *sql.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
