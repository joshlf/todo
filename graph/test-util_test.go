package graph

import (
	"reflect"
)

var void = struct{}{}

func makeTestTasks() Tasks {
	//
	//     A
	//    / \
	//   B   C
	//  / \ / \
	// D   E   F
	//
	return Tasks{
		"A": &Task{Id: "A", Dependencies: TaskIDSet{"B": void, "C": void}},
		"B": &Task{Id: "B", Dependencies: TaskIDSet{"D": void, "E": void}},
		"C": &Task{Id: "C", Dependencies: TaskIDSet{"E": void, "F": void}},
		"D": &Task{Id: "D", Dependencies: TaskIDSet{}},
		"E": &Task{Id: "E", Dependencies: TaskIDSet{}},
		"F": &Task{Id: "F", Dependencies: TaskIDSet{}},
	}
}

func makeTestTaskSlice() []Task {
	//
	//     A
	//    / \
	//   B   C
	//  / \ / \
	// D   E   F
	//
	return []Task{
		Task{Id: "A", Dependencies: TaskIDSet{"B": void, "C": void}},
		Task{Id: "B", Dependencies: TaskIDSet{"D": void, "E": void}},
		Task{Id: "C", Dependencies: TaskIDSet{"E": void, "F": void}},
		Task{Id: "D", Dependencies: TaskIDSet{}},
		Task{Id: "E", Dependencies: TaskIDSet{}},
		Task{Id: "F", Dependencies: TaskIDSet{}},
	}
}

func tasksEqual(a, b Tasks) bool {
	return reflect.DeepEqual(a, b)
}
