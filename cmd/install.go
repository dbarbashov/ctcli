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

var installCmd = &cobra.Command{
	Use: "install <path to package>",
	Short: "install a release",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packagePath, err := filepath.Abs(args[0])
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		rootFlag := cmd.Flag("root")
		rootDir, err := filepath.Abs(rootFlag.Value.String())
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		fn := util.MirrorStdoutToFile(ctcliDir.GetCtcliLogFilePath(rootDir))
		defer fn()

		err = domain.Install(rootDir, packagePath)
		if err != nil {
			log.Fatal(err)
		}
		color.HiGreen("OK\n")
	},
}

func init()  {
	rootCmd.AddCommand(installCmd)
}
