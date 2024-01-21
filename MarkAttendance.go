package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"sync"
)

var (
	token           = "YOUR_BOT_TOKEN"
	attendanceMutex sync.Mutex
	attendanceMap   = make(map[string]bool) // User ID to attendance status
)

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
	case "mark":
		handleMarkCommand(s, m)
	default:
		handleUnknownCommand(s, m)
	}
}

func handlePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func handleHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "Available commands:\n" +
		"!ping - Get a 'Pong!' response\n" +
		"!help - Display this help message\n" +
		"!mark - Mark your attendance"
	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

func handleMarkCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	attendanceMutex.Lock()
	defer attendanceMutex.Unlock()

	// Check if the user has already marked attendance
	if attendanceMap[m.Author.ID] {
		s.ChannelMessageSend(m.ChannelID, "You've already marked your attendance.")
		return
	}

	// Mark the user's attendance
	attendanceMap[m.Author.ID] = true
	s.ChannelMessageSend(m.ChannelID, "Attendance marked for "+m.Author.Username+"!")

	// Additional logic, if needed (e.g., record timestamp, store in a database)
}

func handleUnknownCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Unknown command. Type `!help` for a list of commands.")
}
