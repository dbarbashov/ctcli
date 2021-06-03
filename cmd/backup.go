package cmd

import (
	"ctcli/domain"
	"ctcli/domain/ctcliDir"
	"ctcli/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "make a backup of current release",
	Run: func(cmd *cobra.Command, args []string) {
		rootFlag := cmd.Flag("root")
		rootDir, err := filepath.Abs(rootFlag.Value.String())
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		fn := util.MirrorStdoutToFile(ctcliDir.GetCtcliLogFilePath(rootDir))
		defer fn()
		if err := domain.MakeABackup(rootDir); err != nil {
			log.Fatal(err)
		}
		color.Green("backup was made\n")
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
}
