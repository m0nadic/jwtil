package cmd

import (
	"github.com/spf13/cobra"
	"jwtil/app/util"
	"os"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Pretty prints a JWT token",
	Long:  `Parses a JWT token and pretty prints the different sections.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := util.PrintToken(os.Stdin, false)
		util.ErrorExit(err, 1)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
