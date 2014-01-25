package todo

import (
	"fmt"
	"math"
	"time"
)

type TaskID string

type Dependencies map[TaskID]struct{}

func (d Dependencies) String() string {
	if len(d) == 0 {
		return "[]"
	}
	s := "["
	for t := range d {
		s += string(t) + " "
	}
	b := []byte(s)
	b[len(b)-1] = ']'
	return string(b)
}

type Task struct {
	Id           TaskID
	Start, End   uint64
	Completed    bool
	Dependencies Dependencies
}

func (t *Task) String() string {
	return fmt.Sprintf("{Id:%s Start:%d End:%d Completed:%v Dependencies:%v}", t.Id, t.Start, t.End, t.Completed, t.Dependencies)
}

type Tasks map[TaskID]*Task

func (t Tasks) String() string {
	if len(t) == 0 {
		return "{}"
	}
	fmtStr := "["
	args := make([]interface{}, 0)
	for _, task := range t {
		args = append(args, task.String())
		fmtStr += "%v\n"
	}
	b := []byte(fmtStr)
	b[len(b)-1] = ']'
	fmtStr = string(b)
	return fmt.Sprintf(fmtStr, args...)
}

type TodoList struct {
	Tasks
}

func (t *Task) StartTime() time.Time {
	return time.Unix(int64(t.Start), 0)
}

func (t *Task) EndTime() time.Time {
	return time.Unix(int64(t.End), 0)
}

func (t *Task) StartTimeSet() bool {
	return t.Start != 0
}

func (t *Task) EndTimeSet() bool {
	return t.End != math.MaxUint64
}

func (t *Task) SetStartTime(tm time.Time) {
	t.Start = uint64(tm.Unix())
}

func (t *Task) SetEndTime(tm time.Time) {
	t.End = uint64(tm.Unix())
}
