package cmd

import (
	"fmt"
	"log"

	"github.com/Seifbarouni/sc-controller/pkg/colorize"
	"github.com/Seifbarouni/sc-controller/pkg/validators"
	"github.com/spf13/cobra"
)

var validateCfgCmd = &cobra.Command{
	Use:     "validate",
	Short:   "Validate the sc-controller configuration",
	Aliases: []string{"val"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(colorize.C("white", "Scanning config..."))
		v := validators.Init()
		_, err := v.Validate()
		if err != nil {
			log.Fatal(colorize.C("red", err.Error()))
		}
		fmt.Println(colorize.C("white", "Scan complete."))
		fmt.Println(colorize.C("green", "Configuration is valid!"))
	},
}

func init() {
	rootCmd.AddCommand(validateCfgCmd)
}
