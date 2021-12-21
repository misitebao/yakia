package commands

import (
	"github.com/spf13/cobra"
)

// cmd represents the base command when called without any subcommands
var cmd = &cobra.Command{
	Use:     "yakia",
	Version: "v0.1.0",
	Short:   "Tools for developers to manage projects",
	Long:    "Tools for developers to manage projects",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Yakia version " + cmd.Root().Version)
	},
}

func init() {

	cmd.AddCommand(initCmd)
	// cmd.AddCommand(versionCmd)

	cmd.CompletionOptions.HiddenDefaultCmd = true

	cmd.SetVersionTemplate("Yakia version {{.Version}}\n")

}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
