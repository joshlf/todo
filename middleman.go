package main

import (
	"fmt"
	"github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/middleman"
	"os"
)

// file == true => resource is a filepath
// file == false => resource is a domain name
func getMiddleman(resource string, file bool) (middleman.Middleman, error) {
	if file {
		t, err := jsonFileToTasks(resource)
		if err != nil {
			return nil, err
		}
		return middleman.NewLocal(t, func(t graph.Tasks) error {
			return tasksToJSONFile(t, resource)
		}), nil
	} else {
		return middleman.NewRemote(resource), nil
	}
}

func cleanupCall(m middleman.Middleman) {
	if m == nil {
		return
	}
	err := m.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to close resource cleanly: %v\n", err)
	}
}
