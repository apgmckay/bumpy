package cmd

import (
	"bumpy/package/client"
	"bumpy/package/server"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var rootCmd = BumpyRootCmd()

func init() {
	rootCmd.AddCommand(bumpyServerCmd)
	rootCmd.AddCommand(bumpyMajorCmd)
	rootCmd.AddCommand(bumpyMinorCmd)
	rootCmd.AddCommand(bumpyPatchCmd)

	bumpyMajorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyMinorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyPatchCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
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

var bumpyMajorCmd = &cobra.Command{
	Use:   "major",
	Short: "Bump major version",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New("http://localhost:8080", "1s")
		if err != nil {
			return err
		}

		version := cmd.Flag("version").Value.String()

		stat, _ := os.Stdin.Stat()
		if version == "" {
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				input, err := io.ReadAll(os.Stdin)
				if err != nil {
					log.Fatalf("Failed to read from stdin: %v", err)
				}
				version = strings.TrimSpace(string(input))
			} else {
				return fmt.Errorf("version not provided, Use --version or pipe it via stdin")
			}
		}

		bumpedVersion, err := c.BumpMajor(version)
		if err != nil {
			return err
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpyMinorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Bump minor version",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New("http://localhost:8080", "1s")
		if err != nil {
			return err
		}

		version := cmd.Flag("version").Value.String()

		stat, _ := os.Stdin.Stat()
		if version == "" {
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				input, err := io.ReadAll(os.Stdin)
				if err != nil {
					log.Fatalf("Failed to read from stdin: %v", err)
				}
				version = strings.TrimSpace(string(input))
			} else {
				return fmt.Errorf("version not provided, Use --version or pipe it via stdin")
			}
		}

		bumpedVersion, err := c.BumpMinor(version)
		if err != nil {
			return err
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpyPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Bump patch version",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New("http://localhost:8080", "1s")
		if err != nil {
			return err
		}

		version := cmd.Flag("version").Value.String()

		stat, _ := os.Stdin.Stat()
		if version == "" {
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				input, err := io.ReadAll(os.Stdin)
				if err != nil {
					log.Fatalf("Failed to read from stdin: %v", err)
				}
				version = strings.TrimSpace(string(input))
			} else {
				return fmt.Errorf("version not provided, Use --version or pipe it via stdin")
			}
		}

		bumpedVersion, err := c.BumpPatch(version)
		if err != nil {
			return err
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpyServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start running the bumpy web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.New().Run()
		return nil
	},
}
