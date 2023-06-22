package main

import (
	"fmt"
	"strings"
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

	//fmt.Println(client.Exec("*idn?"))

	out, err := client.Exec("MEASUrement:LIST?")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to acquire measurement names")
	}
	names := strings.Split(out, ",")
	log.Debug().Strs("measurementNames", names).Msg("acquired measurement names")

	values := make([]string, len(names))
	for i, name := range names {
		value, err := client.Exec(fmt.Sprintf("MEASUrement:%s:VALue?", name))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to acquire list of measurements")
		}
		values[i] = value
		log.Debug().Str("measurementName", name).Str("measurementValue", value).Msg("acquired measurement")
	}
	log.Debug().Strs("measurementValues", values).Msg("received measurement values")

	if !log.Debug().Enabled() {
		fmt.Println(strings.Join(names, ", "))
		fmt.Println(strings.Join(values, ", "))
	}
}
