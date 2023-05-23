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
	createZoneName    string
	createDNSZoneName string
)

var createZoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Creates a new DNS Zone.",
	Long:  `Creates a new DNS Zone.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := createZone(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	createCmd.AddCommand(createZoneCmd)

	createZoneCmd.PersistentFlags().StringVar(&createZoneName, "name", "", "The name to set for the instance.")
	_ = createZoneCmd.MarkPersistentFlagRequired("name")

	createZoneCmd.PersistentFlags().StringVar(&createDNSZoneName, "dns-name", "", "The Domain DNS Name.")
	_ = createZoneCmd.MarkPersistentFlagRequired("dns-name")
}

func createZone() error {
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
