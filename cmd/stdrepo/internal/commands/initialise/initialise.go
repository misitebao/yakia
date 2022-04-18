package initialise

import (
	"fmt"

	ini18n "github.com/misitebao/standard-repository/cmd/stdrepo/internal/i18n"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: ini18n.T("commands.short"),
	Long:  ini18n.T("commands.long"),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		// fmt.Println(cmd.Flags().GetString("projecName"))
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// initCmd.Flags().StringP("projecName", "n","", "Set the project name.")
}
