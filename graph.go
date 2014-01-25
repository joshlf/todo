package todo

import (
	"math"
	"time"
)

type TaskID string

type Task struct {
	Id         TaskID
	End, Start uint64
	Depends    []TaskID
}

type Tasks map[TaskID]*Task

type TodoList struct {
	Tasks
}

func (t *Task) StartTime() time.Time {
	return time.Unix(int64(t.Start), 0)
}

func (t *Task) EndTime() time.Time {
	return time.Unix(int64(t.Start), 0)
}

func (t *Task) StartTimeSet() bool {
	return t.Start != 0
}

func (t *Task) EndTimeSet() bool {
	return t.End != math.MaxUint64
}
