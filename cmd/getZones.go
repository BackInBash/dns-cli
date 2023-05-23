package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/BackInBash/dns-cli/internal/api"

	"github.com/spf13/cobra"
)

var getZonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Returns all Zones from the DNS.",
	Long:  `Returns all Zones from the DNS.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getZones(); err != nil {
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

	parameters := &api.GetV1ProjectsProjectIdZonesParams{}

	response, err := client.GetV1ProjectsProjectIdZones(context.Background(), projectId, parameters)
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
	_, err := fmt.Fprintf(writer, "Id\tName\tDNS\tDescription\tActive\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, zone := range zones {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", zone.Id, zone.Name, zone.DnsName, *zone.Description, strconv.FormatBool(*zone.Active))
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}
