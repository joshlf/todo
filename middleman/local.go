package middleman

import (
	"github.com/joshlf13/todo/graph"
)

// Implements Middleman
type Local struct {
	tasks graph.Tasks
}

func NewLocal(tasks graph.Tasks) Local {
	return Local{tasks}
}

func (l Local) AddTask(t graph.Task) (graph.TaskID, error) {
	for id := range t.Dependencies {
		if _, ok := l.tasks[id]; !ok {
			return graph.TaskID(""), newInvalidRefError(id)
		}
	}
	// Make sure the caller can't
	// muck with our data later
	t = t.Copy()
	t.Id = graph.GenerateID()
	l.tasks[t.Id] = &t
	return t.Id, nil
}

func (l Local) AddDependency(from, to graph.TaskID) error {
	f, ok := l.tasks[from]
	if !ok {
		return newInvalidRefError(from)
	}
	if _, ok := l.tasks[to]; !ok {
		return newInvalidRefError(to)
	}
	f.Dependencies.Add(to)
	return nil
}

func (l Local) GetDependencies(id graph.TaskID) (graph.Tasks, error) {
	task, ok := l.tasks[id]
	if !ok {
		return nil, newInvalidRefError(id)
	}
	return task.GetDependenciesTasks(l.tasks).PruneDependencies().Copy(), nil
}

func (l Local) MarkCompleted(id graph.TaskID) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	task.Completed = true
	return nil
}

// Only mark completed if all dependencies are completed
func (l Local) MarkCompletedVerify(id graph.TaskID) (bool, error) {
	task, ok := l.tasks[id]
	if !ok {
		return false, newInvalidRefError(id)
	}
	for _, task = range task.GetDependenciesTasks(l.tasks) {
		if !task.Completed {
			return false, nil
		}
	}
	task.Completed = true
	return true, nil
}

// Mark completed and force mark all dependencies
// as completed recursively
func (l Local) MarkCompletedRecursive(id graph.TaskID) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	task.Completed = true
	l.tasks.DependencyTree(id).Map(func(id graph.TaskID, t *graph.Task) {
		t.Completed = true
	})
	return nil
}
