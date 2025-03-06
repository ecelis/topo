/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Topo cross-platform network utility.\nCopyright © 2025 Ernesto Celis <ernesto@patito.ninja>",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("topo v0.1 -- HEAD")
	},
}
