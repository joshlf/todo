package todo

import (
	"math"
	"testing"
)

func TestStartTime(t *testing.T) {
	task := Task{}
	if task.StartTimeSet() {
		t.Errorf("Start time should be reported as not set (integer value: %v)", task.Start)
	}
	if !task.EndTimeSet() {
		t.Errorf("End time should be reported as set (integer value: %v)", task.End)
	}
	task.Start = 1
	task.End = math.MaxUint64
	if !task.StartTimeSet() {
		t.Errorf("Start time should be reported as set (integer value: %v)", task.Start)
	}
	if task.EndTimeSet() {
		t.Errorf("End time should be reported as not set (integer value: %v)", task.End)
	}
}
