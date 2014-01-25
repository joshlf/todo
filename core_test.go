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
