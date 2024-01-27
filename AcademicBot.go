package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token          = "YOUR TOKEN"
	voiceStates    = make(map[string]map[string]time.Time) // GuildID -> UserID -> JoinTime
	welcomeChannel = ""
)

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.AddHandler(voiceStateUpdate)
	dg.AddHandler(guildMemberAdd)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	fmt.Println("Bot is now running. Press Ctrl+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case strings.Contains(m.Content, "!ping"):
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case strings.HasPrefix(m.Content, "!markattendance"):
		guildID := m.GuildID
		channelID := m.ChannelID

		if voiceStates[guildID] == nil {
			s.ChannelMessageSend(channelID, "No users in voice channels.")
			return
		}
		embed := &discordgo.MessageEmbed{
			Title:  "Attendance List",
			Color:  0x00ff00, // Green color
			Fields: []*discordgo.MessageEmbedField{},
		}

		for userID, joinTime := range voiceStates[guildID] {
			user, err := s.User(userID)
			if err == nil {
				duration := time.Since(joinTime).Truncate(time.Second)
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   user.Username,
					Value:  fmt.Sprintf("Joined %s ago", duration.String()),
					Inline: false,
				})
			}
		}
		s.ChannelMessageSendEmbed(channelID, embed)
	case strings.HasPrefix(m.Content, "!botwelcome"):

		channelID := extractChannelID(m.Content)
		if channelID != "" {
			welcomeChannel = channelID
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Welcome messages will be sent to <#%s>.", welcomeChannel))
		} else {
			s.ChannelMessageSend(m.ChannelID, "Invalid channel. Please mention a valid channel.")
		}
	}
}

func voiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	guildID := vs.GuildID
	userID := vs.UserID

	if vs.ChannelID != "" { // User joined a voice channel
		if voiceStates[guildID] == nil {
			voiceStates[guildID] = make(map[string]time.Time)
		}

		voiceStates[guildID][userID] = time.Now()
	} else { // User left a voice channel
		delete(voiceStates[guildID], userID)
	}
}

func guildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	if welcomeChannel != "" {
		s.ChannelMessageSend(welcomeChannel, fmt.Sprintf("Welcome, %s, to the server!", m.User.Mention()))
	}
}

func extractChannelID(message string) string {
	fields := strings.Fields(message)
	if len(fields) >= 2 {
		channelID := strings.TrimPrefix(fields[1], "<#")
		channelID = strings.TrimSuffix(channelID, ">")
		return channelID
	}
	return ""
}
