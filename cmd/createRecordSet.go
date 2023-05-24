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
	createRecordName string
	createRecordType string
	createRecords    string
	createZoneId     string
)

var createRecordCmd = &cobra.Command{
	Use:   "record",
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

	createRecordCmd.PersistentFlags().StringVar(&createRecordType, "type", "", "The DNS Record Type.")
	_ = createRecordCmd.MarkPersistentFlagRequired("type")

	createRecordCmd.PersistentFlags().StringVar(&createZoneId, "zone-id", "", "The DNS Zone ID.")
	_ = createRecordCmd.MarkPersistentFlagRequired("zone-id")

	createRecordCmd.PersistentFlags().StringVar(&createRecords, "records", "", "The DNS Record IP.")
	_ = createRecordCmd.MarkPersistentFlagRequired("records")
}

func createRecord() error {
	client, err := createClient()
	if err != nil {
		return err
	}

	request := api.PostV1ProjectsProjectIdZonesZoneIdRrsetsJSONRequestBody{
		// Name user given name
		Name: createRecordName,

		// DNS Record IP
		Records: []api.RrsetRecordPost{{Content: createRecords}},

		// DNS Type
		Type: api.RrsetRRSetPostType(createRecordType),
	}
	response, err := client.PostV1ProjectsProjectIdZonesZoneIdRrsets(context.Background(), projectId, createZoneId, request)
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(request)
	if err != nil {
		return fmt.Errorf("failed to create instance: %w", err)
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("unexpected status code: %s", response.Status)
	}

	var instance api.RrsetResponseRRSet
	if err := json.NewDecoder(response.Body).Decode(&instance); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if err := printRset(instance.Rrset); err != nil {
		return err
	}
	return nil
}

func printRset(rset api.DomainRRSet) error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer writer.Flush()
	_, err := fmt.Fprintf(writer, "Id\tName\tRecords\tType\tActive\n")
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%t\n", rset.Id, rset.Name, rset.Records, rset.Type, rset.Active)
	if err != nil {
		return fmt.Errorf("failed to write to tabwriter: %w", err)
	}
	return nil
}
