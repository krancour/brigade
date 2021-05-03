package main

import "github.com/spf13/cobra"

var eventCmd = &cobra.Command{
	Use:     "event",
	Aliases: []string{"events"},
	Short:   "Manage events",
}

var eventCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a single event without deleting it",
	Long: "Unconditionally cancels (and aborts if applicable) a single event " +
		"whose worker is in a non-terminal phase",
	Args: cobra.NoArgs,
	Run:  func(*cobra.Command, []string) {},
}

var eventCancelManyCmd = &cobra.Command{
	Use:     "cancel-many",
	Aliases: []string{"cm"},
	Short:   "Cancel multiple events without deleting them",
	Long: "By default, only cancels events for the specified project with " +
		"their worker in a PENDING phase",
	Args: cobra.NoArgs,
	Run:  func(*cobra.Command, []string) {},
}

var eventCloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone an existing event",
	Long: "Creates a new event with the same source, type, and payload as an " +
		"existing event. The new event will be handled asynchronously according " +
		"to current project configuration, like any other new event.",
	Args: cobra.NoArgs,
	Run:  func(*cobra.Command, []string) {},
}

var eventCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new event",
	Long:  "Creates a new event for the specified project",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var eventDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a single event",
	Long: "Unconditionally deletes (and aborts if applicable) a single event " +
		"whose worker is in any phase",
	Args: cobra.NoArgs,
	Run:  func(*cobra.Command, []string) {},
}

var eventDeleteManyCmd = &cobra.Command{
	Use:     "delete-many",
	Aliases: []string{"dm"},
	Short:   "Delete multiple events",
	Long: "Deletes (and aborts if applicable) events for the specified project " +
		"with their workers in the specified phase(s)",
	Args: cobra.NoArgs,
	Run:  func(*cobra.Command, []string) {},
}

var eventGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a single event",
	Long:  "Retrieves a single event by ID",
	Args:  cobra.NoArgs,
	Run:   func(*cobra.Command, []string) {},
}

var eventListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List events",
	Long:    "Retrieves all events unless specific criteria are provided",
	Args:    cobra.NoArgs,
	Run:     func(*cobra.Command, []string) {},
}

