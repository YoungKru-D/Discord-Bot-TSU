package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var token = "YOUR_BOT_TOKEN"

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddMessageCreateHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if the message starts with the command prefix
	if m.Content == "!ping" {
		// Reply with "Pong!"
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
