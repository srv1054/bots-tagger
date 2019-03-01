package tagger

import (
	"strings"

	"github.com/nlopes/slack"
)

// DrEvil - Dr. Evil Tagger
func DrEvil(myBot *MyBot, ev *slack.MessageEvent) {
	evilWords := []string{
		"laser",
		"laser beams",
		"laser beam",
		"laser-beams",
		"laser-beam",
		"sharks with frickin ",
		"sharks with frickn' ",
		"one million dollars",
		"million dollars",
		"mini-me",
		"mini me",
	}

	if contains(evilWords, strings.ToLower(ev.Msg.Text)) {
		var payload ReactionPayload

		payload.Channel = ev.Channel
		payload.Name = "drevil"
		payload.Token = myBot.SlackToken
		payload.TimeStamp = ev.Timestamp

		err := AddReaction(myBot, payload)
		if err != nil {
			errTrap(myBot, "Dr Evil error catch: ", err)
		}
	}
}

// Bizcat - Business Cat Tagger
func Bizcat(myBot *MyBot, ev *slack.MessageEvent) {

	buzzWords := []string{
		"synergy",
		"disruptor",
		"low-hanging fruit",
		"low hanging fruit",
		"on my radar",
		"like a boss",
		"has legs",
		"break through the clutter",
		"diversify",
		"exit strategy",
		"heavy lifiting",
		"holistic approach",
		"mind share",
		"on the runway",
		"outside the box",
		"organic growth",
		"strategic approach",
		"win-win",
		"win/win",
		"best of breed",
		"cloud computing",
		"core competency",
		"event horizon",
		"herding cats",
		"hyperlocal",
		"make it pop",
		"mindshare",
		"pain point",
		"paralysis by analysis",
		"analysis paralysis",
		"profit center",
		"synergistic",
		"ROI",
		"value added",
		"value-added",
		"date for a date",
		"water under the bridge",
		"robust",
		"center of excellence",
		"30,000 foot view",
		"30000 foot view",
		"target audience",
		"big picture",
		"take offline",
		"drop the ball",
		"dropped the ball",
		"drill down",
		"client focused",
		"circle back",
		"lean in",
		"paradigm shift",
		"level set",
		"top down",
		"transformative",
		"vertical market",
		"mindset",
		"service oriented",
		"throw it over the wall",
	}

	if contains(buzzWords, strings.ToLower(ev.Msg.Text)) {
		var payload ReactionPayload

		payload.Channel = ev.Channel
		payload.Name = "businesscat"
		payload.Token = myBot.SlackToken
		payload.TimeStamp = ev.Timestamp

		err := AddReaction(myBot, payload)
		if err != nil {
			errTrap(myBot, "Business Cat error catch: ", err)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}
