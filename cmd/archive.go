package cmd

import (
	"fmt"
	"spamexpertscli/libs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(archiveCmd)
}

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Get archived domains",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		spamURI := libs.Buildrequest("/api/archive/accounts")

		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}
