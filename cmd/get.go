package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns resources from DNS Zones.",
	Long:  `Returns resources from one or more DNS Zones.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
