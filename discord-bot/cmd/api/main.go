package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("The app")

	bot := createSession("")

	startSession(bot)

	// Create channel, hold it open
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	log.Println("discord bot shutting down")
	bot.Close()
}

func createSession(token string) *discordgo.Session {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Panic(err)
	}

	return dg
}
func startSession(dg *discordgo.Session) {
	err := dg.Open()
	if err != nil {
		log.Panic(err)
	}
}
