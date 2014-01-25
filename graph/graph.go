package graph

import (
	"math"
	"time"
)

type TaskID string

type Task struct {
	Id           TaskID
	Start, End   int64
	Completed    bool
	Dependencies TaskIDSet
}

type Tasks map[TaskID]*Task

type TodoList struct {
	Tasks
}

func (t *Task) StartTime() time.Time {
	return time.Unix(t.Start, 0)
}

func (t *Task) EndTime() time.Time {
	return time.Unix(t.End, 0)
}

func (t *Task) StartTimeSet() bool {
	return t.Start != 0
}

func (t *Task) EndTimeSet() bool {
	return t.End != math.MaxInt64
}

func (t *Task) SetStartTime(tm time.Time) {
	t.Start = tm.Unix()
}

func (t *Task) SetEndTime(tm time.Time) {
	t.End = tm.Unix()
}
