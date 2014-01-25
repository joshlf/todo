package todo

import (
	"testing"
)

func TestSubgraph(t *testing.T) {
	tasks := makeTestTasks()
	tasksPrime := Subgraph("B", tasks)
	allowedTasks := Tasks{
		"B": nil,
		"D": nil,
		"E": nil,
	}
	for id, _ := range tasksPrime {
		if _, ok := allowedTasks[id]; !ok {
			t.Errorf("Did not expect %v to be a member of the subgraph", id)
		}
	}
	for id, _ := range allowedTasks {
		if _, ok := tasksPrime[id]; !ok {
			t.Errorf("Expected %v to be a member of the subgraph", id)
		}
	}
}
