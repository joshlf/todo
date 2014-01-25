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
	t["A"] = &Task{Id: "A", Dependencies: Dependencies(TaskIDSet{"B": void, "C": void})}
	t["B"] = &Task{Id: "B", Dependencies: Dependencies(TaskIDSet{"D": void, "E": void})}
	t["C"] = &Task{Id: "C", Dependencies: Dependencies(TaskIDSet{"E": void, "F": void})}
	t["D"] = &Task{Id: "D", Dependencies: Dependencies(TaskIDSet{})}
	t["E"] = &Task{Id: "E", Dependencies: Dependencies(TaskIDSet{})}
	t["F"] = &Task{Id: "F", Dependencies: Dependencies(TaskIDSet{})}
	return t
}

func tasksEqual(a, b Tasks) bool {
	return reflect.DeepEqual(a, b)
}
