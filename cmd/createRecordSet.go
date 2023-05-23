package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BackInBash/dns-cli/internal/api"
	"github.com/spf13/cobra"
)

var (
	createRecordName string
	createRecordType string
	createRecords    string
)

var createRecordCmd = &cobra.Command{
	Use:   "redord",
	Short: "Creates a new DNS Record.",
	Long:  `Creates a new DNS Record.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := createRecord(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createRecordCmd)

	createRecordCmd.PersistentFlags().StringVar(&createRecordName, "name", "", "The name to set for the instance.")
	_ = createRecordCmd.MarkPersistentFlagRequired("name")

	createZoneCmd.PersistentFlags().StringVar(&createRecordType, "dns-name", "", "The Domain DNS Name.")
	_ = createZoneCmd.MarkPersistentFlagRequired("dns-name")
}

func createRecord() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	request := api.PostV1ProjectsProjectIdZonesJSONRequestBody{
		// DnsName zone name
		DnsName: createDNSZoneName,

		// Name user given name
		Name: createZoneName,
	}
	response, err := client.PostV1ProjectsProjectIdZones(context.Background(), projectId, request)
	if err != nil {
		return fmt.Errorf("failed to create instance: %w", err)
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var instance api.ZoneResponseZone
	if err := json.NewDecoder(response.Body).Decode(&instance); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printZone(instance.Zone); err != nil {
		return err
	}
	return nil
}
