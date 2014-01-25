package main

import (
	"github.com/joshlf13/todo/middleman"
)

// file == true => resource is a filepath
// file == false => resource is a domain name
func getMiddleman(resource string, file bool) (middleman.Middleman, error) {
	if file {
		t, err := jsonFileToTasks(resource)
		if err != nil {
			return nil, err
		}
		return middleman.NewLocal(t), nil
	} else {
		return middleman.NewRemote(resource), nil
	}
}
