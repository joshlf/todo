package graph

// Adapted from http://en.wikipedia.org/wiki/Topological_sorting

// Returns true if the sort succeeded
// or false if the graph contains cycles.
func TopoSort(t Tasks) ([]TaskID, bool) {
	list := make([]TaskID, 0)
	refCounts := make(map[TaskID]uint32, 0)
	for _, task := range t {
		for d := range task.Dependencies {
			// fmt.Println(task, d)
			refCounts[d]++
		}
	}
	nodes := MakeTaskIDSet()
	for id := range t {
		if refCounts[id] == 0 {
			nodes.Add(id)
		}
	}
	for id, ok := nodes.GetRandomOK(); ok; id, ok = nodes.GetRandomOK() {
		nodes.Remove(id)
		list = append(list, id)
		for d := range t[id].Dependencies {
			refCounts[d]--
			if refCounts[d] == 0 {
				nodes.Add(d)
			}
		}
	}
	if len(list) != len(t) {
		return nil, false
	}
	return list, true
}
