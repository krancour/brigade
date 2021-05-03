package main

import "github.com/spf13/cobra"

var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"projects"},
	Short:   "Manage projects",
}

var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long:  "Create a new project from a project definition file",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single project",
	Long:  "Delete a single project by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a single project",
	Long:  "Retrieve a single project by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var projectListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Retrieve all projects",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

var projectUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing project",
	Long:  "Update an existing project from a project definition file",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	projectCmd.AddCommand(
		projectCreateCmd,
		projectDeleteCmd,
		projectGetCmd,
		projectListCmd,
		projectRoleCmd,
		secretCmd,
		projectUpdateCmd,
	)

	projectCreateCmd.Flags().StringVarP(
		&filename,
		"file",
		"f",
		"",
		"A YAML or JSON file that describes the project (required)",
	)
	projectCreateCmd.MarkFlagRequired("file")

	projectDeleteCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Delete the specified project (required)",
	)
	projectDeleteCmd.MarkFlagRequired("id")
	projectDeleteCmd.Flags().BoolVarP(
		&yes,
		"yes",
		"y",
		false,
		"Non-interactively confirm deletion",
	)

	projectGetCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Retrieve the specified project (required)",
	)
	projectGetCmd.MarkFlagRequired("id")

	projectListCmd.Flags().StringVar(
		&continueVal,
		"continue",
		"",
		"Advanced-- passes an opaque value obtained from a previous command back "+
			"to the server to access the next page of results",
	)

	projectUpdateCmd.Flags().StringVarP(
		&filename,
		"file",
		"f",
		"",
		"A YAML or JSON file that describes the project (required)",
	)
	projectUpdateCmd.MarkFlagRequired("file")
}
