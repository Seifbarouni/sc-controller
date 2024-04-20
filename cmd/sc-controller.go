package main

import (
	"fmt"
	"log"

	"github.com/Seifbarouni/sc-controller/cmd/data"
	"github.com/Seifbarouni/sc-controller/cmd/validators"
)

var (
	reset     = "\033[0m"
	bold      = "\033[1m"
	underline = "\033[4m"
	strike    = "\033[9m"
	italic    = "\033[3m"

	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cBlue   = "\033[34m"
	cPurple = "\033[35m"
	cCyan   = "\033[36m"
	cWhite  = "\033[37m"
)

func main() {
	_, err := data.Init()
	if err != nil {
		log.Fatal(cRed + err.Error())
	}
	yaml := validators.Init()
	c, err := yaml.Validate()
	if err != nil {
		log.Fatal(cRed + err.Error())
	}
	fmt.Println(cGreen + "Configuration is valid" + reset)
	log.Println(c)
}
