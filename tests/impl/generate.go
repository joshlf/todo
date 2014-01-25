package impl

import (
	"github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/tests"
	"math/rand"
	"strconv"
)

func MakeTestTasksN(height, width, connectivity int) graph.Tasks {
	t := make(graph.Tasks)

	levels := make([][]*graph.Task, height)
	for i := range levels {
		levels[i] = make([]*graph.Task, width)
		for j := range levels[i] {
			task := &graph.Task{
				Id:           graph.TaskID(strconv.FormatInt(rand.Int63(), 10)),
				Start:        rand.Int63(),
				End:          rand.Int63(),
				Completed:    rand.Int()%2 == 0,
				Dependencies: graph.MakeTaskIDSet(),
			}
			t[task.Id] = task
			levels[i][j] = task
		}
	}

	for depth, level := range levels[:height-1] {
		for _, task := range level {
			connections := rand.Int31n(int32(connectivity))
			for i := 0; i < int(connections); i++ {
				row, col := rand.Int31n(int32(height-depth-1)), rand.Int31n(int32(width))
				task.Dependencies.Add(levels[depth+int(row)+1][col].Id)
			}
		}
	}
	return t
}

// Equivalent to MakeTestTasksN(5, 5, 5)
func MakeTestTasks() graph.Tasks {
	return MakeTestTasksN(5, 5, 5)
}

func init() {
	tests.MakeTestTasksN = func(h, w, c int) interface{} { return MakeTestTasksN(h, w, c) }
	tests.MakeTestTasks = func() interface{} { return MakeTestTasks() }
}
