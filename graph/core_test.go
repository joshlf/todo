package graph

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	tasks := make(Tasks)
	for i := 0; i < 100; i++ {
		id := TaskID(fmt.Sprint(i))
		tasks[id] = &Task{Id: id}
	}
	tasks = tasks.Filter(func(id TaskID, t *Task) bool {
		n, _ := strconv.Atoi(string(id))
		return n%2 == 0
	})
	for id, _ := range tasks {
		if n, _ := strconv.Atoi(string(id)); n%2 != 0 {
			t.Errorf("Expected only even IDs; got %v", id)
		}
	}
}

func TestAcyclic(t *testing.T) {
	// Graph starts off with no root nodes
	graph := makeTestTasks()
	graph[TaskID("D")].Dependencies.Add(TaskID("A"))
	if graph.Acyclic() {
		t.Errorf("Cyclic graph is marked as acyclic")
	}

	// Graph starts off with root node, but has cycle
	graph = makeTestTasks()
	graph[TaskID("D")].Dependencies.Add(TaskID("B"))
	if graph.Acyclic() {
		t.Errorf("Cyclic graph is marked as acyclic")
	}

	// Graph starts off with root node, but has cycle,
	// and that cycles is from a node to itself
	graph = makeTestTasks()
	graph[TaskID("D")].Dependencies.Add(TaskID("D"))
	if graph.Acyclic() {
		t.Errorf("Cyclic graph is marked as acyclic")
	}

	// Graph has no cycle
	graph = makeTestTasks()
	if !graph.Acyclic() {
		t.Errorf("Acyclic graph is marked as cyclic")
	}

	// Empty graphs have no cycles
	graph = make(Tasks)
	if !graph.Acyclic() {
		t.Errorf("Empty graph marked as cyclic")
	}
}
