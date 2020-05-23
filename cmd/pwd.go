package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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