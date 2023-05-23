package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var deleteZoneCmd = &cobra.Command{
	Use:   "zone <zoneId>",
	Short: "Deletes a DNS zone.",
	Long:  `Deletes a DNS zone.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, zoneId := range args {
			if err := deleteZone(zoneId); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteZoneCmd)
}

func deleteZone(zoneId string) error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.DeleteV1ProjectsProjectIdZonesZoneId(context.Background(), projectId, zoneId)
	if err != nil {
		return fmt.Errorf("failed to delete instance: %w", err)
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("Zone %s deleted\n", zoneId)
	return nil
}
