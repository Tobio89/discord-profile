package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (app *App) onReady(s *discordgo.Session, _ *discordgo.Ready) {
	if app.Initialized {
		log.Println("bot reconnected")
		return
	}

	log.Println("bot is ready")
	app.Bot.ChannelMessageSend("904628203430219787", "I AM HERE")
	app.Initialized = true

}

func (app *App) addEventHandlers() {
	app.Bot.AddHandler(app.onReady)
}
