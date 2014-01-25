package json

import (
	"encoding/json"
	"github.com/joshlf13/todo"
)

type file struct {
	Tasks []Task `json:"tasks"`
}

func (f file) toTodoTasks() []todo.Task {
	t := make([]todo.Task, 0)
	for _, task := range f.Tasks {
		t = append(t, task.toTodoTask())
	}
	return t
}

func fromTodoTasks(t []todo.Task) file {
	f := file{make([]Task, 0)}
	for _, task := range t {
		f.Tasks = append(f.Tasks, fromTodoTask(task))
	}
	return f
}

type Task struct {
	Id           string   `json:"id"`
	Start        uint64   `json:"start"`
	End          uint64   `json:"end"`
	Completed    bool     `json:"completed"`
	Dependencies []string `json:"dependencies"`
}

func (t Task) toTodoTask() todo.Task {
	return todo.Task{
		Id:           todo.TaskID(t.Id),
		Start:        t.Start,
		End:          t.End,
		Completed:    t.Completed,
		Dependencies: toTaskIDMap(t.Dependencies),
	}
}

func fromTodoTask(t todo.Task) Task {
	return Task{
		Id:           string(t.Id),
		Start:        t.Start,
		End:          t.End,
		Completed:    t.Completed,
		Dependencies: fromTaskIDMap(t.Dependencies),
	}
}

func toTaskIDMap(s []string) map[todo.TaskID]struct{} {
	m := make(map[todo.TaskID]struct{})
	for _, id := range s {
		m[todo.TaskID(id)] = struct{}{}
	}
	return m
}

func fromTaskIDMap(m map[todo.TaskID]struct{}) []string {
	s := make([]string, 0)
	for id := range m {
		s = append(s, string(id))
	}
	return s
}

func Unmarshal(data []byte) ([]todo.Task, error) {
	f := file{make([]Task, 0)}
	err := json.Unmarshal(data, &f)
	if err != nil {
		return nil, err
	}
	return f.toTodoTasks(), nil
}

func Marshal(t []todo.Task) ([]byte, error) {
	f := fromTodoTasks(t)
	return json.Marshal(f)
}
