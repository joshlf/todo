package graph

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	tasks := MakeTasks()
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

func TestMerge(t *testing.T) {
	tasks := makeTestTasks()
	merged := tasks.Merge(tasks, tasks)
	if !tasksEqual(tasks, merged) {
		t.Errorf("%v and %v should be equal", tasks, merged)
	}

	tasks = makeTestTasks()
	rightBranch := tasks.DependencyTree(TaskID("B"))
	leftBranch := tasks.DependencyTree(TaskID("C"))
	merged = rightBranch.Merge(leftBranch)
	merged[TaskID("A")] = tasks[TaskID("A")]
	if !tasksEqual(tasks, merged) {
		t.Errorf("%v and %v should be equal", tasks, merged)
	}
}

func TestDependencyTree(t *testing.T) {
	tasks := makeTestTasks()
	tasks1 := tasks.DependencyTree(TaskID("A"))
	if !tasksEqual(tasks, tasks1) {
		t.Errorf("%v and %v should be equal", tasks, tasks1)
	}

	tasks2 := tasks.DependencyTree(TaskID("B"))
	if _, ok := tasks2[TaskID("A")]; ok {
		t.Errorf("Tasks shouldn't have task \"A\"")
	}
	if _, ok := tasks2[TaskID("C")]; ok {
		t.Errorf("Tasks shouldn't have task \"A\"")
	}

	// Test to see if it will infinitely recurse
	tasks[TaskID("B")].Dependencies.Add(TaskID("A"))
	tasks.DependencyTree(TaskID("A"))
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
