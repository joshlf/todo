package graph

// TODO this implementation could probably be improved
// to be a bit more efficient.

// Returns set of blocked tasks
func (t Tasks) Blocked() Tasks {
	return t.Filter(func(id TaskID, task *Task) bool {
		return !unblocked(id, task, t)
	})
}

// Returns set of unblocked tests
func (t Tasks) Unblocked() Tasks {
	return t.Filter(func(id TaskID, task *Task) bool {
		return unblocked(id, task, t)
	})
}

// determine if a task is not blocked (has no unfinished dependencies)
func unblocked(id TaskID, task *Task, t Tasks) bool {
	for taskid := range task.Dependencies {
		if !t[taskid].Completed {
			return false
		}
	}

	return true
}
