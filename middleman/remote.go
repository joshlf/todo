package middleman

import (
	"github.com/joshlf13/todo/graph"
	"time"
)

// Implements Middleman
type remote struct {
	domain string
}

func NewRemote(domain string) Middleman {
	return remote{domain}
}

func (r remote) GetTask(id graph.TaskID) (graph.Task, error) {
	return graph.Task{}, nil
}

func (r remote) AddTask(t graph.Task) (graph.TaskID, error) {
	return "", nil
}

func (r remote) SetStartTime(id graph.TaskID, start time.Time) error {
	return nil
}

func (r remote) SetEndTime(id graph.TaskID, end time.Time) error {
	return nil
}

func (r remote) AddDependency(from, to graph.TaskID) error {
	return nil
}

func (r remote) GetDependencies(id graph.TaskID) (graph.Tasks, error) {
	return nil, nil
}

func (r remote) GetUnblocked() (graph.Tasks, error) {
	return nil, nil
}

func (r remote) GetUnblockedDependencies(id graph.TaskID) (graph.Tasks, error) {
	return nil, nil
}

func (r remote) MarkCompleted(id graph.TaskID, obliterate bool) error {
	return nil
}

// Only mark completed if all dependencies are completed
func (r remote) MarkCompletedVerify(id graph.TaskID, obliterate bool) (bool, error) {
	return true, nil
}

// Mark completed and force mark all dependencies
// as completed recursively
func (r remote) MarkCompletedRecursive(id graph.TaskID, obliterate bool) error {
	return nil
}

func (r remote) Close() error {
	return nil
}
