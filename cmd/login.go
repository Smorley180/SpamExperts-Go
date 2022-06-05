package cmd

import (
	"fmt"
	"spamexpertscli/libs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login [user]",
	Short: "Login to an account.",
	Long:  "Login to an account, accepts domain accounts, users and admins owned by the authenticated user.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//spamURI := libs.Buildrequest("/api/authticket/create/username/" + args[0] + "/identifier/" + viper.GetString("identifier"))

		//response := libs.Makerequest(spamURI)

		response := libs.Queryapi("/api/authticket/create/username/" + args[0] + "/identifier/" + viper.GetString("identifier"))

		fmt.Println("https://" + viper.GetString("hostname") + "/?authticket=" + response)
	},
}
