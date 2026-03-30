package main

import (
	"discord-profile/discord-bot/commands"
	"discord-profile/discord-bot/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type App struct {
	Commands    commands.Commands
	Initialized bool
	Config      config.Config
}

var bot *discordgo.Session

func init() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		err := godotenv.Load("dev.env")
		if err != nil {
			log.Panic("could not load envs")
		}
	}
}

func main() {
	log.Println("# discord bot")

	cfg := config.Config{
		Token:   os.Getenv("BOT_TOKEN"),
		BotID:   os.Getenv("BOT_ID"),
		GuildID: os.Getenv("GUILD_ID"),
	}

	bot = createSession(cfg.Token)

	app := App{
		Config:      cfg,
		Initialized: false,
		Commands:    commands.New(bot, &cfg),
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsAllWithoutPrivileged | discordgo.IntentMessageContent

	app.addEventHandlers()
	app.Commands.InitializeCommands()

	log.Println("starting bot session")
	app.StartSession()

	// Create channel, hold it open
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	app.Initialized = false

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
func (app *App) StartSession() {
	err := bot.Open()
	if err != nil {
		log.Panic(err)
	}
}
