/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"gocker/pkg/docker"
	"log"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Get container detail",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		s := docker.ProvideService()
		details, err := s.Get(id)
		if err != nil {
			log.Fatal(err)
		}
		detailString, err := json.Marshal(details)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(detailString))
	},
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	inspectCmd.Flags().String("id", "", "Help message for toggle")

}
