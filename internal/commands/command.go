package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Names() []string // name with aliases
	Exec(ctx *context) error
}

type context struct {
	Session *discordgo.Session
	Message *discordgo.Message // the msg that triggered the cmd
	Args    []string
}

var (
	prefix string
	cmdMap = make(map[string]Command) // maps a string to Command
)

func SetCommandPrefix(p string) {
	prefix = p
}

func RegisterCommand(cmd Command) {
	for _, alias := range cmd.Names() {
		cmdMap[alias] = cmd
	}
}

// will be called for every new message due to AddHandler
func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	// parse the command
	split := strings.Split(m.Content[len(prefix):], " ")
	if len(split) < 1 {
		return
	}
	cmdname := strings.ToLower(split[0])
	cmdargs := split[1:]

	// map the cmdname with a Command
	cmd, ok := cmdMap[cmdname]
	if !ok || cmd == nil {
		return // could not find a valid command
	}

	// prepare to execute
	ctx := &context{
		Session: s,
		Message: m.Message,
		Args:    cmdargs,
	}
	if err := cmd.Exec(ctx); err != nil {
		s.ChannelMessageSend(m.ChannelID, "uhm...")
	}
}
