package tests

import (
	"github.com/joshlf13/todo/graph"
	"math/rand"
	"strconv"
)

func MakeTestTasksN(n int) graph.Tasks {
	t := make(graph.Tasks)

	tt := make([]graph.TaskID, n)
	for j, _ := range rand.Perm(n) {
		for i, _ := range rand.Perm(n) {
			task := new(graph.Task)
			t[graph.TaskID(strconv.Itoa(i+(j*n)))] = task
			tt[i] = graph.TaskID(strconv.Itoa(i + (j * n)))
			task.Id = graph.TaskID(strconv.Itoa(i + (j * n)))
			task.End = rand.Int63()
			task.Start = rand.Int63()
			task.Completed = (rand.Intn(2) == 0)
			task.Dependencies = randCol(tt, n)
		}
		tt = make([]graph.TaskID, n)
	}

	return t
}

func MakeTestTasks() graph.Tasks {
	return MakeTestTasksN(5)
}

func randCol(tt []graph.TaskID, n int) map[graph.TaskID]struct{} {
	m := make(map[graph.TaskID]struct{})
	for i := range rand.Perm(n) {
		if rand.Intn(2) == 0 {
			m[tt[i]] = struct{}{}
		}
	}
	return m
}
