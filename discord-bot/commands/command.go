package commands

import (
	"discord-profile/discord-bot/config"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var bot *discordgo.Session

type Commands struct {
	cfg *config.Config
}

func New(s *discordgo.Session, cfg *config.Config) Commands {
	bot = s
	return Commands{
		cfg: cfg,
	}

}

func (c *Commands) InitializeCommands() {
	c.registerCommands()
	c.registerHandlers()
}

func (c *Commands) registerCommands() {
	_, err := bot.ApplicationCommandCreate(c.cfg.BotID, c.cfg.GuildID, LoginCommand)
	if err != nil {
		log.Println("whilst adding login command: ", err)
	}
	_, err = bot.ApplicationCommandCreate(c.cfg.BotID, c.cfg.GuildID, LoginDeferredCommand)
	if err != nil {
		log.Println("whilst adding login-slowly command: ", err)
	}
}

func (c *Commands) registerHandlers() {
	bot.AddHandler(c.regularCommandGroup)
}

func (c *Commands) regularCommandGroup(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	// options := data.Options
	interactionMember := i.Member

	switch data.Name {
	case "login":
		log.Printf("user %s is trying to login via the bot", interactionMember.User.Username)
		c.HandleLoginRequest(i, interactionMember.User.Username)
		// c.SendResponse(i, "Here's your login link: http://localhost:5173/")

	case "login-slowly":
		log.Printf("user %s is trying to login via the bot using deferred login", interactionMember.User.Username)
		c.SendDeferredResponse(i, "After waiting... here's your login link: http://localhost:5173/", true, 3*time.Second)
	}
}

func (c *Commands) SendResponse(ic *discordgo.InteractionCreate, content string) {
	bot.InteractionRespond(ic.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{Content: content,
				Flags: 1 << 6},
		})
}

func (c *Commands) SendDeferredResponse(ic *discordgo.InteractionCreate, content string, ephemeral bool, delay time.Duration) {
	flags := discordgo.MessageFlags(0)
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}

	if err := bot.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Flags: flags},
	}); err != nil {
		log.Printf("failed to send deferred response: %v", err)
		return
	}

	if delay > 0 {
		time.Sleep(delay)
	}

	if _, err := bot.InteractionResponseEdit(ic.Interaction, &discordgo.WebhookEdit{
		Content: &content,
	}); err != nil {
		log.Printf("failed to edit deferred response: %v", err)
	}
}
