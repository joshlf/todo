package todo

import (
	"reflect"
)

type TaskIDSet map[TaskID]struct{}

var void = struct{}{}

func makeTestTasks() Tasks {
	//
	//     A
	//    / \
	//   B   C
	//  / \ / \
	// D   E   F
	//
	t := make(Tasks)
	t["A"] = &Task{Id: "A", Depends: TaskIDSet{"B": void, "C": void}}
	t["B"] = &Task{Id: "B", Depends: TaskIDSet{"D": void, "E": void}}
	t["C"] = &Task{Id: "C", Depends: TaskIDSet{"E": void, "F": void}}
	t["D"] = &Task{Id: "D", Depends: TaskIDSet{}}
	t["E"] = &Task{Id: "E", Depends: TaskIDSet{}}
	t["F"] = &Task{Id: "F", Depends: TaskIDSet{}}
	return t
}

func tasksEqual(a, b Tasks) bool {
	return reflect.DeepEqual(a, b)
}
