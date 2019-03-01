package tagger

import (
	"encoding/json"
	"fmt"
	"os"
)

// MyBot - Bot Configuration Options
type MyBot struct {
	SlackHook  string `json:"slackhook"`
	SlackToken string `json:"slacktoken"`
	BotID      string `json:"botid"`
	BotName    string `json:"botname"`
	TeamID     string `json:"teamid"`
	TeamName   string `json:"teamname"`
	LogChannel string `json:"logchannel"`
	Version    string `json:"version"`
	Debug      bool   `json:"debug"`
}

// LoadBotConfig - Load Main Bot Configuration TOML
func LoadBotConfig() (myBot *MyBot, err error) {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config.json file: " + err.Error() + ".  Must be in running directory.")
		return myBot, err
	}

	decoded := json.NewDecoder(file)
	err = decoded.Decode(&myBot)
	if err != nil {
		fmt.Println("Error reading invalid config.json file: " + err.Error())
		return myBot, err
	}

	if myBot.Debug {
		fmt.Printf("%+v", myBot)
	}

	return myBot, nil
}

// LoadSprayCans - Load tag.json tagger data file
func LoadSprayCans() (spray SprayCans, err error) {

	file, err := os.Open("tags.json")
	if err != nil {
		fmt.Println("Error opening tags.json file: " + err.Error() + ".  Must be in running directory.")
		return spray, err
	}

	decoded := json.NewDecoder(file)
	err = decoded.Decode(&spray)
	if err != nil {
		fmt.Println("Error reading invalid tags.json file: " + err.Error())
		return spray, err
	}

	return spray, nil
}

// errTrap - Generic error handling function
func errTrap(myBot *MyBot, message string, err error) {
	var attachments Attachment

	if myBot.Debug {
		fmt.Println(message + "(" + err.Error() + ")")
	}
	attachments.Color = "#ff0000"
	attachments.Text = err.Error()
	LogToSlack(message, myBot, attachments)
}
