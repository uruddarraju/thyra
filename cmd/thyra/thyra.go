package main

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/uruddarraju/thyra/cmd/thyra/subcommand"
)

var cmd = &cobra.Command{
	Use:   "thyra",
	Short: "thyra - an api gateway.",
	Long:  `thyra - An opensource API Gateway service that simplifies creating APIs and managing them.`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.Infof("Ran successfully.......")
	},
}

func init() {

	cmd.AddCommand(subcommand.VersionCmd)
	cmd.AddCommand(subcommand.StartCmd)

	viper.SetEnvPrefix("thyra")
	viper.SetConfigName("thyra")
	viper.AddConfigPath("/Users/alekhya/go/src/github.com/uruddarraju/thyra/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")

	flags := cmd.Flags()
	flags.String("started-by", "uruddarraju", "Testing the command line feature using viper....")

	viper.BindPFlag("started-by", flags.Lookup("started-by"))

	err := viper.ReadInConfig()
	if err != nil {
		glog.Fatalf("Error loading thyra configuration: %v", err)
	}
	glog.V(2).Infof("Successfully loaded thyra config")
}

func main() {
	glog.Infof("Starting up thyra api gateway......")
	cmd.Execute()
	return
}
