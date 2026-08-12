package main

import (
	"fmt"
	"os"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/billing"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/configcmd"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/domain"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/server"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/storage"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/config"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

var (
	cfg         *config.Config
	cfgPath     string
	rySDK       *sdk.RainyunSDK
	out         *output.Printer
	flagAPIKey  string
	flagOutput  string
	flagRaw     bool
	flagVerbose bool
)

var rootCmd = &cobra.Command{
	Use:   "ry",
	Short: "Rainyun CLI — manage cloud resources from your terminal",
	Long: `ry is a CLI that lets you manage Rainyun cloud resources from your terminal.

Use "ry [command] --help" for more information about a command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.Load(cfgPath)
		if err != nil {
			return fmt.Errorf("load config: %w", err)
		}

		if cmd.Annotations != nil && cmd.Annotations["skipSDK"] == "true" {
			return nil
		}

		apiKey := cfg.APIKey
		if flagAPIKey != "" {
			apiKey = flagAPIKey
		}
		if env := os.Getenv("RY_APIKEY"); env != "" {
			apiKey = env
		}
		if apiKey == "" {
			return fmt.Errorf("no api key configured — run 'ry config set apikey <key>'")
		}
		rySDK = sdk.New(apiKey)

		outputFormat := cfg.Output
		if flagOutput != "" {
			outputFormat = flagOutput
		}
		if flagRaw {
			outputFormat = "raw"
		}
		out = output.New(outputFormat, os.Stdout)
		if flagVerbose {
			fmt.Fprintf(os.Stderr, "[debug] config: %s\n", cfgPath)
			fmt.Fprintf(os.Stderr, "[debug] output: %s\n", outputFormat)
		}
		return nil
	},
}

func init() {
	cfgPath = config.DefaultPath()
	rootCmd.PersistentFlags().StringVar(&flagAPIKey, "apikey", "", "Rainyun API key")
	rootCmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "", "Output format: table, json, yaml, raw")
	rootCmd.PersistentFlags().BoolVar(&flagRaw, "raw", false, "Output raw API response")
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Enable verbose output")

	completionCmd := &cobra.Command{
		Use:         "completion [bash|zsh|fish]",
		Short:       "Generate shell completion script",
		Annotations: map[string]string{"skipSDK": "true"},
		Args:        cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return rootCmd.GenBashCompletion(os.Stdout)
			case "zsh":
				return rootCmd.GenZshCompletion(os.Stdout)
			case "fish":
				return rootCmd.GenFishCompletion(os.Stdout, true)
			default:
				return fmt.Errorf("unsupported shell: %s", args[0])
			}
		},
	}
	rootCmd.AddCommand(completionCmd)

	rootCmd.AddCommand(configcmd.Cmd(&cfg, cfgPath))
	rootCmd.AddCommand(server.Cmd(&rySDK, &out, &flagRaw))
	rootCmd.AddCommand(storage.Cmd(&rySDK, &out))
	rootCmd.AddCommand(domain.Cmd(&rySDK, &out))
	rootCmd.AddCommand(billing.Cmd(&rySDK, &out))
}
