package cmd

import (
	"github.com/docker/cli/cli"
	"github.com/docker/lunchbox/packager"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load <repotag>",
	Short: "Load an application from a docker image",
	Args:  cli.RequiresMaxArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return packager.Load(firstOrEmpty(args), ".")
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
