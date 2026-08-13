package main

import (
	"fmt"
	"os"

	"github.com/XingfenD/rainyun_api_go_sdk/apis"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/configcmd"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/domain"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/public"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/server"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/commands/storage"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/config"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/constant"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/trace"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

var (
	cfg                  *config.Config
	cfgPath              string
	rySDK                *sdk.RainyunSDK
	out                  *output.Printer
	flagAPIKey           string
	flagOutput           string
	flagVerbose          bool
	flagVerboseBodyLimit int
	flagVerboseFullBody  bool
)

var rootCmd = &cobra.Command{
	Use:     "ry",
	Short:   "Rainyun CLI — manage cloud resources from your terminal",
	Version: constant.Version,
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
		if flagVerboseBodyLimit < 0 {
			return fmt.Errorf("--verbose-body-limit must not be negative, got %d", flagVerboseBodyLimit)
		}
		rySDK = sdk.NewBuilder(apiKey).
			WithTrace(verboseTraceOptions(flagVerbose, flagVerboseBodyLimit, flagVerboseFullBody, trace.NewVerboseRenderer(os.Stderr))).
			Build()

		outputFormat, outputSource := resolveOutput(cfg.Output, flagOutput)
		out = output.New(outputFormat, os.Stdout)
		out.SetRawBody(rySDK.RawResponseBody)
		if flagVerbose {
			fmt.Fprintf(os.Stderr, "[debug] config: %s\n", cfgPath)
			fmt.Fprintf(os.Stderr, "[debug] output: format=%s source=%s\n", outputFormat, outputSource)
		}
		return nil
	},
}

func verboseTraceOptions(enabled bool, limit int, full bool, sink apis.TraceSink) *apis.TraceOptions {
	if !enabled {
		return nil
	}
	opts := apis.NewTraceOptions(sink)
	switch {
	case full:
		opts = opts.WithFullBodyPreview()
	case limit == 0:
		opts = opts.WithoutBodyPreview()
	case limit > 0:
		opts = opts.WithBodyPreviewLimit(limit)
	}
	return &opts
}

func resolveOutput(configFormat, flagFormat string) (format, source string) {
	if flagFormat != "" {
		return flagFormat, "--output"
	}
	return configFormat, "config"
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	cfgPath = config.DefaultPath()
	rootCmd.PersistentFlags().StringVar(&flagAPIKey, "apikey", "", "Rainyun API key")
	rootCmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "", "Output format: table, json, yaml, raw")
	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().IntVar(&flagVerboseBodyLimit, "verbose-body-limit", 64*1024, "Maximum response body bytes shown in verbose output (0 disables the body)")
	rootCmd.PersistentFlags().BoolVar(&flagVerboseFullBody, "verbose-full-body", false, "Show the full response body in verbose output")

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

	versionCmd := &cobra.Command{
		Use:         "version",
		Short:       "Print the CLI version",
		Annotations: map[string]string{"skipSDK": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(constant.Version)
			return nil
		},
	}
	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(configcmd.Cmd(&cfg, cfgPath))
	rootCmd.AddCommand(server.Cmd(&rySDK, &out))
	rootCmd.AddCommand(storage.Cmd(&rySDK, &out))
	rootCmd.AddCommand(domain.Cmd(&rySDK, &out))
	rootCmd.AddCommand(public.Cmd(&rySDK, &out))
}
