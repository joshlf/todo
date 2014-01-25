package todo

import (
	"math"
	"math/rand"
	"testing"
)

func TestTimeSet(t *testing.T) {
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

func TestTimeConvert(t *testing.T) {
	rand.Seed(13040) // So that testing is deterministic
	task := Task{}
	for i := 0; i < 10000; i++ {
		orig := uint64(rand.Int63())
		task.Start = orig
		tm := task.StartTime()
		task.SetStartTime(tm)
		if task.Start != orig {
			t.Errorf("Converted %v -> %v -> %v", orig, tm, task.Start)
		}

		orig = uint64(rand.Int63())
		task.End = orig
		tm = task.EndTime()
		task.SetEndTime(tm)
		if task.End != orig {
			t.Errorf("Converted %v -> %v -> %v", orig, tm, task.End)
		}
	}
}
