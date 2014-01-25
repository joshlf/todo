package todo

import (
	"os"
)

// TODO contains shell interaction functions

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
