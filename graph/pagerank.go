package graph

import (
	"math"
)

type WeightSet map[TaskID] float64

// Sets initial weights to 1
func PageRank1(t Tasks) WeightSet {
	return PageRankN(t, 1.0)
}

// Sets initial weights to n for everything
func PageRankN(t Tasks, n float64) WeightSet {
	ws := make(WeightSet)
	for id := range t {
		ws[id] = n
	}
	return PageRank(t, ws)
}

func PageRank(t Tasks, ws WeightSet) WeightSet {
	epsilon := 0.05 // TODO this is arbitrary. Change?
	d := 0.1 // TODO this is arbitrary. No idea what to set it to. 
	con := (1.0 - d) / float64(len(t))
	wsnew := make(WeightSet)
	for {
		wsum := 0.0
		for id := range t {
			// sum weights of dependents
			var sum float64
			for did := range Dependents(t, id) {
				sum += ws[did]
			}
			wsnew[id] = d*sum + con
			if wsum < math.Abs(wsnew[id] - ws[id]) {
				wsum = math.Abs(wsnew[id] - ws[id])
			}
		}

		if wsum < epsilon {
			break
		}

		ws = wsnew
		wsnew = make(WeightSet)
	}

	return wsnew
}
