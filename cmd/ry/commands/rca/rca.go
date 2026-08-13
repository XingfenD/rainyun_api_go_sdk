// Package rca implements the ry rca command.
package rca

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/rca"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	rcaCmd := &cobra.Command{Use: "rca", Short: "Manage Rainyun cloud apps (RCA)"}

	raindropCmd := &cobra.Command{Use: "raindrop", Short: "Raindrop balance and plans"}

	raindropBalanceCmd := &cobra.Command{
		Use:   "balance",
		Short: "Get raindrop balance",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcaRaindropBalance()
			if err != nil {
				return err
			}
			return (*out).Print(model.RcaRaindropBalance{Balance: resp.Data})
		},
	}

	raindropBuyCmd := &cobra.Command{
		Use:   "buy",
		Short: "Buy raindrops",
		RunE: func(cmd *cobra.Command, args []string) error {
			planID, _ := cmd.Flags().GetInt("plan-id")
			coupon, _ := cmd.Flags().GetInt("coupon")
			if planID == 0 {
				return fmt.Errorf("--plan-id flag is required (see 'ry rca raindrop plans')")
			}
			if _, err := (*rySDK).BuyRaindrop(planID, coupon); err != nil {
				return err
			}
			fmt.Printf("Raindrops bought (plan %d)\n", planID)
			return nil
		},
	}

	raindropLogCmd := &cobra.Command{
		Use:   "log",
		Short: "Get raindrop consume history",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRaindropConsumeLog(`{"page":1,"perPage":50}`)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	raindropPlansCmd := &cobra.Command{
		Use:   "plans",
		Short: "List raindrop plans",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcaRaindropPlansList()
			if err != nil {
				return err
			}
			items := make([]model.RcaRaindropPlan, len(resp.Data))
			for i, p := range resp.Data {
				items[i] = model.RcaRaindropPlan{ID: p.ID, PlanName: p.PlanName,
					Chinese: p.Chinese, Amount: p.Amount, Price: p.Price}
			}
			return (*out).Print(items)
		},
	}

	regionCmd := &cobra.Command{
		Use:   "regions",
		Short: "List RCA regions",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcaRegionInfo()
			if err != nil {
				return err
			}
			items := make([]model.RcaRegion, len(resp.Data))
			for i, r := range resp.Data {
				items[i] = model.RcaRegion{ID: r.ID, Name: r.Name, ChineseName: r.ChineseName}
			}
			return (*out).Print(items)
		},
	}

	projectCmd := &cobra.Command{Use: "project", Short: "Manage RCA projects"}

	projectListCmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).ListRcaProjects(`{"page":1,"perPage":50}`)
			if err != nil {
				return err
			}
			items := make([]model.RcaProject, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				items[i] = toRcaProject(r)
			}
			return (*out).Print(items)
		},
	}

	projectGetCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get project details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcaProjectDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toRcaProject(resp.Data.Data))
		},
	}

	projectCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			regionID, _ := cmd.Flags().GetInt("region-id")
			diskSize, _ := cmd.Flags().GetInt("disk-size")
			if name == "" || regionID == 0 || diskSize == 0 {
				return fmt.Errorf("--name, --region-id and --disk-size flags are required")
			}
			if _, err := (*rySDK).CreateRcaProject(&rca.CreateRcaProjectRequest{
				Name: name, RegionID: regionID, DiskSize: diskSize, ChargeType: "elastic",
			}); err != nil {
				return err
			}
			fmt.Printf("Project %q creating\n", name)
			return nil
		},
	}

	projectDeleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Destroy a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DestroyRcaProject(id); err != nil {
				return err
			}
			fmt.Printf("Project %s destroyed\n", args[0])
			return nil
		},
	}

	appCmd := &cobra.Command{Use: "app", Short: "Manage RCA apps"}

	appListCmd := &cobra.Command{
		Use:   "list <project-id>",
		Short: "List apps of a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcaAppList(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	appGetCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get app details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcaAppDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	appCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Install an app",
		RunE: func(cmd *cobra.Command, args []string) error {
			releaseID, _ := cmd.Flags().GetInt("appstore-release-id")
			projectID, _ := cmd.Flags().GetInt("project-id")
			optionsJSON, _ := cmd.Flags().GetString("options")
			if releaseID == 0 || projectID == 0 {
				return fmt.Errorf("--appstore-release-id and --project-id flags are required")
			}
			req := &rca.InstallRcaAppRequest{
				AppstoreReleaseID: releaseID, ProjectID: projectID,
			}
			if optionsJSON != "" {
				req.Options = json.RawMessage(optionsJSON)
			}
			if _, err := (*rySDK).InstallRcaApp(req); err != nil {
				return err
			}
			fmt.Println("App installing")
			return nil
		},
	}

	appDeleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Uninstall an app",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			removeData, _ := cmd.Flags().GetBool("remove-data")
			if _, err := (*rySDK).UninstallRcaApp(id, &rca.UninstallRcaAppRequest{RemoveData: removeData}); err != nil {
				return err
			}
			fmt.Printf("App %s uninstalled\n", args[0])
			return nil
		},
	}

	lifecycle := func(action string) func(*cobra.Command, []string) error {
		return func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			var err2 error
			switch action {
			case "start":
				_, err2 = (*rySDK).StartRcaApp(id)
			case "stop":
				_, err2 = (*rySDK).StopRcaApp(id)
			case "restart":
				_, err2 = (*rySDK).RestartRcaApp(id)
			}
			if err2 != nil {
				return err2
			}
			fmt.Printf("App %s %sed\n", args[0], action)
			return nil
		}
	}

	appStartCmd := &cobra.Command{Use: "start <id>", Short: "Start an app", Args: cobra.ExactArgs(1), RunE: lifecycle("start")}
	appStopCmd := &cobra.Command{Use: "stop <id>", Short: "Stop an app", Args: cobra.ExactArgs(1), RunE: lifecycle("stop")}
	appRestartCmd := &cobra.Command{Use: "restart <id>", Short: "Restart an app", Args: cobra.ExactArgs(1), RunE: lifecycle("restart")}

	websiteCmd := &cobra.Command{Use: "website", Short: "Manage RCA websites"}

	websiteListCmd := &cobra.Command{
		Use:   "list",
		Short: "List websites",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRcaWebsiteList()
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	websiteGetCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get website details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetRcaWebsiteDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	websiteCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a website",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			websiteType, _ := cmd.Flags().GetString("type")
			projectID, _ := cmd.Flags().GetInt("project-id")
			domains, _ := cmd.Flags().GetString("domains")
			if name == "" || websiteType == "" || projectID == 0 || domains == "" {
				return fmt.Errorf("--name, --type, --project-id and --domains flags are required")
			}
			if _, err := (*rySDK).CreateRcaWebsite(&rca.CreateRcaWebsiteRequest{
				Name: name, Type: websiteType, ProjectID: projectID,
				Domains: strings.Split(domains, ","),
			}); err != nil {
				return err
			}
			fmt.Printf("Website %q creating\n", name)
			return nil
		},
	}

	websiteDeleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete a website",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteRcaWebsite(id); err != nil {
				return err
			}
			fmt.Printf("Website %s deleted\n", args[0])
			return nil
		},
	}

	raindropBuyCmd.Flags().Int("plan-id", 0, "Raindrop plan ID (required)")
	raindropBuyCmd.Flags().Int("coupon", 0, "Coupon ID")
	projectCreateCmd.Flags().String("name", "", "Project name (required)")
	projectCreateCmd.Flags().Int("region-id", 0, "Region ID (required, see 'ry rca regions')")
	projectCreateCmd.Flags().Int("disk-size", 0, "Disk size in GiB (required)")
	appCreateCmd.Flags().Int("appstore-release-id", 0, "Appstore release ID (required)")
	appCreateCmd.Flags().Int("project-id", 0, "Project ID (required)")
	appCreateCmd.Flags().String("options", "", "App options as JSON, e.g. '{\"root_pass\":\"x\"}'")
	appDeleteCmd.Flags().Bool("remove-data", false, "Also remove persistent data")
	websiteCreateCmd.Flags().String("name", "", "Website name (required)")
	websiteCreateCmd.Flags().String("type", "", "Website type (required)")
	websiteCreateCmd.Flags().Int("project-id", 0, "Project ID (required)")
	websiteCreateCmd.Flags().String("domains", "", "Comma-separated domains (required)")

	raindropCmd.AddCommand(raindropBalanceCmd, raindropBuyCmd, raindropLogCmd, raindropPlansCmd)
	projectCmd.AddCommand(projectListCmd, projectGetCmd, projectCreateCmd, projectDeleteCmd)
	appCmd.AddCommand(appListCmd, appGetCmd, appCreateCmd, appDeleteCmd,
		appStartCmd, appStopCmd, appRestartCmd)
	websiteCmd.AddCommand(websiteListCmd, websiteGetCmd, websiteCreateCmd, websiteDeleteCmd)
	rcaCmd.AddCommand(raindropCmd, regionCmd, projectCmd, appCmd, websiteCmd)

	return rcaCmd
}

func toRcaProject(p rca.RcaProject) model.RcaProject {
	return model.RcaProject{
		ID:     strconv.Itoa(p.ID),
		Name:   p.Name,
		Status: p.Status,
		Region: p.Region.ChineseName,
		MaxCPU: p.ResourceLimits.MaxCPU,
		MaxMem: p.ResourceLimits.MaxMemory,
	}
}
