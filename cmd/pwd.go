package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var pwdCmd = &cobra.Command{
	Use:                   "pwd",
	Short:                 "Return working directory name",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		fmt.Println(pwd)
	},
}

func init() {
	rootCmd.AddCommand(pwdCmd)
}