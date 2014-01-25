package graph

import (
	"container/heap"
)

type taskHeap struct {
	t []TaskID
	w WeightMap
}

func (t taskHeap) Len() int { return len(t.t) }

// container/heap is a min heap, so invert the ordering to get a max heap
func (t taskHeap) Less(i, j int) bool { return t.w[t.t[i]] > t.w[t.t[j]] }
func (t taskHeap) Swap(i, j int)      { t.t[i], t.t[j] = t.t[j], t.t[i] }

func (t *taskHeap) Push(x interface{}) {
	t.t = append(t.t, x.(TaskID))
}
func (t *taskHeap) Pop() interface{} {
	old := t.t
	n := len(old)
	x := old[n-1]
	t.t = old[0 : n-1]
	return x
}

// Adapted from http://en.wikipedia.org/wiki/Topological_sorting

// Returns true if the sort succeeded
// or false if the graph contains cycles.
func TopoSort(t Tasks) ([]TaskID, bool) {
	list := make([]TaskID, 0)
	refCounts := make(map[TaskID]uint32, 0)
	for _, task := range t {
		for d := range task.Dependencies {
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

// Returns true if the sort succeeded
// or false if the graph contains cycles.
func TopoSortWeighted(t Tasks, w WeightMap) ([]TaskID, bool) {
	list := make([]TaskID, 0)
	refCounts := make(map[TaskID]uint32, 0)
	for _, task := range t {
		for d := range task.Dependencies {
			refCounts[d]++
		}
	}
	nodes := &taskHeap{make([]TaskID, 0), w}
	heap.Init(nodes)
	for id := range t {
		if refCounts[id] == 0 {
			heap.Push(nodes, id)
		}
	}
	for nodes.Len() > 0 {
		id := heap.Pop(nodes).(TaskID)
		list = append(list, id)
		for d := range t[id].Dependencies {
			refCounts[d]--
			if refCounts[d] == 0 {
				heap.Push(nodes, d)
			}
		}
	}
	if len(list) != len(t) {
		return nil, false
	}
	return list, true
}
