package todo

func Filter(t Tasks, f func(TaskID, *Task)bool) Tasks {
	var m = make(Tasks)
	for id, task := range t {
		if f(id, task) {
			m[id] = task
		}
	}
	return m
}
