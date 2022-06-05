package cmd

import (
	"fmt"
	"spamexpertscli/libs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(domainCmd)
	domainCmd.AddCommand(addCmd)
	domainCmd.AddCommand(removeCmd)
}

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Add or view domains.",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		spamURI := libs.Buildrequest("/api/domain/exists/domain/" + args[0])
		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}

var addCmd = &cobra.Command{
	Use:   "add [domain] [*destinations]",
	Short: "Add a domain to SpamExperts.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		destinations := libs.Arrayformat(args[1:])
		spamURI := libs.Buildrequest("/api/domain/add/domain/" + args[0] + "/destinations/" + destinations)
		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}

var removeCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove a domain from SpamExperts.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		spamURI := libs.Buildrequest("/api/domain/remove/domain/" + args[0])
		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}
