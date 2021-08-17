package commands

import (
	"strings"

	"github.com/ichigo-gyuunyuu/widebot/internal/sounds"
	"github.com/ichigo-gyuunyuu/widebot/internal/utils"
)

type CmdPlaysound struct {
	PlayingSound bool
}

func (c *CmdPlaysound) Names() []string {
	return []string{"playsound", "ps"}
}

func (c *CmdPlaysound) Exec(ctx *context) (err error) {
	if c.PlayingSound {
		return
	}

	var audiobuf [][]byte
	var tmp string

	// handle args
	if len(ctx.Args) <= 0 {
		tmp, err = utils.GetEmoteString("pepeg")
		if err != nil {
			return
		}
		msg := "play what you "
		msg += tmp
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg)
		return
	}
	switch strings.ToLower(ctx.Args[0]) {
	case "eatit":
		audiobuf, err = sounds.GetAudioBuffer("./media/eatit.dca")
		if err != nil {
			return
		}
	case "where":
		audiobuf, err = sounds.GetAudioBuffer("./media/where.dca")
		if err != nil {
			return
		}

	default:
		tmp, err = utils.GetEmoteString("noidontthinkso")
		if err != nil {
			return
		}
		msg := tmp
		_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg)
		return
	}

	// get the guild
	g, err := ctx.Session.State.Guild(ctx.Message.GuildID)
	if err != nil {
		return
	}
	// get the VC
	for _, vs := range g.VoiceStates {
		if vs.UserID == ctx.Message.Author.ID {
			// play the sound
			c.PlayingSound = true
			err = playsound(ctx, g.ID, vs.ChannelID, audiobuf)
			c.PlayingSound = false
			return
		}
	}
	// if user not in any vc
	tmp, err = utils.GetEmoteString("feelsrainman")
	if err != nil {
		return
	}
	msg := "who's there to listen to me "
	msg += tmp
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg)
	return
}

func playsound(ctx *context, guildID, channelID string, audiobuf [][]byte) (err error) {
	// join the provided VC
	vc, err := ctx.Session.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	vc.Speaking(true)
	// send the buffer
	for _, buf := range audiobuf {
		vc.OpusSend <- buf
	}
	vc.Speaking(false)
	vc.Disconnect()
	return
}
