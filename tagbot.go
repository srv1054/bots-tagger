package main

/* Written by srv1054 (https://github.com/srv1054)
   See LICENSE file for usage and "Stealability"

   Tags slack message with Emojis based on what's said.  Defaults to Business Cat
*/

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/srv1054/bots-tagger/tagger"
)

func main() {

	var attachments tagger.Attachment

	// Load Configuration
	myBot, err := tagger.LoadBotConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	slackhook := flag.String("slackhook", "", "Slack Webhook")
	slacktoken := flag.String("slacktoken", "", "Slack Bot Token")

	flag.Parse()

	myBot.SlackHook = *slackhook
	myBot.SlackToken = *slacktoken
	if myBot.SlackHook == "" || myBot.SlackToken == "" {
		fmt.Println("\nWarning CLI parameters: -slacktoken and -slackhook are required!")
		os.Exit(0)
	}

	// Announce startup to logs
	tagger.LogToSlack("*Tagger starting up.* `Version: "+myBot.Version+"`", myBot, attachments)
	if myBot.Debug {
		fmt.Println("Tagger starting up. Version: " + myBot.Version)
	}

	// Slack RTM initilization
	api := slack.New(myBot.SlackToken)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.ConnectedEvent:
			myBot.BotID = ev.Info.User.ID
			myBot.BotName = ev.Info.User.Name
			myBot.TeamID = ev.Info.Team.ID
			myBot.TeamName = ev.Info.Team.Name

		case *slack.MessageEvent:

			if strings.Contains(ev.Msg.Text, "@"+myBot.BotID) || string(ev.Msg.Channel[0]) == "D" {
				// 411 Info or verison info
				if strings.Contains(ev.Msg.Text, "your 411") {
					rtm.SendMessage(rtm.NewOutgoingMessage("My name is "+myBot.BotName+", I tag comments.  My ID is "+myBot.BotID+" and I'm part of the "+myBot.TeamName+" team (ID: "+myBot.TeamID+").  This channels ID is "+ev.Msg.Channel+". Your Slack UID is "+ev.Msg.User, ev.Msg.Channel))
				}

			}

			tagger.Bizcat(myBot, ev)
			tagger.DrEvil(myBot, ev)

		case *slack.PresenceChangeEvent:
			//fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			//fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

		}
	}
}