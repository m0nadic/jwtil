package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a token and output the claims",
	Long:  `Verify a token and pretty prints the claims.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("verify called")
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
