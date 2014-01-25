package middleman

import (
	"fmt"
	"github.com/joshlf13/todo/graph"
)

type invalidRefError string

func (i invalidRefError) Error() string { return string(i) }

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
	AddDependency(from, to graph.TaskID) error
	GetDependencies(id graph.TaskID) (graph.Tasks, error)
	MarkCompleted(id graph.TaskID) error
	MarkCompletedVerify(id graph.TaskID) (bool, error)
	MarkCompletedRecursive(id graph.TaskID) error
}
