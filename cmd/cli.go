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
	rootCmd.AddCommand(bumpyGetCmd)
	rootCmd.AddCommand(bumpySetCmd)

	rootCmd.AddCommand(bumpyServerCmd)

	bumpyGetCmd.AddCommand(bumpyGetMajorCmd)
	bumpyGetCmd.AddCommand(bumpyGetMinorCmd)
	bumpyGetCmd.AddCommand(bumpyGetPatchCmd)

	bumpySetCmd.AddCommand(bumpySetMajorCmd)
	bumpySetCmd.AddCommand(bumpySetMinorCmd)
	bumpySetCmd.AddCommand(bumpySetPatchCmd)

	rootCmd.AddCommand(bumpyBlockedCmd)

	bumpyGetMajorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyGetMajorCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpyGetMajorCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpyGetMajorCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpyGetMajorCmd.MarkPersistentFlagRequired("package-name")

	bumpyGetMinorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyGetMinorCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpyGetMinorCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpyGetMinorCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpyGetMinorCmd.MarkPersistentFlagRequired("package-name")

	bumpyGetPatchCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpyGetPatchCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpyGetPatchCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpyGetPatchCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpyGetPatchCmd.MarkPersistentFlagRequired("package-name")

	//

	bumpySetMajorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpySetMajorCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpySetMajorCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpySetMajorCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpySetMajorCmd.MarkPersistentFlagRequired("package-name")

	bumpySetMinorCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpySetMinorCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpySetMinorCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpySetMinorCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpySetMinorCmd.MarkPersistentFlagRequired("package-name")

	bumpySetPatchCmd.PersistentFlags().StringP("version", "v", "", "version you wish to bump")
	bumpySetPatchCmd.PersistentFlags().StringP("pre-release", "p", "", "pre-release version tag to append to the version number")
	bumpySetPatchCmd.PersistentFlags().StringP("build", "b", "", "build version tag to append to the version number")
	bumpySetPatchCmd.PersistentFlags().StringP("package-name", "n", "", "name of the package")

	bumpySetPatchCmd.MarkPersistentFlagRequired("package-name")
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

var bumpyGetCmd = &cobra.Command{
	Use:   "get",
	Long:  "Get bumpy version information for an application",
	Short: "Get bumpy version information for an application",
}

var bumpySetCmd = &cobra.Command{
	Use:   "set",
	Long:  "Set bumpy version information for an application",
	Short: "Set bumpy version information for an application",
}

var bumpyBlockedCmd = &cobra.Command{
	Use:   "blocked",
	Long:  "Bumpy has a blocked status to indicate if deploys for the package can happen",
	Short: "Bumpy blocked status",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.New("http://localhost:8080", "1s")
		if err != nil {
			return err
		}
		result, err := c.GetBlocked()
		if err != nil {
			return err
		}

		fmt.Println(result)

		return nil
	},
}

var bumpyGetMajorCmd = &cobra.Command{
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

		var bumpedVersion string

		sendEvent, _ := cmd.Flags().GetBool("event")
		if sendEvent {
			bumpedVersion, err = c.PostBumpMajor(params, strings.NewReader(""))
			if err != nil {
				return err
			}
		} else {
			bumpedVersion, err = c.GetBumpMajor(params)
			if err != nil {
				return err
			}
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpyGetMinorCmd = &cobra.Command{
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

		var bumpedVersion string

		sendEvent, _ := cmd.Flags().GetBool("event")
		if sendEvent {
			bumpedVersion, err = c.PostBumpMinor(params, strings.NewReader(""))
			if err != nil {
				return err
			}
		} else {
			bumpedVersion, err = c.GetBumpMinor(params)
			if err != nil {
				return err
			}
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpyGetPatchCmd = &cobra.Command{
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

		var bumpedVersion string

		sendEvent, _ := cmd.Flags().GetBool("event")
		if sendEvent {
			bumpedVersion, err = c.PostBumpPatch(params, strings.NewReader(""))
			if err != nil {
				return err
			}
		} else {
			bumpedVersion, err = c.GetBumpPatch(params)
			if err != nil {
				return err
			}
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpySetPatchCmd = &cobra.Command{
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

		var bumpedVersion string

		bumpedVersion, err = c.PostBumpPatch(params, strings.NewReader(""))
		if err != nil {
			return err
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpySetMinorCmd = &cobra.Command{
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

		var bumpedVersion string

		bumpedVersion, err = c.PostBumpMinor(params, strings.NewReader(""))
		if err != nil {
			return err
		}

		fmt.Println(bumpedVersion)

		return nil
	},
}

var bumpySetMajorCmd = &cobra.Command{
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

		var bumpedVersion string

		bumpedVersion, err = c.PostBumpMajor(params, strings.NewReader(""))
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
