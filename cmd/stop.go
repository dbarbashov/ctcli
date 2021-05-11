package cmd

import (
	"ctcli/domain"
	"ctcli/domain/ctcliDir"
	"github.com/spf13/cobra"
	"path/filepath"
)

var stopCmd = &cobra.Command{
	Use:   "stop [app]",
	Short: "stops a service",
	Run: func(cmd *cobra.Command, args []string) {
		rootFlag := cmd.Flag("root")
		rootDir, err := filepath.Abs(rootFlag.Value.String())
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		if err := ctcliDir.OkIfIsARootDir(rootDir); err != nil {
			cmd.PrintErr(err)
			return
		}
		domain.StopApps(rootDir)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
