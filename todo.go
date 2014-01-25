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
		defer cleanupCall(m)
		t := graph.Task{}
		id, err := m.AddTask(t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add new task to resource: %v", err)
			return
		}
		if dep != "" {
			m.AddDependency(id, graph.TaskID(dep))
		}
		etime := parse(end)
		err = m.SetEndTime(id, etime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to set time to %v: %v\n", etime, err)
			return
		}
		stime := parse(start)
		err = m.SetStartTime(id, stime)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to set start time to %v: %v\n", stime, err)
			return
		}
		m.SetDescription(id, args[0])
		// m.SetRunCmd(id, runcmd)
		m.SetWeight(id, float64(weight))

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
			return
		}
		defer cleanupCall(m)
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
				err = m.SetEndTime(id, parse(end))
			}

			if start != DATE_START {
				err = m.SetStartTime(id, parse(start))
			}

			if runcmd != "" {
				// TODO: Implement
			}

			if weight != 1 {
				err = m.SetWeight(id, float64(weight))
			}
		}
	},
}

var finishCommand = &cobra.Command{
	Use:   "finish ref",
	Short: "Finish one or more tasks in the graph.",
	Long:  "Finish one or more tasks in the graph, and potentially their dependencies.",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error acquiring resource to finish: %v\n", err)
			return
		}
		defer cleanupCall(m)
		// Get the task that this ref refers to.
		// TODO This currently does *not* implement recursive completion. 
		ref := args[0]
		id := graph.TaskID(ref)
		if recursive {
			err = m.MarkCompletedRecursive(id, obliterate)
		} else {
			err = m.MarkCompleted(id, obliterate)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marking ref(s) as completed: %v\n", err)
			return
		}
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
		m, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to acquire resource to get run command: %v", err)
			return
		}
		defer cleanupCall(m)
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
		defer cleanupCall(m)
		unblocked, err := m.GetUnblocked()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to find unblocked commands: %v", err)
			return
		}
		for i, t := range unblocked {
			fmt.Printf("  %v)\n%v\n", i, pretty("    ", t))
		}
	},
}

func pretty(indent string, t *graph.Task) string {
	return fmt.Sprintf("%sDescription: %v\n%sStart: %v\n%sEnd: %v\n%sDependencies: %v\n", indent, t.Description, indent, t.Start, indent, t.End, indent, t.Dependencies)
}

var editCommand = &cobra.Command{
	Use:   "edit ref",
	Short: "Edit one or more tasks' descriptions.",
	Long:  "Edit one or more tasks' descriptions.",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMiddleman(file, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error acquiring resource to edit description: %v\n", err)
			return
		}
		defer cleanupCall(m)
		// Get the task that this ref refers to.
		ref := args[0]
		id := graph.TaskID(ref) // When we have aliases and classes, resolve here.
		// Let user edit string
		var s string
		s, err = m.GetDescription(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error editing description for %v: %v", id, err)
			return
		}
		tmp := shell.EditString(s)
		// Shove back updated string
		if err = m.SetDescription(id, tmp); err != nil {
			fmt.Fprintf(os.Stderr, "Error editing description for %v: %v", id, err)
			return
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
			return
		}
		defer cleanupCall(m)
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
