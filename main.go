package main

import (
	"log"
	"os"

	"github.com/wesleimp/unleash-checkr/cmd/cli"
)

var (
	version = "0.1.0"
)

func main() {
	err := cli.Run(version, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
