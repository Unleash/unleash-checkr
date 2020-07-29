package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	clih "github.com/apex/log/handlers/cli"
	"github.com/wesleimp/unleash-checkr/cmd/cli"
)

var (
	version = "0.1.0"
)

func main() {
	log.SetHandler(clih.Default)

	fmt.Println()
	defer fmt.Println()

	err := cli.Run(version, os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
