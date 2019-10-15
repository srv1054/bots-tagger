package tagger

import (
	"strings"

	"github.com/nlopes/slack"
)

// Responder - Bot in slack response monitor
func Responder(myBot MyBot, Paint SprayCans, ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client) SprayCans {
	var attachments Attachment
	var message string
	var hmessage string
	var payload BotDMPayload
	var helpload BotDMPayload

	// check only messages that refer to the bot itself
	if strings.Contains(ev.Msg.Text, "@"+myBot.BotID) || string(ev.Msg.Channel[0]) == "D" {
		// 411 Info or verison info
		if strings.Contains(ev.Msg.Text, "your 411") {
			rtm.SendMessage(rtm.NewOutgoingMessage("My name is "+myBot.BotName+", I tag comments.  My ID is "+myBot.BotID+" and I'm part of the "+myBot.TeamName+" team (ID: "+myBot.TeamID+").  This channels ID is "+ev.Msg.Channel+". Your Slack UID is "+ev.Msg.User, ev.Msg.Channel))
		}

		if strings.Contains(strings.ToLower(ev.Msg.Text), "reload tags") {
			attachments.Text = ""
			attachments.Color = ""
			userInfo, _ := api.GetUserInfo(ev.Msg.User)
			LogToSlack("Reloading tags per request from "+userInfo.Name, myBot, attachments)
			Paint, _ = LoadSprayCans(myBot.JSONPath)
			rtm.SendMessage(rtm.NewOutgoingMessage("Tags were reloaded from `tags.json`", ev.Msg.Channel))
		}

		if strings.Contains(strings.ToLower(ev.Msg.Text), "show all tags") {
			message = ""
			payload.Attachments = nil

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

				WranglerDM(myBot, payload)

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

						WranglerDM(myBot, payload)

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

		if strings.Contains(strings.ToLower(ev.Msg.Text), "help") {
			cleanMsg := strings.Replace(strings.ToLower(ev.Msg.Text), "<@"+strings.ToLower(myBot.BotID)+"> ", "", -1)
			helpmin := strings.Split(cleanMsg, " ")
			if len(helpmin) < 2 {

				message := ""
				helpload.Attachments = nil

				userInfo, _ := api.GetUserInfo(ev.Msg.User)

				message = "`show all tags` - I will Direct Message you with all of tags I know about and their associated keywords.\n"
				message = message + "`reload tags` - I will re-read and load the tags.json file. If you make edits to it this will make them effective without restarting me.\n"
				message = message + "`show keywords for <tag name>` - I will show all keywords tied to the specific tag (emoji) listed. For Example:\n"
				message = message + ">`@tagger show keywords for :businesscat: (colons are optional)`\n"
				helpload.Text = "Hi, I'm " + myBot.BotName + ".  I tag emojis on to slack messages containing keywords I know about.\nHere's a few slack commands I know:"
				helpload.Channel = userInfo.ID
				attachments.Color = "#00ffff"
				attachments.Text = message
				helpload.Attachments = append(helpload.Attachments, attachments)

				WranglerDM(myBot, helpload)

				message = ""
				helpload.Attachments = nil
			}
		}

	}

	return Paint
}
