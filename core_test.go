package todo

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
	tasks = Filter(tasks, func(id TaskID, t *Task) bool {
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
	graph := makeTestTasks()
	graph[TaskID("D")].Dependencies.Add(TaskID("E"))
	if Acyclic(graph, TaskID("A")) {
		t.Errorf("Cyclic graph is marked as acyclic")
	}

	graph = makeTestTasks()
	if !Acyclic(graph, TaskID("A")) {
		t.Errorf("Acyclic graph is marked as cyclic")
	}
}
