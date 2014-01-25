package todo

import (
	"testing"
)

func TestBlocked(t *testing.T) {
	g := makeTestTasks()
	g["D"].Completed = true
	g["E"].Completed = true
	uset := Blocked(g)
	_, a := uset["A"]
	_, b := uset["B"]
	_, c := uset["C"]
	_, d := uset["D"]
	_, e := uset["E"]
	_, f := uset["F"]
	if !a {
		t.Errorf("Missing element A")
	}
	if !c {
		t.Errorf("Missing element C")
	}
	if d {
		t.Errorf("D should not be in the graph!")
	}
	if b {
		t.Errorf("B should not be in the graph!")
	}
	if e {
		t.Errorf("E should not be in the graph!")
	}
	if f {
		t.Errorf("F should not be in the graph!")
	}
}

func TestUnblocked(t *testing.T) {
	g := makeTestTasks()
	g["D"].Completed = true
	g["E"].Completed = true
	uset := Unblocked(g)
	_, a := uset["A"]
	_, b := uset["B"]
	_, c := uset["C"]
	_, d := uset["D"]
	_, e := uset["E"]
	_, f := uset["F"]
	if a {
		t.Errorf("A should not be in the graph!")
	}
	if c {
		t.Errorf("C should not be in the graph!")
	}
	if !d {
		t.Errorf("Missing element D")
	}
	if !b {
		t.Errorf("Missing element B")
	}
	if !e {
		t.Errorf("Missing element E")
	}
	if !f {
		t.Errorf("Missing element F")
	}
}
