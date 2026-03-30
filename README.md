# Discord Profile

## About

It's a social profile page app that connect to Discord via a bot.

It's built in Go, and uses microservices

## Running it

- Pull the repository
- Create a bot on Discord's developer portal - https://discord.com/developers/applications
- Create an env file
- Copy the bot's token from the Bot page into the env file - `BOT_TOKEN`
- Copy the bot's ID from the OAuth2 page into the env file - `BOT_ID`
- Add `GUILD_ID` to the env file - this can be left as "", or it can be set to the ID of your Discord server.
- Run `make up_build` to build and add the docker images to your local docker