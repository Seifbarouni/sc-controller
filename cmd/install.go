package cmd

import (
	"github.com/Seifbarouni/sc-controller/pkg/workers"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "install the manager process and run the other processes",
	Aliases: []string{"ist"},
	Run: func(cmd *cobra.Command, args []string) {
		workers.Install()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
