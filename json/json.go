package json

import (
	"encoding/json"
	"github.com/joshlf13/todo/graph"
)

func Unmarshal(data []byte) (graph.Task, error) {
	t := graph.Task{}
	err := json.Unmarshal(data, &t)
	return t, err
}

func Marshal(t graph.Task) ([]byte, error) {
	return json.Marshal(FromGraphTask(t))
}

type Task struct {
	Id           string   `json:"id"`
	Start        int64    `json:"start"`
	End          int64    `json:"end"`
	Completed    bool     `json:"completed"`
	Dependencies []string `json:"dependencies"`
}

func (t Task) ToGraphTask() graph.Task {
	return graph.Task{
		Id:           graph.TaskID(t.Id),
		Start:        t.Start,
		End:          t.End,
		Completed:    t.Completed,
		Dependencies: toTaskIDMap(t.Dependencies),
	}
}

func FromGraphTask(t graph.Task) Task {
	return Task{
		Id:           string(t.Id),
		Start:        t.Start,
		End:          t.End,
		Completed:    t.Completed,
		Dependencies: fromTaskIDMap(t.Dependencies),
	}
}

func fromTaskIDMap(m map[graph.TaskID]struct{}) []string {
	s := make([]string, 0)
	for id := range m {
		s = append(s, string(id))
	}
	return s
}

func toTaskIDMap(s []string) map[graph.TaskID]struct{} {
	m := make(map[graph.TaskID]struct{})
	for _, id := range s {
		m[graph.TaskID(id)] = struct{}{}
	}
	return m
}
