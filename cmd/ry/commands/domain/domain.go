package domain

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/domain"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	domainCmd := &cobra.Command{
		Use:   "domain",
		Short: "Manage domains",
	}

	domainListCmd := &cobra.Command{
		Use:   "list",
		Short: "List domains",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetDomainList(&domain.GetDomainListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			domains := make([]model.Domain, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				domains[i] = model.Domain{
					ID:   strconv.Itoa(r.ID),
					Name: r.DomainName,
				}
			}
			return (*out).Print(domains)
		},
	}

	dnsCmd := &cobra.Command{
		Use:   "dns",
		Short: "Manage DNS records",
	}

	dnsListCmd := &cobra.Command{
		Use:   "list <domain-id>",
		Short: "List DNS records",
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetDomainDNSRecordList(id)
			if err != nil {
				return err
			}
			records := make([]model.DNSRecord, len(resp.Data))
			for i, r := range resp.Data {
				records[i] = model.DNSRecord{
					ID:    strconv.Itoa(r.ID),
					Type:  r.RecordType,
					Name:  r.Host,
					Value: r.Value,
					TTL:   r.TTL,
				}
			}
			return (*out).Print(records)
		},
	}

	dnsAddCmd := &cobra.Command{
		Use:   "add <domain-id>",
		Short: "Add a DNS record",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			recType, _ := cmd.Flags().GetString("type")
			host, _ := cmd.Flags().GetString("host")
			value, _ := cmd.Flags().GetString("value")
			ttl, _ := cmd.Flags().GetInt("ttl")
			req := &domain.AddDomainDNSRecordRequest{
				Type:  recType,
				Host:  host,
				Value: value,
				TTL:   ttl,
				Line:  "DEFAULT",
				Level: 10,
			}
			if _, err := (*rySDK).AddDomainDNSRecord(id, req); err != nil {
				return err
			}
			fmt.Printf("DNS record added for domain %s\n", args[0])
			return nil
		},
	}

	dnsDeleteCmd := &cobra.Command{
		Use:   "delete <domain-id> <record-id>",
		Short: "Delete a DNS record",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			domainID, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			recordID, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteDomainDNSRecord(domainID, recordID); err != nil {
				return err
			}
			fmt.Printf("DNS record %s deleted for domain %s\n", args[1], args[0])
			return nil
		},
	}

	dnsAddCmd.Flags().String("type", "A", "Record type (A, AAAA, CNAME, TXT, MX, SRV)")
	dnsAddCmd.Flags().String("host", "@", "Record host name")
	dnsAddCmd.Flags().String("value", "", "Record value")
	dnsAddCmd.Flags().Int("ttl", 600, "TTL in seconds")
	dnsAddCmd.MarkFlagRequired("value")

	dnsCmd.AddCommand(dnsListCmd)
	dnsCmd.AddCommand(dnsAddCmd)
	dnsCmd.AddCommand(dnsDeleteCmd)
	domainCmd.AddCommand(domainListCmd)
	domainCmd.AddCommand(dnsCmd)

	return domainCmd
}
