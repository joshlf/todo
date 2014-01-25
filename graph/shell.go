package graph

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// contains shell interaction functions

func EditString(s string) string {
	// create temporary file
	f, err := ioutil.TempFile("", "todo-")
	if err != nil {
		panic(fmt.Sprintf("todo: EditString: failed to create temporary file: %v", err))
		// TODO this should be a real error eventually
	}

	_, err = f.WriteString(s)
	if err != nil {
		panic(fmt.Sprintf("todo: EditString: failed to write to temporary file: %v", err))
		// TODO change this later
	}

	f.Close()
	fname := f.Name()

	// call editor
	if err = exec.Command(GetEditor(), fname).Run(); err != nil {
		panic(fmt.Sprintf("todo: EditString: failed to launch editor: %v", err))
		// TODO change this later
	}

	// read file
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(fmt.Sprintf("todo: EditString: failed to grab contents of tmp file: %v", err))
		// TODO change this later
	}

	// delete file
	if err = os.Remove(fname); err != nil {
		panic(fmt.Sprintf("todo: EditString: failed to remove tmp file: %v", err))
		// TODO change this later, and is very hacky!
	}

	return string(b)
}

func GetEditor() string {
	ed := os.Getenv("VISUAL")
	if ed == "" {
		ed = os.Getenv("EDITOR")

		// Fail through to vim if there is no editor specified
		if ed == "" {
			ed = "vim"
		}
	}

	return ed
}
