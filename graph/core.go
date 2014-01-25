package graph

func (t Tasks) Filter(f func(TaskID, *Task) bool) Tasks {
	u := MakeTasks()
	for id, task := range t {
		if f(id, task) {
			u[id] = task
		}
	}
	return u
}

func (t Tasks) Uncompleted() Tasks {
	return t.Filter(func(id TaskID, task *Task) bool {
		return !task.Completed
	})
}

func (t Tasks) Completed() Tasks {
	return t.Filter(func(id TaskID, task *Task) bool {
		return task.Completed
	})
}

// id must be in t right now!
// Returns set of tasks that depend on id
func (t Tasks) Dependents(id TaskID) Tasks {
	return t.Filter(func(tid TaskID, task *Task) bool {
		_, ok := t[tid].Dependencies[id]
		return ok
	})
}

// Returns true if graph is acyclic
func (t Tasks) Acyclic() bool {
	_, ok := TopoSort(t)
	return ok
}
