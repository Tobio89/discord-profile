package commands

import (
	rpc "discord-profile/discord-bot/rpc"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (c *Commands) HandleLoginRequest(i *discordgo.InteractionCreate) {
	log.Println("sending RPC request")
	// result := rpc.MakeRPCCall(username)
	result, err := rpc.RPCRequestLogin(rpc.LoginRequest{
		Username: i.Member.User.Username,
		ID:       i.Member.User.ID,
		Token:    i.Member.User.Token,
	})

	if err != nil {
		log.Println("error while making RPC call: ", err)
		c.SendResponse(i, "Sorry, something went wrong while processing your request. Please try again later.")
		return
	}

	log.Println("RPC response: ", result)
	c.SendResponse(i, "Here's your login link: "+result)
}
