package cmd

import (
	"github.com/spf13/cobra"
	"github.com/leo3838ok/go-sharing-01/utils/errors"
)

var rootCmd = &cobra.Command{
	Use:           "gocli",
	Short:         "A CLI tool",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	if cmd, err := rootCmd.ExecuteC(); err != nil {
		errors.Handle(cmd, err)
	}
}