package graph

import (
	"testing"
)

func TestTaskIDSet(t *testing.T) {
	set := MakeTaskIDSet()
	if l := set.Len(); l != 0 {
		t.Errorf("New TaskIDSet should have length 0; has %v", l)
	}
}

func TestTaskIDSetAdd(t *testing.T) {
	set := MakeTaskIDSet()
	set.Add("")
	if !set.Contains("") {
		t.Errorf("TaskIDSet should contain \"\"")
	}
	set.Remove("")
	if set.Contains("") {
		t.Errorf("TaskIDSet shouldn't contain \"\"")
	}

	set = MakeTaskIDSet()
	if set.Contains("-") {
		t.Errorf("TaskIDSet shouldn't contain \"-\"")
	}
}

func TestTaskIDSetEqual(t *testing.T) {
	s1, s2 := MakeTaskIDSet(), MakeTaskIDSet()
	if !s1.Equal(s2) {
		t.Errorf("%v and %v should be equal", s1, s2)
	}
	s1.Add("A")
	s2.Add("A")
	s1.Add("B")
	s2.Add("B")
	if !s1.Equal(s2) {
		t.Errorf("%v and %v should be equal", s1, s2)
	}

	s1.Add("C")
	if s1.Equal(s2) {
		t.Errorf("%v and %v should not be equal", s1, s2)
	}

	s2.Add("D")
	if s1.Equal(s2) {
		t.Errorf("%v and %v should not be equal", s1, s2)
	}
}

func TestTaskIDSetCopy(t *testing.T) {
	s := MakeTaskIDSet()
	s.Add("A")
	s1 := s.Copy()
	if !s.Equal(s1) {
		t.Errorf("%v and %v should be equal", s, s1)
	}

	s1.Remove("A")
	if !s.Contains("A") {
		t.Errorf("TaskIDSet should contain \"A\"")
	}
}

func TestTaskIDSetSub(t *testing.T) {
	set1 := MakeTaskIDSet()
	set1.Add("A")
	set1.Add("B")
	set1.Add("C")

	set2 := MakeTaskIDSet()
	set2.Add("B")
	set3 := set1.Sub(set2)

	expect := MakeTaskIDSet()
	expect.Add("A")
	expect.Add("C")
	if !set3.Equal(expect) {
		t.Errorf("%v should be equal to %v", set3, expect)
	}
}
