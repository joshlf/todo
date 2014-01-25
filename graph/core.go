package graph

func Filter(t Tasks, f func(TaskID, *Task) bool) Tasks {
	u := make(Tasks)
	for id, task := range t {
		if f(id, task) {
			u[id] = task
		}
	}
	return u
}

func Uncompleted(t Tasks) Tasks {
	return Filter(t, func(id TaskID, task *Task) bool {
		return !task.Completed
	})
}

func Completed(t Tasks) Tasks {
	return Filter(t, func(id TaskID, task *Task) bool {
		return task.Completed
	})
}

// id must be in t right now!
// Returns set of tasks that depend on id
func Dependents(t Tasks, id TaskID) Tasks {
	return Filter(t, func(tid TaskID, task *Task) bool {
		_, ok := (t[tid].Dependencies)[id]
		return ok
	})
}

// Returns true if graph is acyclic
func Acyclic(t Tasks, root TaskID) bool {
	idset := MakeTaskIDSet()
	for id, _ := range t {
		idset.Add(id)
	}
	return acyclic(t, root, idset)
}

// Recursive helper function for Acyclic
func acyclic(t Tasks, id TaskID, idset TaskIDSet) bool {
	if idset.Contains(id) {
		return false
	}

	idset.Add(id)
	for tid := range t[id].Dependencies {
		if !acyclic(t, tid, idset) {
			return false
		}
	}

	return true
}
