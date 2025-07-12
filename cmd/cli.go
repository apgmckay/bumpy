package cmd

import (
	"bumpy/package/server"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var rootCmd = BumpyRootCmd()

func init() {
	rootCmd.AddCommand(bumpyServerCmd)
}

func Execute() error {
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		return err
	}
	return nil
}

func BumpyRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bumpy",
		Long:  "Helps you bump you versions using semantic versioning.\n\nPass a version number and bumpy returns a bumped version number.",
		Short: "Helps you bump you versions using semantic versioning",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
}

var bumpyServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start running the bumpy web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.New().Run()
		return nil
	},
}
