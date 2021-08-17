package commands

import (
	"strconv"

	"github.com/ichigo-gyuunyuu/widebot/internal/utils"
)

type CmdPing struct{}

func (c *CmdPing) Names() []string {
	return []string{"ping", "bing"}
}

func (c *CmdPing) Exec(ctx *context) (err error) {
	feelslagman, err := utils.GetEmoteString("feelslagman")
	if err != nil {
		return
	}

	msg := feelslagman
	msg += strconv.FormatInt(ctx.Session.HeartbeatLatency().Milliseconds(), 10) + "ms"
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg)
	return
}
