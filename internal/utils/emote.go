package utils

import "github.com/bwmarrin/discordgo"

type emote struct {
	name    string
	guildID string
	emojiID string
}

var (
	emoteMap  = make(map[string]emote)
	dgsession *discordgo.Session
)

func PopulateEmojiMap(s *discordgo.Session) {
	dgsession = s

	feelslagman := &emote{
		name:    "FeelsLagMan",
		guildID: "709983111340884008",
		emojiID: "875783177808007188",
	}

	pepeg := &emote{
		name:    "Pepeg",
		guildID: "709983111340884008",
		emojiID: "876797429788311582",
	}

	noidontthinkso := &emote{
		name:    "NOIDONTTHINKSO",
		guildID: "709983111340884008",
		emojiID: "876797809624490014",
	}

	feelsrainman := &emote{
		name:    "FeelsRainMan",
		guildID: "709983111340884008",
		emojiID: "877108692988854282",
	}

	emoteMap["feelslagman"] = *feelslagman
	emoteMap["pepeg"] = *pepeg
	emoteMap["noidontthinkso"] = *noidontthinkso
	emoteMap["feelsrainman"] = *feelsrainman

	return
}

func GetDiscordEmote(e string) (emote *discordgo.Emoji, err error) {
	emote, err = dgsession.State.Emoji(emoteMap[e].guildID, emoteMap[e].emojiID)
	return
}

func GetEmoteString(e string) (emote string, err error) {
	em, err := dgsession.State.Emoji(emoteMap[e].guildID, emoteMap[e].emojiID)
	if err != nil {
		return
	}
	emote = em.MessageFormat()
	return
}
