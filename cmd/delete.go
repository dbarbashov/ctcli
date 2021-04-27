package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "delete a release",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented")
	},
}

func init()  {
	rootCmd.AddCommand(deleteCmd)
}