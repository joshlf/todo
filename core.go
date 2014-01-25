package todo

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
func Acyclic(t Tasks) bool {
	idset := MakeTaskIDSet()
	for id, _ := range t {
		idset.Add(id)
	}
	return acyclic(t, idset)
}

func acyclic(t Tasks, idset TaskIDSet) bool {
	for id, task := range t {

	}

	return true
}
