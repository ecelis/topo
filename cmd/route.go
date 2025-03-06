/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ecelis/topo/internal/platform"
	"github.com/ecelis/topo/internal/route"
	"github.com/spf13/cobra"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "Display system network routes",
	Run: func(cmd *cobra.Command, args []string) {
		expected, err := route.ReadConfig(rootCmd.PersistentFlags().Lookup("config").Value.String())
		if err != nil {
			log.Printf("%v", err)
			os.Exit(1)
		}

		checker := &platform.RouteChecker{
			Routes: expected,
		}
		err = checker.CheckAndUpdateRoutes()

		if err != nil {
			fmt.Println("Error checking routes:", err)
			os.Exit(1)
		} else {
			fmt.Println("Routes checked succesfully")
		}
		// fmt.Println(route.)
	},
}

func init() {
	rootCmd.AddCommand(routeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
