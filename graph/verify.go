package graph

import (
	"fmt"
)

func GraphFromTasks(t []Task) (Tasks, error) {
	ts := make(Tasks)
	for i, task := range t {
		ts[task.Id] = &t[i]
	}

	for _, task := range ts {
		for id := range task.Dependencies {
			_, ok := ts[id]
			if !ok {
				return nil, fmt.Errorf("todo: GraphFromTasks: dependency %v does not resolve.", id)
			}
		}
	}

	if !ts.Acyclic() {
		return nil, fmt.Errorf("todo: GraphFromTasks: graph has a cycle!")
	}

	for id, task := range ts {
		if task.Start > task.End {
			return nil, fmt.Errorf("todo: GraphFromTasks: end time of task %v is before start.", id)
		}
	}

	return ts, nil
}
