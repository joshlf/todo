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

func MakeTasks() Tasks { return make(Tasks) }

func (t Tasks) Copy() Tasks {
	return t.MapImmutable(func(id TaskID, task Task) *Task {
		return &task
	})
}

func (t *Task) Copy() *Task {
	newT := *t // Deep copy
	newT.Dependencies = t.Dependencies.Copy()
	return &newT
}

func (t *Task) GetDependenciesTasks(tt Tasks) Tasks {
	ret := MakeTasks()
	for id := range t.Dependencies {
		task, ok := tt[id]
		if ok {
			ret[id] = task
		}
	}
	return ret
}

// Remove all dependencies which are invalid
// references in tt
func (t *Task) PruneDependencies(tasks Tasks) *Task {
	newT := *t // Deep copy
	remove := MakeTaskIDSet()
	for id := range t.Dependencies {
		if _, ok := tasks[id]; !ok {
			remove.Add(id)
		}
	}
	newT.Dependencies = t.Dependencies.Sub(remove)
	return &newT
}

// Remove all dependencies which are invalid
// references in t
func (t Tasks) PruneDependencies() Tasks {
	newT := MakeTasks()
	for id, task := range t {
		newT[id] = task.PruneDependencies(t)
	}
	return newT
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
