package main

import (
	"fmt"
	"time"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cli struct {
	Host    string `arg:"host" name:"host"`
	Port    int    `arg:"port" default:"4000"`
	Debug   bool   `name:"debug" default:"false"`
	Timeout int    `name:"timeout" default:"1"`
}

func main() {
	kong.Parse(&cli)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cli.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	client := &Client{
		Timeout: time.Duration(cli.Timeout) * time.Second,
		Address: fmt.Sprintf("%s:%d", cli.Host, cli.Port),
	}
	if err := client.Open(); err != nil {
		log.Fatal().Err(err).Msg("failed to open client connection")
	}
	defer client.Close()

	fmt.Println(client.Exec("*idn?"))
	fmt.Println(client.Exec("MEASUrement:LIST?"))
}
