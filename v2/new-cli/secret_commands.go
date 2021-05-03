package main

import "github.com/spf13/cobra"

var secretCmd = &cobra.Command{
	Use:     "secret",
	Aliases: []string{"secrets"},
	Short:   "Manage project secrets",
}

var secretListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Retrieve all project secrets",
	Long:    "Retrieve all project secrets; values are always redacted",
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

var secretSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Define or redefine the value of one or more secrets",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var secretUnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Clear the value of one or more secrets",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	secretCmd.AddCommand(
		secretListCmd,
		secretSetCmd,
		secretUnsetCmd,
	)

	secretListCmd.Flags().StringVar(
		&continueVal,
		"continue",
		"",
		"Advanced-- passes an opaque value obtained from a previous command "+
			"back to the server to access the next page of results",
	)
	secretListCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Retrieve secrets for the specified project (required)",
	)
	secretListCmd.MarkFlagRequired("project")

	secretSetCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Set secrets for the specified project (required)",
	)
	secretSetCmd.MarkFlagRequired("project")
	secretSetCmd.Flags().StringVarP(
		&filename,
		"file",
		"f",
		"",
		`A "flat" JSON or YAML file containing secrets as key/value pairs`,
	)
	secretSetCmd.Flags().StringArrayVarP(
		&set,
		"set",
		"s",
		[]string{},
		"Set a secret using the specified key=value pair. Secrets "+
			"specified using this flag take precedence over any specified "+
			"using the --file flag",
	)

	secretUnsetCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Clear secrets for the specified project (required)",
	)
	secretUnsetCmd.MarkFlagRequired("project")
	secretUnsetCmd.Flags().StringArrayVarP(
		&unset,
		"unset",
		"u",
		[]string{},
		"Clear a secret having the specified key (required)",
	)

}
