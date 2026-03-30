package commands

import (
	rpc "discord-profile/discord-bot/rpc"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (c *Commands) HandleLoginRequest(i *discordgo.InteractionCreate, username string) {
	log.Println("sending RPC request")
	result := rpc.MakeRPCCall(username)
	c.SendResponse(i, "Here's your login link: "+result)
}
