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

var getRecordsCmd = &cobra.Command{
	Use:   "records",
	Short: "Returns all Records from a DNS Zone.",
	Long:  `Returns all Records from a DNS Zone.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := getRecords(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	},
}

func init() {
	getCmd.AddCommand(getRecordsCmd)

	getRecordsCmd.PersistentFlags().StringVar(&getZoneId, "zone-id", "", "The UUID of the DNS Zone.")
	_ = getRecordsCmd.MarkPersistentFlagRequired("zone-id")
}

func getRecords() error {
	client, err := createClient()
	if err != nil {
		return err
	}
	response, err := client.GetV1ProjectsProjectIdZonesZoneIdRrsets(context.Background(), projectId, getZoneId, &api.GetV1ProjectsProjectIdZonesZoneIdRrsetsParams{})
	if err != nil {
		return fmt.Errorf("failed to get zone: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var rsets api.RrsetResponseRRSetAll
	if err := json.NewDecoder(response.Body).Decode(&rsets); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printRecords(rsets.RrSets); err != nil {
		return err
	}
	return nil
}

func printRecords(rsets []api.DomainRRSet) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tType\tEndpoints\tActive\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	for _, rset := range rsets {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\n", rset.Id, rset.Name, rset.Type, rset.Records, strconv.FormatBool(*rset.Active))
		if err != nil {
			return fmt.Errorf("failed to write to tabwriter: %w", err)
		}
	}
	return nil
}
