/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gocker/pkg/docker"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal(err)
		}
		s := docker.ProvideService()

		containers, err := s.List()
		if all {
			containers, err = s.ListAll()
		}

		if err != nil {
			log.Fatal(err)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "CONTAINER ID", "IMAGE", "CREATED", "STATUS"})

		for i, container := range containers {
			t.AppendRows([]table.Row{
				{
					converInterface(fmt.Sprintf("%d", i)),
					converInterface(container.ID[0:12]),
					converInterface(container.Image),
					converInterface(fmt.Sprintf("%d", container.Created)),
					converInterface(container.Status),
				},
			})

		}
		t.AppendSeparator()
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.Flags().BoolP("all", "a", false, "gocker ps --all")
}

func converInterface(variable string) (a interface{}) {
	a = variable
	return
}
