package commands

import (
	"github.com/misitebao/yakia/cmd/yakia/internal/commands/dev"
	"github.com/misitebao/yakia/cmd/yakia/internal/commands/doctor"
	"github.com/misitebao/yakia/cmd/yakia/internal/commands/initialise"
	ini18n "github.com/misitebao/yakia/cmd/yakia/internal/i18n"
	"github.com/spf13/cobra"
)

// cmd represents the base command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "yakia",
	Short: ini18n.T("commands.short"),
	Long:  ini18n.T("commands.long"),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args)

	},
}

func init() {
	// Add subcommand
	cmd.AddCommand(VersonCmd)
	cmd.AddCommand(initialise.Cmd)
	cmd.AddCommand(doctor.Cmd)
	cmd.AddCommand(dev.Cmd)

	cmd.CompletionOptions.HiddenDefaultCmd = true

}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
