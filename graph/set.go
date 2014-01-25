package graph

type TaskIDSet map[TaskID]struct{}

func MakeTaskIDSet() TaskIDSet { return make(TaskIDSet) }

func (t TaskIDSet) Add(id TaskID) { t[id] = struct{}{} }

func (t TaskIDSet) Remove(id TaskID) { delete(t, id) }

func (t TaskIDSet) Sub(u TaskIDSet) TaskIDSet {
	v := MakeTaskIDSet()
	for id := range t {
		if !u.Contains(id) {
			v.Add(id)
		}
	}
	return v
}

func (t TaskIDSet) Equal(u TaskIDSet) bool {
	for id := range t {
		if !u.Contains(id) {
			return false
		}
	}
	for id := range u {
		if !t.Contains(id) {
			return false
		}
	}
	return true
}

func (t TaskIDSet) Copy() TaskIDSet {
	u := MakeTaskIDSet()
	for id := range t {
		u.Add(id)
	}
	return u
}

func (t TaskIDSet) GetRandom() TaskID { id, _ := t.GetRandomOK(); return id }

func (t TaskIDSet) GetRandomOK() (TaskID, bool) {
	for id := range t {
		return id, true
	}
	return "", false
}

func (t TaskIDSet) Contains(id TaskID) bool { _, ok := t[id]; return ok }

func (t TaskIDSet) Len() int { return len(t) }
