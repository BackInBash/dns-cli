package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates, updates and deletes resources for DNS.",
	Long:  `Creates, updates and deletes resources for DNS.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
