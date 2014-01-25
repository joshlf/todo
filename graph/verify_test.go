package graph

import (
	"testing"
)

func TestGraphFromTasks(t *testing.T) {
	t2 := makeTestTasks()
	t1, _ := GraphFromTasks(makeTestTaskSlice())
	if !tasksEqual(t1, t2) {
		t.Errorf("Tasks are not equal.")
	}
}
