package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a claim object",
	Long:  `Signs a claim object from a file or from the standard input.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sign called")
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
}
