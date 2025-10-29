package cmd

import (
	"bumpy/package/server"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	client "github.com/apgmckay/bumpy-client"
	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

var rootCmd = BumpyRootCmd()

func init() {
	rootCmd.AddCommand(bumpyServerCmd)
	rootCmd.AddCommand(bumpyMajorCmd)
	rootCmd.AddCommand(bumpyMinorCmd)
	rootCmd.AddCommand(bumpyPatchCmd)

	bumpyMajorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyMajorCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpyMajorCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpyMajorCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpyMajorCmd.MarkPersistentFlagRequired("package-name")

	bumpyMinorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyMinorCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpyMinorCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpyMinorCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpyMinorCmd.MarkPersistentFlagRequired("package-name")

	bumpyPatchCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyPatchCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpyPatchCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpyPatchCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpyPatchCmd.MarkPersistentFlagRequired("package-name")
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

		params := make(map[string]string, 2)

		params["version"] = version
		params["pre-release"] = cmd.Flag("pre-release").Value.String()
		params["build"] = cmd.Flag("build").Value.String()
		params["package_name"] = cmd.Flag("package-name").Value.String()

		bumpedVersion, err := c.GetMajor(params)
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

		params := make(map[string]string, 2)

		params["version"] = version
		params["pre-release"] = cmd.Flag("pre-release").Value.String()
		params["build"] = cmd.Flag("build").Value.String()
		params["package_name"] = cmd.Flag("package-name").Value.String()

		bumpedVersion, err := c.GetMinor(params)
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

		params := make(map[string]string, 2)

		params["version"] = version
		params["pre-release"] = cmd.Flag("pre-release").Value.String()
		params["build"] = cmd.Flag("build").Value.String()
		params["package_name"] = cmd.Flag("package-name").Value.String()

		bumpedVersion, err := c.GetPatch(params)
		if err != nil {
			return err
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpyServerCmd = &cobra.Command{
	Use:    "server",
	Short:  "Start running the bumpy web server",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		server.New().Run()
		return nil
	},
}
