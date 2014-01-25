package graph

import (
	"testing"
)

func TestGraphFromTasks(t *testing.T) {
	tasks := []Task{
	Task{Id: "A", Dependencies: TaskIDSet{"B": void, "C": void}},
	Task{Id: "B", Dependencies: TaskIDSet{"D": void, "E": void}},
	Task{Id: "C", Dependencies: TaskIDSet{"E": void, "F": void}},
	Task{Id: "D", Dependencies: TaskIDSet{}},
	Task{Id: "E", Dependencies: TaskIDSet{}},
	Task{Id: "F", Dependencies: TaskIDSet{}}}

	t1, _ := GraphFromTasks(tasks)
	t2 := makeTestTasks()
	if !tasksEqual(t1, t2) {
		t.Errorf("Tasks are not equal.")
	}
}
