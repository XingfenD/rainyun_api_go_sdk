package game

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addPlansCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	plansCmd := &cobra.Command{
		Use:   "plans",
		Short: "List game server plans",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRgsPlanList()
			if err != nil {
				return err
			}
			plans := make([]model.GamePlan, 0, len(resp.Data))
			for _, p := range resp.Data {
				plans = append(plans, toGamePlan(p))
			}
			return (*out).Print(plans)
		},
	}

	discountCmd := &cobra.Command{
		Use:   "discount-percent",
		Short: "Get game server discount percent",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRgsDiscountPercent(&rgs.GetRgsDiscountPercentRequest{})
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	osCmd := &cobra.Command{
		Use:   "os",
		Short: "List game server OS templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRgsOSList()
			if err != nil {
				return err
			}
			return (*out).Print(toGameOS(resp.Data))
		},
	}

	gameCmd.AddCommand(plansCmd)
	gameCmd.AddCommand(discountCmd)
	gameCmd.AddCommand(osCmd)
}

func toGamePlan(p rgs.RgsPlan) model.GamePlan {
	return model.GamePlan{
		ID:          p.ID,
		Region:      p.Region,
		Subtype:     p.Subtype,
		Plan:        p.PlanName,
		Machine:     p.Machine,
		ChargeType:  p.ChargeType,
		Chinese:     p.Chinese,
		Selling:     p.IsSelling,
		CPUPrice:    p.CPUPrice,
		MemPrice:    p.MemoryPrice,
		NetOutPrice: p.NetOutPrice,
	}
}

func toGameOS(items []public.RgsOSListItem) []model.GameOS {
	oses := make([]model.GameOS, 0, len(items))
	for _, o := range items {
		oses = append(oses, model.GameOS{
			ID:        o.ID,
			Name:      o.Name,
			Chinese:   o.ChineseName,
			Region:    o.Region,
			Subtype:   o.Subtype,
			Version:   o.Version,
			Available: o.IsAvailable,
		})
	}
	return oses
}
