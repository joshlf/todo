package middleman

import (
	"github.com/joshlf13/todo/graph"
	"time"
)

// Implements Middleman
type local struct {
	tasks graph.Tasks
	close func(graph.Tasks) error
}

func NewLocal(tasks graph.Tasks, close func(graph.Tasks) error) Middleman {
	return local{tasks, close}
}

func (l local) GetTask(id graph.TaskID) (graph.Task, error) {
	task, ok := l.tasks[id]
	if !ok {
		return graph.Task{}, newInvalidRefError(id)
	}
	return task.Copy(), nil
}

func (l local) AddTask(t graph.Task) (graph.TaskID, error) {
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

func (l local) SetStartTime(id graph.TaskID, start time.Time) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	if task.EndTime().Before(start) {
		return invalidTimeError("Start time after end time")
	}
	task.SetStartTime(start)
	return nil
}

func (l local) SetEndTime(id graph.TaskID, end time.Time) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	if task.StartTime().After(end) {
		return invalidTimeError("End time before start time")
	}
	task.SetEndTime(end)
	return nil
}

func (l local) SetTimes(id graph.TaskID, start, end time.Time) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	if start.After(end) {
		return invalidTimeError("End time before start time")
	}
	task.SetStartTime(start)
	task.SetEndTime(end)
	return nil
}

func (l local) SetWeight(id graph.TaskID, w float64) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	task.Weight = w
	return nil
}

// Get and set description
func (l local) GetDescription(id graph.TaskID) (string, error) {
	task, ok := l.tasks[id]
	if !ok {
		return "", newInvalidRefError(id)
	}
	return task.Description, nil
}

func (l local) SetDescription(id graph.TaskID, s string) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	task.Description = s
	return nil
}

func (l local) AddDependency(from, to graph.TaskID) error {
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

func (l local) GetDependencies(id graph.TaskID) (graph.Tasks, error) {
	task, ok := l.tasks[id]
	if !ok {
		return nil, newInvalidRefError(id)
	}
	return task.GetDependenciesTasks(l.tasks).PruneDependencies(), nil
}

func (l local) GetUnblocked() (graph.Tasks, error) {
	return l.tasks.Unblocked().Uncompleted().PruneDependencies(), nil
}

func (l local) GetUnblockedDependencies(id graph.TaskID) (graph.Tasks, error) {

	if _, ok := l.tasks[id]; !ok {
		return nil, newInvalidRefError(id)
	}
	return l.tasks.DependencyTree(id).Unblocked().Uncompleted().PruneDependencies(), nil
}

func (l local) MarkCompleted(id graph.TaskID, obliterate bool) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}
	if obliterate {
		delete(l.tasks, id)
		l.tasks.PruneDependenciesMutate()
	} else {
		task.Completed = true
	}
	return nil
}

// Only mark completed if all dependencies are completed
func (l local) MarkCompletedVerify(id graph.TaskID, obliterate bool) (bool, error) {
	task, ok := l.tasks[id]
	if !ok {
		return false, newInvalidRefError(id)
	}
	for _, task = range task.GetDependenciesTasks(l.tasks) {
		if !task.Completed {
			return false, nil
		}
	}

	if obliterate {
		delete(l.tasks, id)
		l.tasks.PruneDependenciesMutate()
	} else {
		task.Completed = true
	}
	return true, nil
}

// Mark completed and force mark all dependencies
// as completed recursively
func (l local) MarkCompletedRecursive(id graph.TaskID, obliterate bool) error {
	task, ok := l.tasks[id]
	if !ok {
		return newInvalidRefError(id)
	}

	depTree := l.tasks.DependencyTree(id)
	if obliterate {
		for id = range depTree {
			delete(l.tasks, id)
		}
		l.tasks.PruneDependenciesMutate()
	} else {
		task.Completed = true
		depTree.Map(func(id graph.TaskID, t *graph.Task) {
			t.Completed = true
		})
	}
	return nil
}

func (l local) Close() error {
	return l.close(l.tasks)
}
