package data

type Repository interface {
	GetUserByDiscordID(discordID string) (*DiscordUser, error)
	InsertDiscordUser(discordUser DiscordUser) (int, error)
	ReplaceMagicLinkForUser(magicLink MagicLink) (int64, error)
	ConsumeMagicLink(tokenHash string) (string, error)
}
