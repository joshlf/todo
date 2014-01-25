package todo

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// contains shell interaction functions

func EditString(s string) string {
	// create temporary file
	f, err := ioutil.TempFile("", "todo-")
	if err != nil {
		panic("todo: EditString: failed to create temporary file.")
		// TODO this should be a real error eventually
	}

	_, werr := f.WriteString(s)
	if werr != nil {
		panic("todo: EditString: failed to write to temporary file.")
		// TODO change this later
	}

	f.Close()
	fname := f.Name()

	// call editor
	if exec.Command(GetEditor(), fname).Run() != nil {
		panic("todo: EditString: failed to launch editor.")
		// TODO change this later
	}

	// read file
	b, rferr := ioutil.ReadFile(fname)
	if rferr != nil {
		panic("todo: EditString: failed to grab contents of tmp file.")
		// TODO change this later
	}

	// delete file
	if os.Remove(fname) != nil {
		panic("todo: EditString: failed to remove tmp file.")
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
