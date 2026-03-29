package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (app *App) onReady(s *discordgo.Session, _ *discordgo.Ready) {
	logMessage := "Bot was turned on"
	log.Println(logMessage)

	app.Bot.ChannelMessageSend("904628203430219787", "I AM HERE")

}

func (app *App) addEventHandlers() {
	app.Bot.AddHandler(app.onReady)
}
