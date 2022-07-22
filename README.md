# Waffer - Discord Bot
[![Go Report Card](https://goreportcard.com/badge/github.com/nicolito128/waffer)](https://goreportcard.com/report/github.com/nicolito128/waffer)

A discord bot written in Go.


<img align="right" alt="Waffer logo" src="https://i.imgur.com/guq55Wb_d.png" width="150">

## Deploy
Clone the repository:

    git clone https://github.com/nicolito128/waffer

Get packages:

    go get

You need to set a "BOT_TOKEN" and a "BOT_PREFIX" environment variables, for more information about that consult .env.example. Also, you can use `--token` and `--prefix` flags to start the bot.

Run app:

    go run main.go --prefix <PREFIX> --token <YOUR_TOKEN>

## Interest links
* [bwmarrin/discordgo][1]

[1]: https://github.com/bwmarrin/discordgo