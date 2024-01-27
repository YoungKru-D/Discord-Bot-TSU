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

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore messages from the bot itself
	}

	// Check if the message starts with the command prefix
	if m.Content[0] != '!' {
		return // Ignore messages that don't start with '!'
	}

	// Parse the command and arguments
	command := m.Content[1:]
	switch command {
	case "ping":
		handlePingCommand(s, m)
	case "help":
		handleHelpCommand(s, m)
	default:
		handleUnknownCommand(s, m)
	}
}

func handlePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Reply with "Pong!"
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func handleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Provide help information
	helpMessage := "Available commands:\n" +
		"!ping - Get a 'Pong!' response\n" +
		"!help - Display this help message"
	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

func handleUnknownCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Reply with an unknown command message
	s.ChannelMessageSend(m.ChannelID, "Unknown command. Type `!help` for a list of commands.")
}
