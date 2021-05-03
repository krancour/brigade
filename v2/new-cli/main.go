package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// For flags
	aborted                   bool
	any                       bool
	browse                    bool
	canceled                  bool
	container                 string
	continueVal               string
	eventType                 string
	failed                    bool
	description               string
	disableInteractivePrompts bool
	filename                  string
	follow                    bool
	id                        string
	ignoreCertErrors          bool
	job                       string
	nonTerminal               bool
	password                  string
	payload                   string
	payloadFile               string
	pending                   bool
	project                   string
	role                      string
	root                      bool
	running                   bool
	server                    string
	serviceAccount            string
	serviceAccounts           []string
	source                    string
	set                       []string
	starting                  bool
	succeeded                 bool
	terminal                  bool
	timedOut                  bool
	unknown                   bool
	unset                     []string
	user                      string
	users                     []string
	yes                       bool
)

var rootCmd = &cobra.Command{
	Use:   "brig",
	Short: "Event Driven Scripting for Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(
		eventCmd,
		loginCmd,
		logoutCmd,
		projectCmd,
		roleCmd,
		serviceAccountCmd,
		userCmd,
	)
	rootCmd.PersistentFlags().BoolVarP(
		&disableInteractivePrompts,
		"non-interactive",
		"n",
		false,
		"Disable all interactive prompts",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&ignoreCertErrors,
		"insecure",
		"k",
		false,
		"Ignore certificate errors when using HTTPS",
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
