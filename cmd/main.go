package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"

	"github.com/rydyb/osci"
	"github.com/rydyb/telnet"
)

var cli struct {
	Debug        bool     `name:"debug" default:"false"`
	Timeout      int      `name:"timeout" default:"1"`
	Host         string   `name:"host" required:""`
	Port         int      `name:"port" required:""`
	Identity     struct{} `cmd:"identity"`
	Measurements struct{} `cmd:"measurements"`
}

func main() {
	args := kong.Parse(&cli)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cli.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	client := &osci.Client{
		Client: telnet.Client{
			Timeout: time.Duration(cli.Timeout) * time.Second,
			Address: fmt.Sprintf("%s:%d", cli.Host, cli.Port),
		},
	}
	if err := client.Open(); err != nil {
		log.Fatal().Err(err).Msg("failed to open client connection")
	}
	defer client.Close()

	switch args.Command() {
	case "identity":
		out, err := client.Identity()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to query identity")
		}
		fmt.Println(out)
	case "measurements":
		out, err := client.Measurements()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to query measurements")
		}
		fmt.Println(strings.Join(maps.Keys(out), ", "))
		fmt.Println(strings.Join(maps.Values(out), ", "))
	default:
		log.Fatal().Msgf("invalid command: %s", args.Command())
	}
}
