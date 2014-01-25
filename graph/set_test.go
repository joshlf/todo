package graph

import (
	"testing"
)

func TestTaskIDSet(t *testing.T) {
	set := MakeTaskIDSet()
	if l := set.Len(); l != 0 {
		t.Errorf("New TaskIDSet should have length 0; has %v", l)
	}
	set.Add("")
	if !set.Contains("") {
		t.Errorf("TaskIDSet should contain \"\"")
	}
	set.Remove("")
	if set.Contains("") {
		t.Errorf("TaskIDSet shouldn't contain \"\"")
	}
	if set.Contains("-") {
		t.Errorf("TaskIDSet shouldn't contain \"-\"")
	}
}
