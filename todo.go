package main

import (
	"fmt"
	"github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/server"
	_ "github.com/joshlf13/todo/shell"
	"github.com/spf13/cobra"
	"os"
	"time"
)

const (
	DATE_END   = "00:00:00 1/1/2040"
	DATE_START = "00:00:00 1/1/1970"
)

func parse(d string) time.Time {
	t, err := time.Parse("15:04:05 1/2/2006", d)
	if err != nil {
		panic("Fix me!")
	} else {
		return t
	}
}

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
		m, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to acquire resource to add new task: %v", err)
			return
		}
		t := graph.Task{}
		id, err := m.AddTask(t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add new task to resource: %v", err)
			return
		}
		if dep != "" {
			m.AddDependency(id, graph.TaskID(dep))
		}
		m.SetStartTime(id, parse(start))
		m.SetEndTime(id, parse(end))
		// m.SetDescription(id, args[0])
		// m.SetRunCmd(id, runcmd)
		// m.SetWeight(id, weight)

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
		if len(args) == 0 {
			// TODO: Print help message. HOW DO I DO THIS, COBRA?!?!?!
			return
		}

		ref := args[0]
		m, err := getMiddleman(file, true)
		ids := []graph.TaskID{graph.TaskID(ref)} // When we have aliases and classes, resolve ref.
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to modify %v: %v\n", ref, err)
		} else {
			for _, id := range ids {
				if alias != "" {
					// TODO: Implement aliases
				}

				if class != "" {
					// TODO: Implement classes
				}

				if dep != "" {
					err = m.AddDependency(id, graph.TaskID(dep))
				}

				if end != DATE_END {
					m.SetEndTime(id, parse(end))
				}

				if start != DATE_START {
					m.SetStartTime(id, parse(start))
				}

				if runcmd != "" {
					// TODO: Implement
				}

				if weight != 1 {
					// TODO: Implement
				}
			}
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
		if background {
			// TODO: daemonize
		}
		_, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to acquire resource to get run command: %v", err)
			return
		}
		// TODO: Uncomment below when we have run commands
		// ref := args[0]
		// id := graph.TaskID(ref) // When we have aliases and classes, resolve here.
		// shell.RunCommand(id, m.GetRunCmd(id), background)
	},
}

var showCommand = &cobra.Command{
	Use:   "show ref [ref ...]",
	Short: "Show all unblocked tasks.",
	Long:  "Show all tasks that match the query. By default the command will just shows unblocked tasks, but this can be changed.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Uncompleted tasks:")
		m, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to acquire resource: %v", err)
			return
		}
		unblocked, err := m.GetUnblocked()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to find unblocked commands: %v", err)
			return
		}
		for i, t := range unblocked {
			fmt.Printf("\t%v) %v", i, t)
		}
	},
}

var editCommand = &cobra.Command{
	Use:   "edit ref",
	Short: "Edit one or more tasks' descriptions.",
	Long:  "Edit one or more tasks' descriptions.",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getMiddleman(file, true)
		// Get the task that this ref refers to.
		ref := args[0]
		id := graph.TaskID(ref) // When we have aliases and classes, resolve here.
		// Check for issue (e.g., ref was class)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error editing description for %v: %v", id, err)
		} else {
			// Let user edit string
			// Uncomment below once we have descriptions
			// tmp := shell.EditString(m.GetDescription(id))
			// Shove back updated string
			// m.SetDescription(id, tmp)
		}
	},
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start a todo server.",
	Long:  "Start a todo server that provides a web navigable interface and an API to allow you to view and potentially modify your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		if background {
			// TODO: daemonize
		}
		m, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to acquire resource before starting server: %v", err)
		}
		server.StartServer(m, port, noRestart)
	},
}

var optionCommand = &cobra.Command{
	Use:   "option",
	Short: "Set an option for your todo graph.",
	Long:  "Set an option for your todo graph, affecting future behaviour of todo commands.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Options don't work yet. :(")
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