func init() {
	eventCancelCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Cancel (and abort if applicable) the specified event (required)",
	)
	eventCancelCmd.MarkFlagRequired("id")

	eventCancelManyCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Cancel events for the specified project only (required)",
	)
	eventCancelManyCmd.MarkFlagRequired("project")
	eventCancelManyCmd.Flags().BoolVar(
		&running,
		"running",
		false,
		"If set, will additionally abort and cancel events with their worker in "+
			"a RUNNING phase",
	)
	eventCancelManyCmd.Flags().BoolVar(
		&starting,
		"starting",
		false,
		"If set, will additionally abort and cancel events with their worker in "+
			"a STARTING phase",
	)

	eventCloneCmd.Flags().BoolVarP(
		&follow,
		"follow",
		"f",
		false,
		"Synchronously wait for the event to be processed and stream logs from "+
			"its worker",
	)
	eventCloneCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Clone the specified event (required)",
	)
	eventCloneCmd.MarkFlagRequired("id")

	// TODO: There are some controversial options here that I'd like to do away
	// with.
	eventCreateCmd.Flags().BoolVarP(
		&follow,
		"follow",
		"f",
		false,
		"Synchronously wait for the event to be processed and stream logs from "+
			"its worker",
	)
	eventCreateCmd.Flags().StringVar(
		&payload,
		"payload",
		"",
		"The event payload",
	)
	eventCreateCmd.Flags().StringVar(
		&payloadFile,
		"payload-file",
		"",
		"The location of a file containing the event payload",
	)
	eventCreateCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Create an event for the specified project (required)",
	)
	eventCreateCmd.MarkFlagRequired("project")
	eventCreateCmd.Flags().StringVarP(
		&source,
		"source",
		"s",
		"brigade.sh/cli",
		"Override the default event source",
	)
	eventCreateCmd.Flags().StringVarP(
		&eventType,
		"type",
		"t",
		"exec",
		"Override the default event type",
	)

	eventDeleteCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Delete (and abort if applicable) the specified event (required)",
	)
	eventDeleteCmd.MarkFlagRequired("id")
	eventDeleteCmd.Flags().BoolVarP(
		&yes,
		"yes",
		"y",
		false,
		"Non-interactively confirm deletion",
	)

	eventDeleteManyCmd.Flags().BoolVar(
		&aborted,
		"aborted",
		false,
		"If set, will delete events with their worker in an ABORTED phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&any,
		"any-phase",
		false,
		"If set, will delete events with their worker in any phase; mutually "+
			"exclusive with all other phase flags",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&canceled,
		"canceled",
		false,
		"If set, will delete events with their worker in a CANCELED phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&failed,
		"failed",
		false,
		"If set, will delete events with their worker in a FAILED phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&pending,
		"pending",
		false,
		"If set, will delete events with their worker in a PENDING phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"Delete events for the specified project only (required)",
	)
	eventDeleteManyCmd.MarkFlagRequired("project")
	eventDeleteManyCmd.Flags().BoolVar(
		&running,
		"running",
		false,
		"If set, will abort and delete events with their worker in a RUNNING "+
			"phase; mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&starting,
		"starting",
		false,
		"If set, will delete events with their worker in a STARTING phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&succeeded,
		"succeeded",
		false,
		"If set, will delete events with their worker in a SUCCEEDED phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&terminal,
		"terminal",
		false,
		"If set, will delete events with their worker in any terminal phase; "+
			"mutually exclusive with all other phase flags",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&timedOut,
		"timed-out",
		false,
		"If set, will delete events with their worker in a TIMED_OUT phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVar(
		&unknown,
		"unknown",
		false,
		"If set, will delete events with their worker in an UNKNOWN phase; "+
			"mutually exclusive with --any-phase and --terminal",
	)
	eventDeleteManyCmd.Flags().BoolVarP(
		&yes,
		"yes",
		"y",
		false,
		"Non-interactively confirm deletion",
	)

	eventGetCmd.Flags().StringVarP(
		&id,
		"id",
		"i",
		"",
		"Retrieve the specified event (required)",
	)
	eventGetCmd.MarkFlagRequired("id")

	eventListCmd.Flags().BoolVar(
		&aborted,
		"aborted",
		false,
		"If set, will retrieve events with their worker in an ABORTED phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().BoolVar(
		&canceled,
		"canceled",
		false,
		"If set, will retrieve events with their worker in a CANCELED phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().StringVar(
		&continueVal,
		"continue",
		"",
		"Advanced-- passes an opaque value obtained from a previous command "+
			"back to the server to access the next page of results",
	)
	eventListCmd.Flags().BoolVar(
		&failed,
		"failed",
		false,
		"If set, will retrieve events with their worker in a FAILED phase; "+
			"mutually exclusive with  --terminal and --non-terminal",
	)
	eventListCmd.Flags().BoolVar(
		&nonTerminal,
		"non-terminal",
		false,
		"If set, will retrieve events with their worker in any non-terminal "+
			"phase; mutually exclusive with all other phase flags",
	)
	eventListCmd.Flags().BoolVar(
		&pending,
		"pending",
		false,
		"If set, will retrieve events with their worker in a PENDING phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"If set, will retrieve events only for the specified project",
	)
	eventListCmd.Flags().BoolVar(
		&running,
		"running",
		false,
		"If set, will retrieve events with their worker in RUNNING phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().BoolVar(
		&starting,
		"starting",
		false,
		"If set, will retrieve events with their worker in a STARTING phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().BoolVar(
		&succeeded,
		"succeeded",
		false,
		"If set, will retrieve events with their worker in a SUCCEEDED phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().BoolVar(
		&terminal,
		"terminal",
		false,
		"If set, will retrieve events with their worker in any terminal phase; "+
			"mutually exclusive with all other phase flags",
	)
	eventListCmd.Flags().BoolVar(
		&timedOut,
		"timed-out",
		false,
		"If set, will retrieve events with their worker in a TIMED_OUT phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)
	eventListCmd.Flags().BoolVar(
		&unknown,
		"unknown",
		false,
		"If set, will retrieve events with their worker in an UNKNOWN phase; "+
			"mutually exclusive with --terminal and --non-terminal",
	)

	eventCmd.AddCommand(
		eventCancelCmd,
		eventCancelManyCmd,
		eventCloneCmd,
		eventCreateCmd,
		eventDeleteCmd,
		eventDeleteManyCmd,
		eventGetCmd,
		eventListCmd,
		logCmd,
	)
}
