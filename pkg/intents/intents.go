package intents

import "github.com/bwmarrin/discordgo"

func SetIntents() discordgo.Intent {
	return discordgo.IntentsAll
}
