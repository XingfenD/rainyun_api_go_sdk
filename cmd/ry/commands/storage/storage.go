package storage

import (
	"fmt"
	"strconv"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/ros"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	storageCmd := &cobra.Command{
		Use:   "storage",
		Short: "Manage object storage",
	}

	storageListCmd := &cobra.Command{
		Use:   "list",
		Short: "List storage instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRosInstanceList(&ros.GetRosInstanceListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			instances := make([]model.StorageInstance, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				instances[i] = model.StorageInstance{
					ID:     strconv.Itoa(r.ID),
					Name:   r.Tag,
					Status: r.Status,
				}
			}
			return (*out).Print(instances)
		},
	}

	bucketCmd := &cobra.Command{
		Use:   "bucket <instance-id>",
		Short: "Manage buckets",
		Args:  cobra.ExactArgs(1),
	}

	bucketListCmd := &cobra.Command{
		Use:   "list",
		Short: "List buckets",
		RunE: func(cmd *cobra.Command, args []string) error {
			instanceID, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid instance id: %s", args[0])
			}
			resp, err := (*rySDK).GetRosBucketListByInstance(instanceID)
			if err != nil {
				return err
			}
			buckets := make([]model.Bucket, len(resp.Data.Records))
			for i, b := range resp.Data.Records {
				buckets[i] = model.Bucket{
					ID:   strconv.Itoa(b.ID),
					Name: b.Name,
				}
			}
			return (*out).Print(buckets)
		},
	}

	bucketCreateCmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a bucket",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			instanceID, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid instance id: %s", args[0])
			}
			name := args[1]
			if _, err := (*rySDK).CreateRosBucket(ros.CreateRosBucketRequest{
				BucketName: name,
				InstanceID: instanceID,
			}); err != nil {
				return err
			}
			fmt.Printf("Bucket %q created\n", name)
			return nil
		},
	}

	bucketCmd.AddCommand(bucketListCmd)
	bucketCmd.AddCommand(bucketCreateCmd)
	storageCmd.AddCommand(storageListCmd)
	storageCmd.AddCommand(bucketCmd)

	return storageCmd
}
