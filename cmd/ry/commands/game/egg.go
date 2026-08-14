package game

import (
	"fmt"
	"strings"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/public"
	"github.com/XingfenD/rainyun_api_go_sdk/apis/rgs"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func addEggCommands(gameCmd *cobra.Command, rySDK **sdk.RainyunSDK, out **output.Printer) {
	eggCmd := &cobra.Command{
		Use:   "egg",
		Short: "Browse game types and servers",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List games (eggs)",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetEggList()
			if err != nil {
				return err
			}
			return (*out).Print(toGameEggs(resp.Data))
		},
	}

	typeCmd := &cobra.Command{
		Use:   "type",
		Short: "List game types (egg types)",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetEggTypeList()
			if err != nil {
				return err
			}
			return (*out).Print(toGameEggTypes(resp.Data))
		},
	}

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "List game server types",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetRgsEggServerList()
			if err != nil {
				return err
			}
			servers := make([]model.GameEggServer, 0, len(resp.Data))
			for _, s := range resp.Data {
				servers = append(servers, toGameEggServer(s))
			}
			return (*out).Print(servers)
		},
	}

	changeCmd := &cobra.Command{
		Use:   "change <id>",
		Short: "Change game type (egg) of a game server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			eggTypeID, _ := cmd.Flags().GetInt("egg-type-id")
			if eggTypeID == 0 {
				return fmt.Errorf("--egg-type-id flag is required")
			}
			saveDirsStr, _ := cmd.Flags().GetString("save-dirs")
			var saveDirs []string
			if saveDirsStr != "" {
				saveDirs = strings.Split(saveDirsStr, ",")
			}
			if _, err := (*rySDK).ChangeRgsEgg(id, eggTypeID, saveDirs); err != nil {
				return err
			}
			fmt.Printf("Game type change requested for game server %d\n", id)
			return nil
		},
	}
	changeCmd.Flags().Int("egg-type-id", 0, "Target game type ID (required)")
	changeCmd.Flags().String("save-dirs", "", "Comma-separated dirs to keep")

	eggCmd.AddCommand(listCmd)
	eggCmd.AddCommand(typeCmd)
	eggCmd.AddCommand(serverCmd)
	eggCmd.AddCommand(changeCmd)

	gameCmd.AddCommand(eggCmd)
}

func toGameEggServer(s rgs.RgsEggServer) model.GameEggServer {
	return model.GameEggServer{
		Server:  s.Server,
		EggName: s.EggName,
		Desc:    s.Desc,
		Order:   s.Order,
	}
}

func toGameEggs(items []public.EggItem) []model.GameEgg {
	eggs := make([]model.GameEgg, 0, len(items))
	for _, e := range items {
		eggs = append(eggs, model.GameEgg{
			Name: e.Name, Title: e.Title, EggGroup: e.EggGroup, Desc: e.Desc, Order: e.Order,
		})
	}
	return eggs
}

func toGameEggTypes(items []public.EggTypeItem) []model.GameEggType {
	types := make([]model.GameEggType, 0, len(items))
	for _, t := range items {
		types = append(types, model.GameEggType{
			ID:            t.ID,
			Name:          t.EggName,
			Game:          t.Egg.Name,
			ServerType:    t.Env.ServerType,
			ServerVersion: t.Env.ServerVersion,
			Order:         t.Order,
		})
	}
	return types
}
