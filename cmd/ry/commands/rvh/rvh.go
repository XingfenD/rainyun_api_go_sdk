// Package rvh implements the ry rvh command.
package rvh

import (
	"fmt"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rvh"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	rvhCmd := &cobra.Command{Use: "rvh", Short: "Manage Rainyun virtual hosts"}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List virtual hosts",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRvhList(&rvh.GetRvhListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	getCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get virtual host details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRvhDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	plansCmd := &cobra.Command{
		Use:   "plans",
		Short: "List virtual host plans",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRvhPlanList()
			if err != nil {
				return err
			}
			items := make([]model.RvhPlan, len(resp.Data))
			for i, p := range resp.Data {
				items[i] = toRvhPlan(p)
			}
			return (*out).Print(items)
		},
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a virtual host",
		RunE: func(cmd *cobra.Command, args []string) error {
			planID, _ := cmd.Flags().GetInt("plan-id")
			duration, _ := cmd.Flags().GetInt("duration")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if planID == 0 || duration == 0 {
				return fmt.Errorf("--plan-id and --duration flags are required")
			}
			if _, err := (*rySDK).CreateRvh(&rvh.CreateRvhRequest{
				PlanID: planID, Duration: duration, WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Println("Virtual host creating")
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Release a virtual host",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).FreeRvh(id); err != nil {
				return err
			}
			fmt.Printf("Virtual host %s released\n", args[0])
			return nil
		},
	}

	renewCmd := &cobra.Command{
		Use:   "renew <id>",
		Short: "Renew a virtual host",
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
			if _, err := (*rySDK).RenewRvh(id, rvh.RenewRvhRequest{
				Duration: duration, WithCouponID: coupon,
			}); err != nil {
				return err
			}
			fmt.Printf("Virtual host %s renewed\n", args[0])
			return nil
		},
	}

	backupCmd := &cobra.Command{Use: "backup", Short: "Manage virtual host backups"}

	backupCreateCmd := &cobra.Command{
		Use:   "create <id> <label>",
		Short: "Create a backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CreateRvhBackup(id, rvh.CreateRvhBackupRequest{Label: args[1]}); err != nil {
				return err
			}
			fmt.Printf("Backup %q creating\n", args[1])
			return nil
		},
	}

	backupDeleteCmd := &cobra.Command{
		Use:   "delete <id> <backup-id>",
		Short: "Delete a backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			bid, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRvhBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %s deleted\n", args[1])
			return nil
		},
	}

	backupRestoreCmd := &cobra.Command{
		Use:   "restore <id> <backup-id>",
		Short: "Restore a backup",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			bid, err := cliutil.ParseID(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).RestoreRvhBackup(id, bid); err != nil {
				return err
			}
			fmt.Printf("Backup %s restoring\n", args[1])
			return nil
		},
	}

	domainCmd := &cobra.Command{Use: "domain", Short: "Manage virtual host domains"}

	domainBindCmd := &cobra.Command{
		Use:   "bind <id> <domain>",
		Short: "Bind a domain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			sslCertID, _ := cmd.Flags().GetInt("ssl-cert-id")
			sslForce, _ := cmd.Flags().GetInt("ssl-force")
			if _, err := (*rySDK).BindRvhDomain(id, rvh.BindRvhDomainRequest{
				Domain: args[1], SSLCertID: sslCertID, SSLForce: sslForce,
			}); err != nil {
				return err
			}
			fmt.Printf("Domain %s bound\n", args[1])
			return nil
		},
	}

	domainUnbindCmd := &cobra.Command{
		Use:   "unbind <id> <domain>",
		Short: "Unbind a domain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).UnbindRvhDomain(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Domain %s unbound\n", args[1])
			return nil
		},
	}

	createCmd.Flags().Int("plan-id", 0, "Plan ID (required, see 'ry rvh plans')")
	createCmd.Flags().Int("duration", 0, "Duration in months: 1/3/6/12 (required)")
	createCmd.Flags().Int("coupon", 0, "Coupon ID")
	renewCmd.Flags().Int("duration", 0, "Duration in months: 1/3/6/12 (required)")
	renewCmd.Flags().Int("coupon", 0, "Coupon ID")
	domainBindCmd.Flags().Int("ssl-cert-id", -1, "SSL certificate ID (-1 for http)")
	domainBindCmd.Flags().Int("ssl-force", 0, "SSL force redirect: 0/1")

	backupCmd.AddCommand(backupCreateCmd, backupDeleteCmd, backupRestoreCmd)
	domainCmd.AddCommand(domainBindCmd, domainUnbindCmd)
	rvhCmd.AddCommand(listCmd, getCmd, plansCmd, createCmd, deleteCmd, renewCmd, backupCmd, domainCmd)

	return rvhCmd
}

func toRvhPlan(p rvh.RvhPlan) model.RvhPlan {
	return model.RvhPlan{
		ID:      p.ID,
		Plan:    p.PlanName,
		Chinese: p.Chinese,
		Price:   p.Price,
		Disk:    p.Disk,
	}
}
