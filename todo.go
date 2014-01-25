package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"math"
    "github.com/joshlf13/todo/server"
    "github.com/joshlf13/todo/graph"
)

var todoList graph.TodoList

var file string

var alias, class, runcmd, dep string
var weight, start, end int

var obliterate, recursive, requireDeps bool

var background bool

var port int
var noRestart, readonly bool

var TodoCommand = &cobra.Command{
	Use:   "todo",
	Short: "A todo list manager with some more interesting features.",
	Long:  "A todo list manager that allows for dependencies, start and end times, executing commands and more.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dsafdsfsadf")
	},
}

var addCommand = &cobra.Command{
	Use:   "add description",
	Short: "Add a new task to the graph.",
	Long:  "Add a new task to the graph, specifying its properties (aliases, classes, times, etc.).",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var modifyCommand = &cobra.Command{
	Use:   "modify ref",
	Short: "Modify a new task in the graph.",
	Long:  "Modify a new task in the graph, changing its properties (aliases, classes, times, etc.).",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var finishCommand = &cobra.Command{
	Use:   "finish ref",
	Short: "Finish one or more tasks in the graph.",
	Long:  "Finish one or more tasks in the graph, and potentially their dependencies.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var runCommand = &cobra.Command{
	Use:   "run ref",
	Short: "Run a task's command.",
	Long:  "Run a task's command. Can be run forked to run in the background if desired.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var showCommand = &cobra.Command{
	Use:   "show ref [ref ...]",
	Short: "Show all unblocked tasks.",
	Long:  "Show all tasks that match the query. By default the command will just shows unblocked tasks, but this can be changed.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("About to show")
		for _, arg := range args {
			fmt.Printf("Showing: %s\n", arg)
		}
	},
}

var editCommand = &cobra.Command{
	Use:   "edit ref",
	Short: "Edit one or more tasks' descriptions.",
	Long:  "Edit one or more tasks' descriptions.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start a todo server.",
	Long:  "Start a todo server that provides a web navigable interface and an API to allow you to view and potentially modify your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
        server.StartServer(todoList, port, noRestart)
	},
}

var optionCommand = &cobra.Command{
	Use:   "option",
	Short: "Set an option for your todo graph.",
	Long:  "Set an option for your todo graph, affecting future behaviour of todo commands.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	TodoCommand.Flags().StringVarP(&file, "file", "f", "todo.json", "JSON file to read graph from.")

	// Task commands
	TodoCommand.AddCommand(addCommand, modifyCommand, finishCommand, runCommand, showCommand, editCommand)

	// Other commands
	TodoCommand.AddCommand(serverCommand, optionCommand)

	TodoCommand.Execute()
}

func init() {
	for _, cmd := range []*cobra.Command{addCommand, modifyCommand} {
		cmd.Flags().StringVarP(&alias, "alias", "a", "", "Create an alias that references this task.")
		cmd.Flags().StringVarP(&class, "class", "c", "", "Place this alias within a class.")
		cmd.Flags().StringVarP(&runcmd, "run", "r", "", "Run a command when doing this task.")
		cmd.Flags().StringVarP(&dep, "dep", "d", "", "Add a dependency for this task.")
		cmd.Flags().IntVarP(&weight, "weight", "w", 1, "Assign this task a weight.")
		cmd.Flags().IntVarP(&start, "start", "s", 0, "Start time for this task.")
		cmd.Flags().IntVarP(&end, "end", "e", math.MaxInt64, "End time for this task.")
	}

	finishCommand.Flags().BoolVarP(&obliterate, "obliterate", "o", false, "Delete this task instead of just marking it completed.")
	finishCommand.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursively finish this task and its dependencies.")
	finishCommand.Flags().BoolVarP(&requireDeps, "require-deps", "d", false, "Disallow marking this task finished until all of its dependencies are finished.")

	runCommand.Flags().BoolVarP(&background, "background", "b", false, "Fork to the background.")

	serverCommand.Flags().IntVarP(&port, "port", "p", 8080, "Port the server should use.")
	serverCommand.Flags().BoolVarP(&noRestart, "no-restart", "n", false, "Stop the server from restarting when the JSON file changes. (Only applies when running read-only.)")
	serverCommand.Flags().BoolVarP(&readonly, "readonly", "r", false, "Prevent anyone from modifying the todo graph over HTTP.")
}
