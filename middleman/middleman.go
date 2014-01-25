package middleman

import (
	"fmt"
	"github.com/joshlf13/todo/graph"
	"time"
)

type invalidTimeError string
type invalidRefError string

func (i invalidTimeError) Error() string { return string(i) }
func (i invalidRefError) Error() string  { return string(i) }

func newInvalidRefError(ref graph.TaskID) invalidRefError {
	return invalidRefError(fmt.Sprintf("Invalid reference: %v", ref))
}

// Returns whether the error signifies
// that a reference was invalid
func IsInvalidRefError(e error) bool {
	_, ok := e.(invalidRefError)
	return ok
}

type Middleman interface {
	GetTask(id graph.TaskID) (graph.Task, error)

	AddTask(t graph.Task) (graph.TaskID, error)

	SetStartTime(id graph.TaskID, start time.Time) error
	SetEndTime(id graph.TaskID, end time.Time) error

	AddDependency(from, to graph.TaskID) error

	// Get all dependencies of id
	GetDependencies(id graph.TaskID) (graph.Tasks, error)

	// Get all tasks which are both unblocked and uncompleted
	GetUnblocked() (graph.Tasks, error)

	// Get all tasks which are dependencies of id,
	// and are both unblocked and uncompleted
	GetUnblockedDependencies(id graph.TaskID) (graph.Tasks, error)
	MarkCompleted(id graph.TaskID, obliterate bool) error
	MarkCompletedVerify(id graph.TaskID, obliterate bool) (bool, error)
	MarkCompletedRecursive(id graph.TaskID, obliterate bool) error

	Close() error
}
