package cmd

import (
	"fmt"
	"spamexpertscli/libs"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(destinationCmd)
	destinationCmd.AddCommand(setCmd)
	destinationCmd.AddCommand(getCmd)
}

var destinationCmd = &cobra.Command{
	Use:   "destination",
	Short: "Manipulate the destination of a domain.",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		spamURI := libs.Buildrequest("/api/domain/getroute/domain/" + args[0])

		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}

var getCmd = &cobra.Command{
	Use:   "get [domain]",
	Short: "Get the destinations for a domain.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		spamURI := libs.Buildrequest("/api/domain/getroute/domain/" + args[0])

		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
}

var setCmd = &cobra.Command{
	Use:   "set [domain] [*destinations]",
	Short: "Set the destinations for a domain.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		destinations := libs.Arrayformat(args[1:])
		spamURI := libs.Buildrequest("/api/domain/edit/domain/" + args[0] + "/destinations/" + destinations)
		response := libs.Makerequest(spamURI)

		fmt.Println(response)
	},
} //
