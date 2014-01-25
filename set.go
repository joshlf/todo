package todo

type TaskIDSet map[TaskID]struct{}

func MakeTaskIDSet() TaskIDSet { return make(TaskIDSet) }

func (t TaskIDSet) Add(id TaskID) { t[id] = struct{}{} }

func (t TaskIDSet) Remove(id TaskID) { delete(t, id) }

func (t TaskIDSet) Contains(id TaskID) bool { _, ok := t[id]; return ok }

func (t TaskIDSet) Len() int { return len(t) }
