package configcmd

import (
	"fmt"
	"strings"

	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/config"

	"github.com/spf13/cobra"
)

func Cmd(cfg **config.Config, cfgPath string) *cobra.Command {
	configCmd := &cobra.Command{
		Use:         "config",
		Short:       "Manage ry configuration",
		Annotations: map[string]string{"skipSDK": "true"},
	}

	configShowCmd := &cobra.Command{
		Use:         "show",
		Short:       "Show current configuration",
		Annotations: map[string]string{"skipSDK": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			c := *cfg
			fmt.Printf("Config file: %s\n", cfgPath)
			masked := c.APIKey
			if len(masked) > 4 {
				masked = masked[:2] + strings.Repeat("*", len(masked)-4) + masked[len(masked)-2:]
			}
			fmt.Printf("API key:     %s\n", masked)
			fmt.Printf("Output:      %s\n", c.Output)
			return nil
		},
	}

	configSetCmd := &cobra.Command{
		Use:         "set <key> [value]",
		Short:       "Set a config value",
		Args:        cobra.RangeArgs(1, 2),
		Annotations: map[string]string{"skipSDK": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			var value string

			if len(args) == 2 {
				value = args[1]
			} else {
				// Interactive mode — prompt for the value.
				if key == "apikey" {
					v, err := config.ReadSecret("Enter API key: ")
					if err != nil {
						return fmt.Errorf("read API key: %w", err)
					}
					value = v
				} else {
					v, err := config.ReadLine("Enter " + key + ": ")
					if err != nil {
						return fmt.Errorf("read %s: %w", key, err)
					}
					value = v
				}
			}

			c := *cfg
			switch key {
			case "apikey":
				c.APIKey = value
			case "output":
				c.Output = value
			default:
				return fmt.Errorf("unknown config key: %s (supported: apikey, output)", key)
			}
			if err := config.Save(cfgPath, c); err != nil {
				return fmt.Errorf("save config: %w", err)
			}
			if key == "apikey" {
				fmt.Printf("Set %s = %s\n", key, "(hidden)")
			} else {
				fmt.Printf("Set %s = %s\n", key, value)
			}
			return nil
		},
	}

	configPathCmd := &cobra.Command{
		Use:         "path",
		Short:       "Print config file path",
		Annotations: map[string]string{"skipSDK": "true"},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(cfgPath)
			return nil
		},
	}

	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configPathCmd)

	return configCmd
}
