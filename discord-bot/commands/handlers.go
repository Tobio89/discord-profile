package commands

import (
	rpc "discord-profile/discord-bot/rpc"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (c *Commands) HandleLoginRequest(i *discordgo.InteractionCreate) {
	// Set response to be ephemeral and deferred
	flags := discordgo.MessageFlags(0)
	flags = discordgo.MessageFlagsEphemeral

	// Start the deferred response (waiting...)
	if err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Flags: flags},
	}); err != nil {
		log.Printf("failed to send deferred response: %v", err)
		return
	}

	// Request signup using RPC
	result, err := rpc.RPCRequestLogin(rpc.RPCLoginPayload{
		Username: i.Member.User.Username,
		ID:       i.Member.User.ID,
		Token:    i.Member.User.Token,
	})

	var content string

	// RPC Error handling
	if err != nil || result.Error == true {
		log.Println("error while making RPC call: ", err)
		content = "Sorry, something went wrong while processing your request. Please try again later."
		if _, err := bot.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		}); err != nil {
			log.Printf("failed to edit deferred response: %v", err)
		}
		return
	}

	// Set response content based on RPC result
	content = "Looks like you're new to the app! Use `/create-profile` to get started."
	if result.Success == true {
		content = "Here's your login link: " + result.URL
	}

	// Respond to signup interaction
	if _, err := bot.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	}); err != nil {
		log.Printf("failed to edit deferred response: %v", err)
	}

}

func (c *Commands) HandleSignupRequest(i *discordgo.InteractionCreate) {
	// Set response to be ephemeral and deferred
	flags := discordgo.MessageFlags(0)
	flags = discordgo.MessageFlagsEphemeral

	// Start the deferred response (waiting...)
	if err := bot.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Flags: flags},
	}); err != nil {
		log.Printf("failed to send deferred response: %v", err)
		return
	}

	// Request signup using RPC
	result, err := rpc.RPCRequestSignup(rpc.RPCSignupPayload{
		Username: i.Member.User.Username,
		ID:       i.Member.User.ID,
		Token:    i.Member.User.Token,
	})

	var content string

	// RPC Error handling
	if err != nil {
		log.Println("error while making RPC call: ", err)
		content = "Sorry, something went wrong while processing your request. Please try again later."
		if _, err := bot.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &content,
		}); err != nil {
			log.Printf("failed to edit deferred response: %v", err)
		}
		return
	}

	// Set response content based on RPC result
	content = "Signup successful! You can now log in using the /login command."
	if result.AlreadyExists {
		content = "It looks like you already have an account. You can log in using the /login command."
	}

	// Respond to signup interaction
	if _, err := bot.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	}); err != nil {
		log.Printf("failed to edit deferred response: %v", err)
	}

}
