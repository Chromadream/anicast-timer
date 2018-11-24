package main

import (
	"fmt"

	"github.com/Chromadream/anicast-timer/utility"
	"github.com/bwmarrin/discordgo"
)

var (
	botID          string
	discord        discordgo.Session
	commandPrefix  string
	recordingTimer utility.Timer
)

const recordStart = "The recording starts now, please use `&duration` for timekeeping!"
const currentLength = "Current episode length at this point: %s"
const finalLength = "Recording is done. Raw episode length is: %s"

func main() {
	discord, err := discordgo.New(utility.CompleteKey(DiscordKey))
	if err != nil {
		fmt.Println(fmt.Errorf("%+v", err))
		return
	}
	user, err := discord.User("@me")
	botID = user.ID
	if err != nil {
		fmt.Println(fmt.Errorf("%+v", err))
		return
	}
	fmt.Println(user)
	discord.AddHandler(checkMessage)
	discord.AddHandler(setReadyGo)
	err = discord.Open()
	if err != nil {
		fmt.Println(fmt.Errorf("%+v", err))
		return
	}
	defer discord.Close()
	commandPrefix = "&"
	<-make(chan struct{})
}

func checkMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	sender := m.Author
	channel := m.ChannelID
	if sender.ID == botID || sender.Bot {
		return
	}
	fmt.Println(m.Content)
	switch m.Content {
	case ":craig:, join":
		startTimer(s, channel)
	case "&start":
		startTimer(s, channel)
	case "&duration":
		duration := recordingTimer.GetDuration()
		message := fmt.Sprintf(currentLength, duration)
		_, err := s.ChannelMessageSend(channel, message)
		if err != nil {
			fmt.Println(fmt.Errorf("%+v", err))
			return
		}
	case ":craig:, leave":
		endTimer(s, channel)
	case "&end":
		endTimer(s, channel)
	}
}

func startTimer(s *discordgo.Session, channel string) {
	recordingTimer = utility.StartTiming()
	_, err := s.ChannelMessageSend(channel, recordStart)
	if err != nil {
		fmt.Println(fmt.Errorf("%+v", err))
		return
	}
}

func endTimer(s *discordgo.Session, channel string) {
	duration := recordingTimer.GetDuration()
	message := fmt.Sprintf(finalLength, duration)
	_, err := s.ChannelMessageSend(channel, message)
	if err != nil {
		fmt.Println(fmt.Errorf("%+v", err))
		return
	}
}

func setReadyGo(s *discordgo.Session, r *discordgo.Ready) {
	err := discord.UpdateStatus(0, "a fucking bot")
	if err != nil {
		fmt.Println(fmt.Errorf("%+v", err))
		return
	}
	servers := discord.State.Guilds
	fmt.Println(servers)
}
