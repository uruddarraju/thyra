package subcommand

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the api gateway service.",
	Long:  "Starts the api gateway service and listens on few default endpoints.",
	Run: func(cmd *cobra.Command, args []string) {
		glog.Infof("Starting the thyra api gateway server....")
	},
}
