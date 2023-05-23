package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes DNS resources.",
	Long:  `Deletes DNS resources.`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
