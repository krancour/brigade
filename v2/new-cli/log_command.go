package main

import "github.com/spf13/cobra"

var logCmd = &cobra.Command{
	Use:     "log",
	Aliases: []string{"logs"},
	Short:   "View worker or job logs",
	Long:    "View worker or job logs for a specified event",
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

func init() {
	logCmd.Flags().StringVarP(
		&container,
		"container",
		"c",
		"",
		"View logs from the specified container; if not set, displays logs from "+
			"the worker or job's primary container",
	)
	logCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"View logs from the specified event",
	)
	logCmd.MarkFlagRequired("id")
	logCmd.Flags().BoolVarP(
		&follow,
		"follow",
		"f",
		false,
		"If set, will stream logs until interrupted",
	)
	logCmd.Flags().StringVarP(
		&job,
		"job",
		"j",
		"",
		"View logs from the specified job; if not set, displays worker logs",
	)
}
