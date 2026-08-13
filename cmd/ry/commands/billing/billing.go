package billing

import (
	"strconv"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/expense"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	billingCmd := &cobra.Command{
		Use:   "billing",
		Short: "Billing and orders",
	}

	ordersCmd := &cobra.Command{
		Use:   "orders",
		Short: "List orders",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetOrdersList(&expense.GetOrdersListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			orders := make([]model.Order, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				t, _ := time.Parse("2006-01-02 15:04:05", r.CreateDate)
				orders[i] = model.Order{
					ID:      strconv.Itoa(r.ID),
					Product: r.ProductName,
					Amount:  r.Amount,
					Status:  r.Status,
					Created: t,
				}
			}
			return (*out).Print(orders)
		},
	}

	billingCmd.AddCommand(ordersCmd)

	return billingCmd
}
