package graph

import (
	"reflect"
	"testing"
)

func TestTopoSortWeighted(t *testing.T) {
	tasks := makeTestTasks()
	weights := WeightMap{
		TaskID("A"): 1,
		TaskID("B"): 2,
		TaskID("C"): 3,
		TaskID("D"): 4,
		TaskID("E"): 5,
		TaskID("F"): 6,
	}
	target := []TaskID{
		TaskID("A"),
		TaskID("C"),
		TaskID("F"),
		TaskID("B"),
		TaskID("E"),
		TaskID("D"),
	}
	sorted, _ := TopoSortWeighted(tasks, weights)
	if !reflect.DeepEqual(sorted, target) {
		t.Errorf("Expected %v; got %v", target, sorted)
	}
}
