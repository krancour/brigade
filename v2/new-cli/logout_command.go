package main

import "github.com/spf13/cobra"

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of Brigade",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}
