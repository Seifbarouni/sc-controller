package main

import (
	"fmt"
	"log"

	"github.com/Seifbarouni/sc-controller/pkg/colorize"
	"github.com/Seifbarouni/sc-controller/pkg/data"
	"github.com/Seifbarouni/sc-controller/pkg/validators"
)

func main() {
	_, err := data.Init()
	if err != nil {
		log.Fatal(colorize.C("red", err.Error()))
	}
	yaml := validators.Init()
	c, err := yaml.Validate()
	if err != nil {
		log.Fatal(colorize.C("red", err.Error()))
	}
	fmt.Println(colorize.C("green", "Configuration is valid!"))
	_ = c
}
