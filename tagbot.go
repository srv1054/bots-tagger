package main

/* Written by srv1054 (https://github.com/srv1054)
   See LICENSE file for usage and "Stealability"

   Tags slack message with Emojis based on what's said.  Defaults to Business Cat
*/

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nlopes/slack"
	"github.com/srv1054/bots-tagger/tagger"
)

func main() {

	var attachments tagger.Attachment
	var Paint tagger.SprayCans
	var message string
	var hmessage string
	var payload tagger.BotDMPayload

	// Load Configuration
	myBot, err := tagger.LoadBotConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	myBot.Version = "1.5"

	slackhook := flag.String("slackhook", "", "Slack Webhook")
	slacktoken := flag.String("slacktoken", "", "Slack Bot Token")
	version := flag.Bool("v", false, "Tagger Version")

	flag.Parse()

	if *version {
		fmt.Println("Tagger v" + myBot.Version)
		os.Exit(0)
	}

	myBot.SlackHook = *slackhook
	myBot.SlackToken = *slacktoken
	if myBot.SlackHook == "" || myBot.SlackToken == "" {
		fmt.Println("\nWarning CLI parameters: -slacktoken and -slackhook are required!")
		os.Exit(0)
	}

	// Load tag.json data
	Paint, err = tagger.LoadSprayCans()
	if err != nil {
		fmt.Println("Could not load tags.json, exiting tagger")
		tagger.LogToSlack("Could not load `tags.json`, exiting tagger", myBot, attachments)
		os.Exit(2)
	} else {
		tagger.LogToSlack("Spray Paint Data loaded from `tags.json`! I'm watching *"+strconv.Itoa(len(Paint))+"* different tags.", myBot, attachments)
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

				if strings.Contains(strings.ToLower(ev.Msg.Text), "reload tags") {
					userInfo, _ := api.GetUserInfo(ev.Msg.User)
					tagger.LogToSlack("Reloading tags per request from "+userInfo.Name, myBot, attachments)
					Paint, _ = tagger.LoadSprayCans()
					rtm.SendMessage(rtm.NewOutgoingMessage("Tags were reloaded from `tags.json`", ev.Msg.Channel))
				}

				if strings.Contains(strings.ToLower(ev.Msg.Text), "show all tags") {
					message = ""
					for _, p := range Paint {
						hmessage = "Keywords for tag :" + p.Spray + ":\n"
						for _, w := range p.Words {
							message = message + w + "\n"
						}

						userInfo, _ := api.GetUserInfo(ev.Msg.User)

						payload.Text = hmessage
						payload.Channel = userInfo.ID
						attachments.Color = "#00ff00"
						attachments.Text = message
						payload.Attachments = append(payload.Attachments, attachments)

						tagger.WranglerDM(myBot, payload)

						message = ""
						payload.Attachments = nil
					}
					rtm.SendMessage(rtm.NewOutgoingMessage("I have sent you a DM with your results! :love_letter:", ev.Msg.Channel))

				}

				if strings.Contains(strings.ToLower(ev.Msg.Text), "show keywords for ") {
					message = ""
					hmessage = ""

					cleanMsg := strings.Replace(strings.ToLower(ev.Msg.Text), "<@"+strings.ToLower(myBot.BotID)+"> ", "", -1)
					colonStrip := strings.Replace(cleanMsg, ":", "", -1)
					whatsaid := strings.Split(colonStrip, " ")

					if len(whatsaid) == 4 {
						tag := whatsaid[3]

						for _, p := range Paint {
							if tag == p.Spray {
								hmessage = "Keywords for tag :" + p.Spray + ":\n"
								for _, w := range p.Words {
									message = message + w + "\n"
								}

								userInfo, _ := api.GetUserInfo(ev.Msg.User)

								payload.Text = hmessage
								payload.Channel = userInfo.ID
								attachments.Color = "#00ff00"
								attachments.Text = message
								payload.Attachments = append(payload.Attachments, attachments)

								tagger.WranglerDM(myBot, payload)

								message = ""
								payload.Attachments = nil
							}
						}
						if hmessage == "" {
							rtm.SendMessage(rtm.NewOutgoingMessage("I could not find a tag with keywords for :"+tag+":, sorry!", ev.Msg.Channel))
						} else {
							rtm.SendMessage(rtm.NewOutgoingMessage("I have sent you a DM with your results! :love_letter:", ev.Msg.Channel))
						}
					} else {
						rtm.SendMessage(rtm.NewOutgoingMessage("I don't understand what you are asking me to do!", ev.Msg.Channel))
					}
				}

			}

			tagger.TagIt(myBot, Paint, ev)

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
