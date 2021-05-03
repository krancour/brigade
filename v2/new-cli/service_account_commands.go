package main

import "github.com/spf13/cobra"

var serviceAccountCmd = &cobra.Command{
	Use:     "service-account",
	Aliases: []string{"service-accounts", "sa"},
	Short:   "Manage service accounts",
}

var serviceAccountCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new service account",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var serviceAccountGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a single service account",
	Long:  "Retrieve a single service account by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var serviceAccountListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Retrieve all service accounts",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

var serviceAccountLockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Lock a service account out of Brigade",
	Long:  "Lock a service account out of Brigade by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var serviceAccountUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Restore a service account's access to Brigade",
	Long:  "Restore a service account's access to Brigade by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	serviceAccountCmd.AddCommand(
		serviceAccountCreateCmd,
		serviceAccountGetCmd,
		serviceAccountListCmd,
		serviceAccountLockCmd,
		serviceAccountUnlockCmd,
	)

	serviceAccountCreateCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Create a service account with the specified ID (required)",
	)
	serviceAccountCreateCmd.MarkFlagRequired("id")
	serviceAccountCreateCmd.Flags().StringVarP(
		&description,
		"description",
		"d",
		"",
		"Create a service account with the specified description (required)",
	)
	serviceAccountCreateCmd.MarkFlagRequired("description")

	serviceAccountGetCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Retrieve the specified service account (required)",
	)
	serviceAccountGetCmd.MarkFlagRequired("id")

	serviceAccountListCmd.Flags().StringVar(
		&continueVal,
		"continue",
		"",
		"Advanced-- passes an opaque value obtained from a previous command back "+
			"to the server to access the next page of results",
	)

	serviceAccountLockCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Lock the specified service account (required)",
	)
	serviceAccountLockCmd.MarkFlagRequired("id")

	serviceAccountUnlockCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Unlock the specified service account (required)",
	)
	serviceAccountUnlockCmd.MarkFlagRequired("id")
}
