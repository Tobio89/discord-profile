package data

type Repository interface {
	GetUserByDiscordID(discordID string) (*DiscordUser, error)
	InsertDiscordUser(discordUser DiscordUser) (int, error)
}
