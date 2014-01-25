package graph

import (
	"math"
)

// Sets initial weights to 1
func PageRank1(t Tasks) WeightMap {
	return PageRankN(t, 1.0)
}

// Sets initial weights to n for everything
func PageRankN(t Tasks, n float64) WeightMap {
	ws := make(WeightMap)
	for id := range t {
		ws[id] = n
	}
	return PageRank(t, ws)
}

func PageRank(t Tasks, ws WeightMap) WeightMap {
	epsilon := 0.05 // TODO this is arbitrary. Change?
	d := 0.5        // TODO this is arbitrary. No idea what to set it to.
	con := (1.0 - d) / float64(len(t))
	wsnew := make(WeightMap)

	dependents := make(map[TaskID]TaskIDSet)
	for id, task := range t {
		for did := range task.Dependencies {
			deps, ok := dependents[did]
			if !ok {
				deps = MakeTaskIDSet()
				dependents[did] = deps
			}
			deps.Add(id)
		}
	}

	for {
		wsum := 0.0
		for id := range t {
			// sum weights of dependents
			var sum float64
			for did := range dependents[id] {
				sum += ws[did]
			}
			wsnew[id] = d*sum + con
			if wsum < math.Abs(wsnew[id]-ws[id]) {
				wsum = math.Abs(wsnew[id] - ws[id])
			}
		}

		if wsum < epsilon {
			break
		}

		ws = wsnew
		wsnew = make(WeightMap)
	}

	return wsnew
}
