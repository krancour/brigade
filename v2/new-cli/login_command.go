package main

import "github.com/spf13/cobra"

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Brigade",
	Long: "By default, initiates authentication using a third-party identity " +
		"provider. This may not be supported by all Brigade API servers.",
	Args: cobra.NoArgs,
	Run:  func(*cobra.Command, []string) {},
}

func init() {
	loginCmd.Flags().BoolVarP(
		&browse,
		"browse",
		"b",
		false,
		"Use the system's default web browser to complete authentication; not "+
			"applicable when --root is used",
	)
	loginCmd.Flags().StringVarP(
		&password,
		"password",
		"p",
		"",
		"Specify the password for non-interactive root user login; only "+
			"applicable when --root is used",
	)
	loginCmd.Flags().BoolVarP(
		&root,
		"root",
		"r",
		false,
		"Log in as the root user; does not use any third party authentication",
	)
	loginCmd.Flags().StringVarP(
		&server,
		"server",
		"s",
		"",
		"Log into the API server at the specified address (required)",
	)
	loginCmd.MarkFlagRequired("server")

}
