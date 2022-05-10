package doctor

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "doctor",
	Short: "Initialize directories and files",
	Long:  `Initialize directories and files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doctor called")
		// fmt.Println(cmd.Flags().GetString("projecName"))
	},
}

func init() {

}
