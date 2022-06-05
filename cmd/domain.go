package cmd

import (
	"fmt"
	"spamexpertscli/libs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(domainCmd)
}

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Add or view domains.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		spamURI := libs.Buildrequest("/api/domainslist/get/")

		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}
