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

var getZonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Returns all Zones from the DNS.",
	Long:  `Returns all Zones from the DNS.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getZone(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getZonesCmd)
}

func getZones() error {
	client, err := createClient()
	if err != nil {
		return err
	}

	size := int(100)
	start := int(1)
	empty := string("")
	parameters := api.GetV1ProjectsProjectIdZonesParams{
		PageSize:    &size,
		Page:        &start,
		DnsNameEq:   &empty,
		DnsNameLike: &empty,
	}

	response, err := client.GetV1ProjectsProjectIdZones(context.Background(), projectId, &parameters)
	if err != nil {
		return fmt.Errorf("failed to get zones: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var zone api.ZoneResponseZoneAll
	if err := json.NewDecoder(response.Body).Decode(&zone); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printZones(zone.Zones); err != nil {
		return err
	}
	return nil
}

func printZones(zones []api.DomainZone) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tDescription\tActive\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, zone := range zones {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%t\n", zone.Id, zone.Name, zone.Description, zone.Active)
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}
