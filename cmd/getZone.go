package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/BackInBash/dns-cli/internal/api"

	"github.com/spf13/cobra"
)

var (
	getZoneId string
)

var getZoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "Returns a Zone from the DNS.",
	Long:  `Returns a Zone from the DNS.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getZone(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getZoneCmd)

	getZoneCmd.PersistentFlags().StringVar(&getZoneId, "zone-id", "", "The UUID of the DNS Zone.")
	_ = getZoneCmd.MarkPersistentFlagRequired("zone-id")
}

func getZone() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.GetV1ProjectsProjectIdZonesZoneId(context.Background(), projectId, getZoneId)
	if err != nil {
		return fmt.Errorf("failed to get zone: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var zone api.ZoneResponseZone
	if err := json.NewDecoder(response.Body).Decode(&zone); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printZone(zone.Zone); err != nil {
		return err
	}
	return nil
}

func printZone(zone api.DomainZone) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tDescription\tActive\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%t\n", zone.Id, zone.Name, zone.Description, zone.Active)
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	return nil
}
