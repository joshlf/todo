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

func (l TodoList) ResolveSingle(ref string) (Task, error) {
    var t Task
    // TODO: Implement
    return t, nil
}

func (l TodoList) NewTask() Task {
    var t Task
    // TODO: place into Dependencies
    return t
}

func (t *Task) GetTaskID() string {
    return ""
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

func (t *Task) GetDescription() string {
    // TODO: Implement description
    return ""
}

func (t *Task) SetDescription(desc string) {
    // TODO: Implement description
}

func (t *Task) SetWeight(w int) {
    // TODO: Implement weights
}

func (t *Task) GetRunCmd() string {
    // TODO: Implement run command
    return ""
}

func (t *Task) SetRunCmd(cmd string) {
    // TODO: Implement run command
}

func (t *Task) AddDependencies(deps []string) {
    // TODO: Do this.
}
