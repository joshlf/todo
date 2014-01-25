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

// Intended for mutation
func (t Tasks) Map(f func(TaskID, *Task)) {
	for id, task := range t {
		f(id, task)
	}
}

// Does not provide access to mutate t
func (t Tasks) MapImmutable(f func(TaskID, Task) *Task) Tasks {
	newT := MakeTasks()
	for id, task := range t {
		newT[id] = f(id, task.Copy())
	}
	return newT
}

// If identically-named nodes in different
// Tasks have conflicting data, the behavior
// of Merge is undefined.
func (t Tasks) Merge(u ...Tasks) Tasks {
	newT := MakeTasks()
	u = append(u, t)
	for _, t := range u {
		for id, task := range t {
			newT[id] = task
		}
	}
	return newT
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

// Returns all dependencies recursively
func (t Tasks) DependencyTree(id TaskID) Tasks {
	newT := MakeTasks()
	dependentTasks(t, id, newT)
	return newT
}

func dependentTasks(t Tasks, id TaskID, newT Tasks) {
	task, ok := t[id]
	if ok {
		// Don't do unnecessary work;
		// also, don't infinitely recurse
		// on cyclic graphs
		if _, ok := newT[id]; ok {
			return
		}

		newT[id] = task
		for id = range task.Dependencies {
			dependentTasks(t, id, newT)
		}
	}
}

// Returns true if graph is acyclic
func (t Tasks) Acyclic() bool {
	_, ok := TopoSort(t)
	return ok
}
