package graph

import (
	"math"
	"time"
	"fmt"
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
	epsilon := 0.5 // TODO this is arbitrary. Change?
	d := 0.5 // TODO this is arbitrary. No idea what to set it to. 
	con := (1.0 - d) / float64(len(t))
	wsnew := make(WeightSet)

	m := make(map[TaskID]Tasks)
	for id := range t {
		m[id] = Dependents(t, id)
	}

	t1 := time.Now()
	for {
		wsum := 0.0
		for id := range t {
			// sum weights of dependents
			var sum float64
			for did := range m[id] {
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

	t2 := time.Now()
	fmt.Println("body", t2.Sub(t1))
	return wsnew
}
