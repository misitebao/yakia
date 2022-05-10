package dev

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dev",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev called")
		// fmt.Println(cmd.Flags().GetString("projecName"))
	},
}

func init() {

}
