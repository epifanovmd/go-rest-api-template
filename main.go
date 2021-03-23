package main

import (
	"go-rest-api-template/app"
	"go-rest-api-template/app/config"
	"log"

	"github.com/BurntSushi/toml"
)

func main() {

	configuration := config.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", configuration)
	if err != nil {
		log.Fatal(err)
	}

	application := &app.App{}

	if err := application.Start(configuration); err != nil {
		log.Fatal(err)
	}
}
