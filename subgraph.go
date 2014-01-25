package todo

func Subgraph(id TaskID, t Tasks) Tasks {
	ret := make(Tasks)
	subgraphHelper(id, t, ret)
	return ret
}

func subgraphHelper(id TaskID, orig, new Tasks) {
	task, ok := orig[id]
	if ok {
		new[id] = task
		for _, depID := range task.Depends {
			subgraphHelper(depID, orig, new)
		}
	}
}
