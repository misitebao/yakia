package commands

import (
	"github.com/misitebao/standard-repository/cmd/stdrepo/internal/commands/doctor"
	"github.com/misitebao/standard-repository/cmd/stdrepo/internal/commands/initialise"
	ini18n "github.com/misitebao/standard-repository/cmd/stdrepo/internal/i18n"
	"github.com/spf13/cobra"
)

// cmd represents the base command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "stdrepo",
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

	cmd.CompletionOptions.HiddenDefaultCmd = true

}

func Execute() {
	cobra.CheckErr(cmd.Execute())
}
