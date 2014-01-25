package tests

import (
	"github.com/joshlf13/todo"
	"math/rand"
	"strconv"
)

func MakeTestTasks() todo.Tasks {
	t := make(todo.Tasks)

	n := 5
	tt := make([]todo.TaskID, n)
	for j, _ := range rand.Perm(n) {
		for i, _ := range rand.Perm(n) {
			task := new(todo.Task)
			t[todo.TaskID(strconv.Itoa(i+(j*n)))] = task
			tt[i] = todo.TaskID(strconv.Itoa(i + (j * n)))
			task.Id = todo.TaskID(strconv.Itoa(i + (j * n)))
			task.End = uint64(rand.Int63())
			task.Start = uint64(rand.Int63())
			task.Completed = (rand.Intn(2) == 0)
			task.Dependencies = randCol(tt, n)
		}
		tt = make([]todo.TaskID, n)
	}

	return t
}

func randCol(tt []todo.TaskID, n int) map[todo.TaskID]struct{} {
	m := make(map[todo.TaskID]struct{})
	for i := range rand.Perm(n) {
		if rand.Intn(2) == 0 {
			m[tt[i]] = struct{}{}
		}
	}
	return m
}
