package cmd

import (
	"log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config File (Defaults to $HOME/spamapi.conf)")
	viper.BindPFlag("username", rootCmd.Flags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.Flags().Lookup("password"))
	viper.BindPFlag("hostname", rootCmd.Flags().Lookup("hostname"))
	viper.BindPFlag("identifier", rootCmd.Flags().Lookup("identifier"))

}

func initConfig() {
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("spamapi.conf")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Can't read config:", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "spamexpertscli [command]",
	Short: "spamexpertscli is for managing a SpamExperts server from the command line.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		println(viper.GetString("password"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
