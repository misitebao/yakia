/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v1.3.0-beta.3"

// versionCmd represents the version command
var VersonCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number standard repository ClI",
	Long:  `Print the version number standard repository ClI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
