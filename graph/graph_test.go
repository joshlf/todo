package graph

import (
	"math"
	"math/rand"
	"testing"
)

func TestTaskPruneDependencies(t *testing.T) {
	tasks := makeTestTasks()
	delete(tasks, TaskID("B"))
	newA := tasks[TaskID("A")].PruneDependencies(tasks)
	expectedDependencies := MakeTaskIDSet()

	// Make sure that dependencies are properly pruned
	expectedDependencies.Add(TaskID("C"))
	if !newA.Dependencies.Equal(expectedDependencies) {
		t.Errorf("Expected dependencies %v; got %v", expectedDependencies, newA.Dependencies)
	}

	// Make sure that the old graph isn't affected
	expectedDependencies.Add(TaskID("B"))
	if !tasks[TaskID("A")].Dependencies.Equal(expectedDependencies) {
		t.Errorf("Expected dependencies %v; got %v", expectedDependencies, newA.Dependencies)
	}
}

func TestTasksPruneDependencies(t *testing.T) {
	tasks := makeTestTasks()
	delete(tasks, TaskID("B"))
	newTasks := tasks.PruneDependencies()
	newA := newTasks[TaskID("A")]
	expectedDependencies := MakeTaskIDSet()

	// Make sure that dependencies are properly pruned
	expectedDependencies.Add(TaskID("C"))
	if !newA.Dependencies.Equal(expectedDependencies) {
		t.Errorf("Expected dependencies %v; got %v", expectedDependencies, newA.Dependencies)
	}

	// Make sure that the old graph isn't affected
	expectedDependencies.Add(TaskID("B"))
	if !tasks[TaskID("A")].Dependencies.Equal(expectedDependencies) {
		t.Errorf("Expected dependencies %v; got %v", expectedDependencies, newA.Dependencies)
	}
}

func TestTimeSet(t *testing.T) {
	task := Task{}
	if task.StartTimeSet() {
		t.Errorf("Start time should be reported as not set (integer value: %v)", task.Start)
	}
	if !task.EndTimeSet() {
		t.Errorf("End time should be reported as set (integer value: %v)", task.End)
	}
	task.Start = 1
	task.End = math.MaxInt64
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
		orig := rand.Int63()
		task.Start = orig
		tm := task.StartTime()
		task.SetStartTime(tm)
		if task.Start != orig {
			t.Errorf("Converted %v -> %v -> %v", orig, tm, task.Start)
		}

		orig = rand.Int63()
		task.End = orig
		tm = task.EndTime()
		task.SetEndTime(tm)
		if task.End != orig {
			t.Errorf("Converted %v -> %v -> %v", orig, tm, task.End)
		}
	}
}
