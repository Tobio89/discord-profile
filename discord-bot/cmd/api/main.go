package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type App struct {
	Bot *discordgo.Session
}

func init() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Panic("could not load dev.env file")
	}
}

func main() {
	log.Println("# discord bot")

	token := os.Getenv("BOT_TOKEN")

	bot := createSession(token)

	app := &App{Bot: bot}

	app.addEventHandlers()

	log.Println("starting bot session")
	startSession(bot)

	// Create channel, hold it open
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	log.Println("discord bot shutting down")
	app.Bot.Close()
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
