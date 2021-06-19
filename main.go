package main

import (
	"fmt"
	"log"

	"github.com/cjdenio/dockerenv-cli/pkg/api"
)

func main() {
	data, err := api.Image("postgres")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", data)
}
