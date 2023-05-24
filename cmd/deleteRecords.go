package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	deleteZoneId string
)

var deleteRecordCmd = &cobra.Command{
	Use:   "record <recordId>",
	Short: "Deletes a DNS Record.",
	Long:  `Deletes a DNS Record.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, recordId := range args {
			if err := deleteRecord(recordId); err != nil {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteRecordCmd)

	deleteRecordCmd.PersistentFlags().StringVar(&deleteZoneId, "zone-id", "", "The UUID of the DNS Zone.")
	_ = deleteRecordCmd.MarkPersistentFlagRequired("zone-id")
}

func deleteRecord(recordId string) error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.DeleteV1ProjectsProjectIdZonesZoneIdRrsetsRrSetId(context.Background(), projectId, deleteZoneId, recordId)
	if err != nil {
		return fmt.Errorf("failed to delete instance: %w", err)
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}
	fmt.Printf("Record %s deleted\n", recordId)
	return nil
}
