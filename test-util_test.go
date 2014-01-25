package todo

func makeTestTasks() Tasks {
	//
	//     A
	//    / \
	//   B   C
	//  / \ / \
	// D   E   F
	//
	t := make(Tasks)
	t["A"] = &Task{Id: "A", Depends: []TaskID{"B", "C"}}
	t["B"] = &Task{Id: "B", Depends: []TaskID{"D", "E"}}
	t["C"] = &Task{Id: "C", Depends: []TaskID{"E", "F"}}
	t["D"] = &Task{Id: "D", Depends: []TaskID{}}
	t["E"] = &Task{Id: "E", Depends: []TaskID{}}
	t["F"] = &Task{Id: "F", Depends: []TaskID{}}
	return t
}
