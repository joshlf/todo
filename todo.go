package main

import (
	"fmt"
	"github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/server"
	"github.com/joshlf13/todo/shell"
	"github.com/spf13/cobra"
	"os"
    "time"
)

const (
    DATE_END = "00:00:00 1/1/2040"
    DATE_START = "00:00:00 1/1/1970"
)

func parse(d string) time.Time {
    t, err := time.Parse("15:04:05 1/2/2006", d)
    if (err != nil) {
        panic("Fix me!")
    } else {
        return t
    }
}

var todoList graph.TodoList

var file string

var alias, class, runcmd, dep, start, end string
var weight int

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
		fmt.Println("something (ran add)")
		deps := []string{dep}
		// Create task, set attributes and then add deps
		t := todoList.NewTask()
		t.SetDescription(args[0])
		t.SetRunCmd(runcmd)
		t.SetStartTime(parse(start))
		t.SetEndTime(parse(end))
		t.SetWeight(weight)
		t.AddDependencies(deps)

		// Point `alias' to `taskid' if needed
		if alias != "" {
			// TODO: Implement aliases
		}

		// Add task to `class' if needed
		if class != "" {
			// TODO: Implement classes
		}
	},
}

var modifyCommand = &cobra.Command{
	Use:   "modify ref",
	Short: "Modify a new task in the graph.",
	Long:  "Modify a new task in the graph, changing its properties (aliases, classes, times, etc.).",
	Run: func(cmd *cobra.Command, args []string) {
		// Point `alias' to `taskid'
		if alias != "" {

		}

		if class != "" {

		}

		if dep != "" {

		}

		if end != DATE_END {

		}

		if runcmd != "" {

		}

		if start != DATE_START {

		}

		if weight != 1 {

		}

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
		t, _ := todoList.ResolveSingle(args[0])
		shell.RunCommand(t.GetTaskID(), t.GetRunCmd(), background)
	},
}

var showCommand = &cobra.Command{
	Use:   "show ref [ref ...]",
	Short: "Show all unblocked tasks.",
	Long:  "Show all tasks that match the query. By default the command will just shows unblocked tasks, but this can be changed.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Uncompleted tasks:")
		uncompleted := graph.Uncompleted(todoList.Tasks)
		for i, t := range uncompleted {
			fmt.Printf("\t%v) %v", i, t)
		}
	},
}

var editCommand = &cobra.Command{
	Use:   "edit ref",
	Short: "Edit one or more tasks' descriptions.",
	Long:  "Edit one or more tasks' descriptions.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the task that this ref refers to.
		t, err := todoList.ResolveSingle(args[0])
		// Check for issue (e.g., ref was class)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error editing description: %v", err)
		} else {
			// Let user edit string
			tmp := shell.EditString(t.GetDescription())
			// Shove back updated string
			t.SetDescription(tmp)
		}
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
		cmd.Flags().StringVarP(&start, "start", "s", DATE_START, "Start time for this task.")
		cmd.Flags().StringVarP(&end, "end", "e", DATE_END, "End time for this task.")
	}

	finishCommand.Flags().BoolVarP(&obliterate, "obliterate", "o", false, "Delete this task instead of just marking it completed.")
	finishCommand.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursively finish this task and its dependencies.")
	finishCommand.Flags().BoolVarP(&requireDeps, "require-deps", "d", false, "Disallow marking this task finished until all of its dependencies are finished.")

	runCommand.Flags().BoolVarP(&background, "background", "b", false, "Fork to the background.")

	serverCommand.Flags().IntVarP(&port, "port", "p", 8080, "Port the server should use.")
	serverCommand.Flags().BoolVarP(&noRestart, "no-restart", "n", false, "Stop the server from restarting when the JSON file changes. (Only applies when running read-only.)")
	serverCommand.Flags().BoolVarP(&readonly, "readonly", "r", false, "Prevent anyone from modifying the todo graph over HTTP.")
}
