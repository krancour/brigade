package main

import (
	"fmt"

	"github.com/brigadecore/brigade/sdk/v2/system"
	"github.com/spf13/cobra"
)

var roleCmd = &cobra.Command{
	Use:     "role",
	Aliases: []string{"roles"},
	Short:   "Manage system-level role assignments",
}

var roleGrantCmd = &cobra.Command{
	Use:   "grant",
	Short: "Grant a system-level role to a user or service account",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var roleListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List users and/or service accounts and their system-level roles",
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

var roleRevokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke a system-level role from a user or service account",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var roleGrantAdminCmd = &cobra.Command{
	Use:   string(system.RoleAdmin),
	Short: fmt.Sprintf("Grant the %s role", system.RoleAdmin),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables system management including "+
			"system-level permissions for other users and service accounts.",
		system.RoleAdmin,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleGrantEventCreatorCmd = &cobra.Command{
	Use:   string(system.RoleEventCreator),
	Short: fmt.Sprintf("Grant the %s role", system.RoleEventCreator),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables creation of events for all projects.",
		system.RoleEventCreator,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleGrantProjectCreatorCmd = &cobra.Command{
	Use:   string(system.RoleProjectCreator),
	Short: fmt.Sprintf("Grant the %s role", system.RoleProjectCreator),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables creation of new projects.",
		system.RoleProjectCreator,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleGrantReaderCmd = &cobra.Command{
	Use:   string(system.RoleReader),
	Short: fmt.Sprintf("Grant the %s role", system.RoleReader),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables global read-only access to Brigade.",
		system.RoleReader,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleRevokeAdminCmd = &cobra.Command{
	Use:   string(system.RoleAdmin),
	Short: fmt.Sprintf("Revoke the %s role", system.RoleAdmin),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables system management including "+
			"system-level permissions for other users and service accounts.",
		system.RoleAdmin,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleRevokeEventCreatorCmd = &cobra.Command{
	Use:   string(system.RoleEventCreator),
	Short: fmt.Sprintf("Revoke the %s role", system.RoleEventCreator),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables creation of events for all projects.",
		system.RoleEventCreator,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleRevokeProjectCreatorCmd = &cobra.Command{
	Use: string(system.RoleProjectCreator),
	Short: fmt.Sprintf(
		"Revoke the %s role",
		system.RoleProjectCreator,
	),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables creation of new projects.",
		system.RoleProjectCreator,
	),
	Run: func(*cobra.Command, []string) {},
}

var roleRevokeReaderCmd = &cobra.Command{
	Use:   string(system.RoleReader),
	Short: fmt.Sprintf("Revoke the %s role", system.RoleReader),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables global read-only access to Brigade.",
		system.RoleReader,
	),
	Run: func(*cobra.Command, []string) {},
}

func init() {
	roleCmd.AddCommand(
		roleGrantCmd,
		roleListCmd,
		roleRevokeCmd,
	)

	roleGrantCmd.AddCommand(
		roleGrantAdminCmd,
		roleGrantEventCreatorCmd,
		roleGrantProjectCreatorCmd,
		roleGrantReaderCmd,
	)

	roleRevokeCmd.AddCommand(
		roleRevokeAdminCmd,
		roleRevokeEventCreatorCmd,
		roleRevokeProjectCreatorCmd,
		roleRevokeReaderCmd,
	)

	roleListCmd.Flags().StringVarP(
		&role,
		"role",
		"r",
		"",
		"Narrow results to the specified role",
	)
	roleListCmd.Flags().StringVarP(
		&serviceAccount,
		"service-account",
		"s",
		"",
		"Narrow results to the specified service account; mutually exclusive "+
			"with --user",
	)
	roleListCmd.Flags().StringVarP(
		&user,
		"user",
		"u",
		"",
		"Narrow results to the specified user; mutually exclusive with "+
			"--service-account",
	)

	addRoleGrantFlags(roleGrantAdminCmd)
	addRoleGrantFlags(roleGrantEventCreatorCmd)
	roleGrantEventCreatorCmd.Flags().StringVar(
		&source,
		"source",
		"",
		"Permit creation of events from the specified source only (required)",
	)
	roleGrantEventCreatorCmd.MarkFlagRequired("source")
	addRoleGrantFlags(roleGrantProjectCreatorCmd)
	addRoleGrantFlags(roleGrantReaderCmd)

	addRoleRevokeFlags(roleRevokeAdminCmd)
	addRoleRevokeFlags(roleRevokeEventCreatorCmd)
	roleRevokeEventCreatorCmd.Flags().StringVar(
		&source,
		"source",
		"",
		"Revoke creation of events from the specified source only (required)",
	)
	roleRevokeEventCreatorCmd.MarkFlagRequired("source")
	addRoleRevokeFlags(roleRevokeProjectCreatorCmd)
	addRoleRevokeFlags(roleRevokeReaderCmd)
}

func addRoleGrantFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayVarP(
		&serviceAccounts,
		"service-account",
		"s",
		[]string{},
		"Grant the role to the specified service account",
	)
	cmd.Flags().StringArrayVarP(
		&users,
		"user",
		"u",
		[]string{},
		"Grant the role to the specified user",
	)
}

func addRoleRevokeFlags(cmd *cobra.Command) {
	cmd.Flags().StringArrayVarP(
		&serviceAccounts,
		"service-account",
		"s",
		[]string{},
		"Revoke the role from the specified service account",
	)
	cmd.Flags().StringArrayVarP(
		&users,
		"user",
		"u",
		[]string{},
		"Revoke the role from the specified user",
	)
}
