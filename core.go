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
