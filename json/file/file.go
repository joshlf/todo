package file

import (
	"encoding/json"
	"github.com/joshlf13/todo/graph"
	myJson "github.com/joshlf13/todo/json"
)

type file struct {
	Tasks []myJson.Task `json:"tasks"`
}

func (f file) toGraphTasks() []graph.Task {
	t := make([]graph.Task, 0)
	for _, task := range f.Tasks {
		t = append(t, task.ToGraphTask())
	}
	return t
}

func fromGraphTasks(t []graph.Task) file {
	f := file{make([]myJson.Task, 0)}
	for _, task := range t {
		f.Tasks = append(f.Tasks, myJson.FromGraphTask(task))
	}
	return f
}

func Marshal(t []graph.Task) ([]byte, error) {
	f := fromGraphTasks(t)
	return json.MarshalIndent(f, "", " ")
}

func Unmarshal(data []byte) ([]graph.Task, error) {
	f := file{make([]myJson.Task, 0)}
	err := json.Unmarshal(data, &f)
	if err != nil {
		return nil, err
	}
	return f.toGraphTasks(), nil
}
