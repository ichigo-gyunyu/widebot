package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ichigo-gyuunyuu/widebot/internal/commands"
	"github.com/ichigo-gyuunyuu/widebot/internal/config"
)

func main() {
	// get the config
	const configFile = "./configs/config.yaml"
	cfg, err := config.ParseYAMLConfig(configFile)
	if err != nil {
		panic(err)
	}

	// new discordgo session
	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	// info about guilds
	s.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates
	registerCommands(s, cfg)

	// open a new websocket connection
	if err = s.Open(); err != nil {
		panic(err)
	}
	fmt.Println("connected")

	fmt.Println("tis running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc // keep it running

	// close when notified
	s.Close()
}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	commands.SetCommandPrefix(cfg.Prefix)

	commands.RegisterCommand(&commands.CmdPing{})

	// callback for messagecreate events
	s.AddHandler(commands.HandleMessage)
}
