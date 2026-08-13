package public

import (
	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	publicCmd := &cobra.Command{
		Use:   "public",
		Short: "Browse public Rainyun info (news, status, app config)",
	}

	publicCmd.AddCommand(appConfigCmd(rySDK, out))
	publicCmd.AddCommand(newsCmd(rySDK, out))
	publicCmd.AddCommand(statusCmd(rySDK, out))

	return publicCmd
}

func appConfigCmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	return &cobra.Command{
		Use:   "app-config",
		Short: "List app config entries",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetAppConfig()
			if err != nil {
				return err
			}
			var items []model.AppConfig
			for _, c := range resp.Data {
				for _, v := range c.Value {
					items = append(items, model.AppConfig{
						Type:    c.Type,
						Title:   v.Title,
						Name:    v.Name,
						Page:    v.Page,
						Order:   v.Order,
						Content: v.Content,
					})
				}
			}
			return (*out).Print(items)
		},
	}
}

func newsCmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	return &cobra.Command{
		Use:   "news",
		Short: "List news announcements",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetNews()
			if err != nil {
				return err
			}
			items := make([]model.News, len(resp.Data))
			for i, n := range resp.Data {
				items[i] = model.News{Type: n.Type, Title: n.Title, Time: n.TimeStamp, URL: n.URL}
			}
			return (*out).Print(items)
		},
	}
}

func statusCmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "List node network status",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetStatus(&public.GetStatusRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			records := make([]model.NodeStatus, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				records[i] = model.NodeStatus{
					UUID:    r.UUID,
					Name:    r.ChineseName,
					Product: r.Product,
					CPU:     r.CPU,
					Memory:  r.Memory,
					NetOut:  r.NetOut,
					Status:  r.Status,
					Updated: r.UpdateTime,
				}
			}
			return (*out).Print(records)
		},
	}
}
