package commands

import "strconv"

type CmdPing struct{}

func (c *CmdPing) Names() []string {
	return []string{"ping", "bing"}
}

func (c *CmdPing) Exec(ctx *context) (err error) {
	// feelslagman
	emoji, err := ctx.Session.State.Emoji("709983111340884008", "875783177808007188")
	if err != nil {
		return
	}
	msg := emoji.MessageFormat()
	msg += strconv.FormatInt(ctx.Session.HeartbeatLatency().Milliseconds(), 10) + "ms"
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg)
	return
}
