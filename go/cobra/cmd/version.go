package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "version",
		Short: "Print the version or upgrade",

		Run: func(cmd *cobra.Command, args []string) {
			currentVersion := "v1.0.0"

			fmt.Printf("Current version: %s\n", currentVersion)

			sdk, err := NewSDK(currentVersion)
			if err != nil {
				panic(err)
			}

			updateInfo, err := sdk.GetUpdateInfo()
			if err != nil {
				return // you should determine how to handle this error
			}

			if updateInfo != nil {
				fmt.Printf("Latest version: %s\n", updateInfo.LatestVersion)

				canUpgrade, err := sdk.CanSupportUpgrade()
				if err != nil {
					return // you should determine how to handle this error
				}

				if canUpgrade {
					fmt.Printf("To upgrade, run \"%s version upgrade\"\n\n", os.Args[0])
				} else {
					fmt.Printf("Up upgrade, run \"%s\"\n\n", sdk.ExternalUpgradeCommand())
				}
			}
		},
	}

	cmd.AddCommand(upgradeCmd())

	return &cmd
}

func upgradeCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "upgrade",
		Short: "Install any pending upgrade to this CLI",

		Run: func(cmd *cobra.Command, args []string) {
			currentVersion := "v1.0.0"

			sdk, err := NewSDK(currentVersion)
			if err != nil {
				panic(err)
			}

			if err := sdk.Upgrade(); err != nil {
				// our example repo doesn't have the right assets currently
				// so this will error
				return
				// panic(err)
			}
		},
	}

	return &cmd
}
