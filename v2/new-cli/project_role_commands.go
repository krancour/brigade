package main

import (
	"fmt"

	"github.com/brigadecore/brigade/sdk/v2/core"
	"github.com/spf13/cobra"
)

var projectRoleCmd = &cobra.Command{
	Use:     "role",
	Aliases: []string{"roles"},
	Short:   "Manage project-level role assignments",
}

var projectRoleGrantCmd = &cobra.Command{
	Use:   "grant",
	Short: "Grant a project-level role to a user or service account",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var projectRoleListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List users and/or service accounts and their project-level roles",
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

var projectRoleRevokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke a project-level role from a user or service account",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var projectRoleGrantAdminCmd = &cobra.Command{
	Use:   string(core.RoleProjectAdmin),
	Short: fmt.Sprintf("Grant the %s role", core.RoleProjectAdmin),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables management of all "+
			"aspects of the project, including its secrets, as well as "+
			"project-level permissions for other users and service "+
			"accounts.",
		core.RoleProjectAdmin,
	),
	Run: func(*cobra.Command, []string) {},
}

var projectRoleGrantDeveloperCmd = &cobra.Command{
	Use:   string(core.RoleProjectDeveloper),
	Short: fmt.Sprintf("Grant the %s role", core.RoleProjectDeveloper),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables updating the project "+
			"definition, but does NOT enable management of the project's "+
			"secrets or project-level permissions for other users and "+
			"service accounts.",
		core.RoleProjectDeveloper,
	),
	Run: func(*cobra.Command, []string) {},
}

var projectRoleGrantUserCmd = &cobra.Command{
	Use:   string(core.RoleProjectUser),
	Short: fmt.Sprintf("Grant the %s role", core.RoleProjectUser),
	Long: fmt.Sprintf(
		"Grant the %s role, which enables creation and management "+
			"of events associated with the project",
		core.RoleProjectUser,
	),
	Run: func(*cobra.Command, []string) {},
}

var projectRoleRevokeAdminCmd = &cobra.Command{
	Use:   string(core.RoleProjectAdmin),
	Short: fmt.Sprintf("Revoke the %s role", core.RoleProjectAdmin),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables management of all "+
			"aspects of the project, including its secrets, as well as "+
			"project-level permissions for other users and service "+
			"accounts.",
		core.RoleProjectAdmin,
	),
	Run: func(*cobra.Command, []string) {},
}

var projectRoleRevokeDeveloperCmd = &cobra.Command{
	Use:   string(core.RoleProjectDeveloper),
	Short: fmt.Sprintf("Revoke the %s role", core.RoleProjectDeveloper),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables updating the project "+
			"definition, but does NOT enable management of the project's "+
			"secrets or project-level permissions for other users and "+
			"service accounts.",
		core.RoleProjectDeveloper,
	),
	Run: func(*cobra.Command, []string) {},
}

var projectRoleRevokeUserCmd = &cobra.Command{
	Use:   string(core.RoleProjectUser),
	Short: fmt.Sprintf("Revoke the %s role", core.RoleProjectUser),
	Long: fmt.Sprintf(
		"Revoke the %s role, which enables creation and "+
			"management of events associated with the project",
		core.RoleProjectUser,
	),
	Run: func(*cobra.Command, []string) {},
}

func init() {
	projectRoleCmd.AddCommand(
		projectRoleGrantCmd,
		projectRoleListCmd,
		projectRoleRevokeCmd,
	)

	projectRoleGrantCmd.AddCommand(
		projectRoleGrantAdminCmd,
		projectRoleGrantDeveloperCmd,
		projectRoleGrantUserCmd,
	)

	projectRoleRevokeCmd.AddCommand(
		projectRoleRevokeAdminCmd,
		projectRoleRevokeDeveloperCmd,
		projectRoleRevokeUserCmd,
	)

	projectRoleListCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"List principals and their roles for the specified project (required)",
	)
	projectRoleListCmd.MarkFlagRequired("project")
	projectRoleListCmd.Flags().StringVarP(
		&role,
		"role",
		"r",
		"",
		"Narrow results to the specified role",
	)
	projectRoleListCmd.Flags().StringVarP(
		&serviceAccount,
		"service-account",
		"s",
		"",
		"Narrow results to the specified service account; mutually exclusive "+
			"with --user",
	)
	projectRoleListCmd.Flags().StringVarP(
		&user,
		"user",
		"u",
		"",
		"Narrow results to the specified user; mutually exclusive with "+
			"--service-account",
	)

	addProjectRoleGrantFlags(projectRoleGrantAdminCmd)
	addProjectRoleGrantFlags(projectRoleGrantDeveloperCmd)
	addProjectRoleGrantFlags(projectRoleGrantUserCmd)

	addProjectRoleRevokeFlags(projectRoleRevokeAdminCmd)
	addProjectRoleRevokeFlags(projectRoleRevokeDeveloperCmd)
	addProjectRoleRevokeFlags(projectRoleRevokeUserCmd)
}

func addProjectRoleGrantFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Grant the role for the specified project (required)",
	)
	cmd.MarkFlagRequired("project")
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

func addProjectRoleRevokeFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Revoke the role for the specified project (required)",
	)
	cmd.MarkFlagRequired("project")
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
