package main

import (
	"github.com/joshlf13/todo/graph"
	json "github.com/joshlf13/todo/json/file"
	"io/ioutil"
	"os"
)

func jsonFileToTasksDefault() (graph.Tasks, error) { return jsonFileToTasks(file) }
func tasksToJSONFileDefault(t graph.Tasks) error   { return tasksToJSONFile(t, file) }

func jsonFileToTasks(fname string) (graph.Tasks, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	tasks, err := json.Unmarshal(b)
	if err != nil {
		return nil, err
	}
	t, err := graph.GraphFromTasks(tasks)
	if err != nil {
		return nil, err
	}
	return t, err
}

func tasksToJSONFile(t graph.Tasks, fname string) error {
	b, err := json.Marshal(t.Slice())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fname, b, 0)
}
