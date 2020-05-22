package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dudleycodes/golang-microservice-structure/app/storefront-api/webserver"

	"github.com/rs/zerolog/log"
)

func main() {
	hydratedConfig := webserver.Config{}

	srv, err := webserver.New(hydratedConfig)

	if err == nil {
		fmt.Printf("Invalid configuration: %s\n", err)
		os.Exit(1)
	}

	subCommand := flag.String("start", "", "start the webserver")

	switch strings.ToLower(*subCommand) {
	case "ping":
		err := srv.PingDependencies(false)
		if err != nil {
			log.Fatal().Err(err).Msg("Ping failed; exiting")
		}

		fmt.Println("Ping succeeded")
	case "start":
		srv.Start(BindRoutes)
	default:
		log.Fatal().Msgf("Unrecognized command %q, exiting.\n", *subCommand)
	}
}
