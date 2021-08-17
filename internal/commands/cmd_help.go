package commands

type CmdHelp struct{}

func (c *CmdHelp) Names() []string {
	return []string{"help", "h"}
}

func (c *CmdHelp) Exec(ctx *context) (err error) {
	m := "current list <https://gist.github.com/ichigo-gyuunyuu/f2e84adbe3a8243ffe28e29e42ceabd2>"
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, m)
	return
}
