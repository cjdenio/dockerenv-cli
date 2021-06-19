package main

import (
	"log"

	"github.com/cjdenio/dockerenv-cli/pkg/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
