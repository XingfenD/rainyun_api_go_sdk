package server

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addEDiskCommands(serverCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	ediskCmd := &cobra.Command{
		Use:   "edisk",
		Short: "Manage server elastic cloud disks",
	}

	ediskCreateCmd := &cobra.Command{
		Use:   "create <id>",
		Short: "Create an elastic cloud disk",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			size, _ := cmd.Flags().GetInt("size")
			diskType, _ := cmd.Flags().GetString("type")
			backup, _ := cmd.Flags().GetBool("backup")
			tag, _ := cmd.Flags().GetString("tag")
			req := &rcs.RcsManagesElasticCloudDisksRequest{
				Actions: []struct {
					Type   string `json:"type"`
					Action any    `json:"action"`
				}{
					{Type: "create", Action: rcs.RcsManagesElasticCloudDisksCreate{
						SizeInGb: size,
						DiskType: diskType,
						Backup:   backup,
						Tag:      tag,
					}},
				},
			}
			if _, err := (*rySDK).RcsManagesElasticCloudDisks(id, req); err != nil {
				return err
			}
			fmt.Printf("Elastic disk created for server %d\n", id)
			return nil
		},
	}
	ediskCreateCmd.Flags().Int("size", 0, "Disk size in GB (required)")
	ediskCreateCmd.Flags().String("type", "ssd", "Disk type (ssd/hdd)")
	ediskCreateCmd.Flags().Bool("backup", false, "Enable backup support")
	ediskCreateCmd.Flags().String("tag", "", "Disk tag")
	ediskCreateCmd.MarkFlagRequired("size")

	ediskExpandCmd := &cobra.Command{
		Use:   "expand <id>",
		Short: "Expand an elastic cloud disk",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := parseID(args[0])
			if err != nil {
				return err
			}
			ediskID, _ := cmd.Flags().GetInt("edisk-id")
			size, _ := cmd.Flags().GetInt("size")
			backup, _ := cmd.Flags().GetBool("backup")
			req := &rcs.RcsManagesElasticCloudDisksRequest{
				Actions: []struct {
					Type   string `json:"type"`
					Action any    `json:"action"`
				}{
					{Type: "expand", Action: rcs.RcsManagesElasticCloudDisksExpand{
						EdiskID:  ediskID,
						SizeInGb: size,
						Backup:   backup,
					}},
				},
			}
			if _, err := (*rySDK).RcsManagesElasticCloudDisks(id, req); err != nil {
				return err
			}
			fmt.Printf("Elastic disk %d expanded for server %d\n", ediskID, id)
			return nil
		},
	}
	ediskExpandCmd.Flags().Int("edisk-id", 0, "Elastic disk ID (required)")
	ediskExpandCmd.Flags().Int("size", 0, "Target size in GB (required)")
	ediskExpandCmd.Flags().Bool("backup", false, "Enable backup support")
	ediskExpandCmd.MarkFlagRequired("edisk-id")
	ediskExpandCmd.MarkFlagRequired("size")

	ediskCmd.AddCommand(ediskCreateCmd, ediskExpandCmd)
	serverCmd.AddCommand(ediskCmd)
}
