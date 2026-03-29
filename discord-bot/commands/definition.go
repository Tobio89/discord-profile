package commands

import "github.com/bwmarrin/discordgo"

var LoginCommand = &discordgo.ApplicationCommand{
	Name:        "login",
	Type:        discordgo.ChatApplicationCommand,
	Description: "Login to the profile page via the bot",
}

var LoginDeferredCommand = &discordgo.ApplicationCommand{
	Name:        "login-slowly",
	Type:        discordgo.ChatApplicationCommand,
	Description: "Login to the profile page via the bot",
}
