// Package rcdn implements the ry rcdn command.
package rcdn

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rcdn"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	rcdnCmd := &cobra.Command{
		Use:   "rcdn",
		Short: "Manage RCDN instances and domains",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List RCDN instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcdnInstanceList(&rcdn.GetRcdnInstanceListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			items := make([]model.RcdnInstance, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				items[i] = toRcdnInstance(r)
			}
			return (*out).Print(items)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get RCDN instance details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcdnInstanceDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toRcdnInstance(resp.Data))
		},
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an RCDN instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			planID, _ := cmd.Flags().GetInt("plan-id")
			duration, _ := cmd.Flags().GetInt("duration")
			domains, _ := cmd.Flags().GetString("domains")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if planID == 0 {
				return fmt.Errorf("--plan-id flag is required")
			}
			if duration == 0 {
				return fmt.Errorf("--duration flag is required (months: 1/3/6/12)")
			}
			req := &rcdn.CreateRcdnInstanceRequest{
				PlanID:       planID,
				Duration:     duration,
				WithCouponID: coupon,
				Config:       &rcdn.RcdnBaseConfig{},
			}
			if domains != "" {
				req.Domains = strings.Split(domains, ",")
			}
			if _, err := (*rySDK).CreateRcdnInstance(req); err != nil {
				return err
			}
			fmt.Println("RCDN instance creating")
			return nil
		},
	}

	renewCmd := &cobra.Command{
		Use:   "renew <id>",
		Short: "Renew an RCDN instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if duration == 0 {
				return fmt.Errorf("--duration flag is required")
			}
			if _, err := (*rySDK).RenewRcdnInstance(id, rcdn.RenewRcdnInstanceRequest{
				Duration: duration, WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Printf("RCDN instance %s renewed\n", args[0])
			return nil
		},
	}

	scaleCmd := &cobra.Command{
		Use:   "scale <id>",
		Short: "Scale an RCDN instance to a new plan",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			destPlan, _ := cmd.Flags().GetInt("dest-plan")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if destPlan == 0 {
				return fmt.Errorf("--dest-plan flag is required")
			}
			if _, err := (*rySDK).ScaleRcdnInstance(id, rcdn.ScaleRcdnInstanceRequest{
				DestPlan: destPlan, WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Printf("RCDN instance %s scaling\n", args[0])
			return nil
		},
	}

	domainCmd := &cobra.Command{Use: "domain", Short: "Manage RCDN accelerated domains"}

	domainListCmd := &cobra.Command{
		Use:   "list",
		Short: "List accelerated domains",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcdnDomainList(&rcdn.GetRcdnDomainListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			items := make([]model.RcdnDomain, len(resp.Data.Records))
			for i, d := range resp.Data.Records {
				items[i] = toRcdnDomain(d)
			}
			return (*out).Print(items)
		},
	}

	domainAddCmd := &cobra.Command{
		Use:   "add <domain>",
		Short: "Add an accelerated domain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			instanceID, _ := cmd.Flags().GetInt("instance-id")
			copyFrom, _ := cmd.Flags().GetString("copy-from")
			if instanceID == 0 {
				return fmt.Errorf("--instance-id flag is required")
			}
			if _, err := (*rySDK).AddRcdnDomain(&rcdn.AddRcdnDomainRequest{
				Domain: args[0], InstanceID: instanceID, CopyFromDomain: copyFrom,
			}); err != nil {
				return err
			}
			fmt.Printf("Domain %s added\n", args[0])
			return nil
		},
	}

	domainGetCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get accelerated domain details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcdnDomainDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toRcdnDomain(resp.Data))
		},
	}

	domainDeleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete an accelerated domain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRcdnDomain(id); err != nil {
				return err
			}
			fmt.Printf("Domain %s deleted\n", args[0])
			return nil
		},
	}

	createCmd.Flags().Int("plan-id", 0, "Plan ID (required)")
	createCmd.Flags().Int("duration", 0, "Duration in months: 1/3/6/12 (required)")
	createCmd.Flags().String("domains", "", "Comma-separated initial domains")
	createCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewCmd.Flags().Int("duration", 0, "Duration in months: 1/3/6/12 (required)")
	renewCmd.Flags().Int("coupon", 0, "Coupon ID")
	scaleCmd.Flags().Int("dest-plan", 0, "Destination plan ID (required)")
	scaleCmd.Flags().Int("coupon", 0, "Coupon ID")
	domainAddCmd.Flags().Int("instance-id", 0, "RCDN instance ID (required)")
	domainAddCmd.Flags().String("copy-from", "", "Copy config from an existing domain")

	rcdnCmd.AddCommand(listCmd, getCmd, createCmd, renewCmd, scaleCmd)
	domainCmd.AddCommand(domainListCmd, domainAddCmd, domainGetCmd, domainDeleteCmd)
	rcdnCmd.AddCommand(domainCmd)

	return rcdnCmd
}

func toRcdnInstance(r rcdn.RcdnInstance) model.RcdnInstance {
	return model.RcdnInstance{
		ID:          strconv.Itoa(r.ID),
		Status:      r.Status,
		Plan:        r.Plan.Chinese,
		Tag:         r.Tag,
		Region:      r.Node.Region,
		TrafficUsed: r.UsageTraffic,
		AutoRenew:   r.AutoRenew,
		CreatedAt:   time.Unix(int64(r.CreateDate), 0),
		ExpireAt:    time.Unix(int64(r.ExpDate), 0),
	}
}

func toRcdnDomain(d rcdn.RcdnDomain) model.RcdnDomain {
	return model.RcdnDomain{
		ID:     strconv.Itoa(d.ID),
		Domain: d.Domain,
		CNAME:  d.CNAME,
		Region: d.Region,
	}
}
