package graph

import (
	"fmt"
)

func (t TaskIDSet) String() string {
	if t.Len() == 0 {
		return "[]"
	}
	s := "["
	for task := range t {
		s += string(task) + " "
	}
	b := []byte(s)
	b[len(b)-1] = ']'
	return string(b)
}
func (t *Task) String() string {
	return fmt.Sprintf("{Id:%s Start:%d End:%d Completed:%v Dependencies:%v}", t.Id, t.Start, t.End, t.Completed, t.Dependencies)
}

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
