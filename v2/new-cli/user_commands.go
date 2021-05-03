package main

import "github.com/spf13/cobra"

var userCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"users"},
	Short:   "Manage users",
}

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a single user",
	Long:  "Retrieve a single user by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var userListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Retrieve all users",
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

var userLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock a user out of Brigade",
	Long:  "Lock a user out of Brigade by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var userUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Restore a user's access to Brigade",
	Long:  "Restore a user's access to Brigade by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	userCmd.AddCommand(
		userGetCmd,
		userListCmd,
		userLockCmd,
		userUnlockCmd,
	)

	userGetCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Retrieve the specified user (required)",
	)
	userGetCmd.MarkFlagRequired("id")

	userListCmd.Flags().StringVar(
		&continueVal,
		"continue",
		"",
		"Advanced-- passes an opaque value obtained from a previous command back "+
			"to the server to access the next page of results",
	)

	userLockCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Lock the specified user (required)",
	)
	userLockCmd.MarkFlagRequired("id")

	userUnlockCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Unlock the specified user (required)",
	)
	userUnlockCmd.MarkFlagRequired("id")
}
