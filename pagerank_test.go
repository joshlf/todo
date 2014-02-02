package main

import (
	//	"fmt"
	"github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/tests"
	_ "github.com/joshlf13/todo/tests/impl"
	"testing"
)

func TestPageRank1(t *testing.T) {
	g := tests.MakeTestTasksN(100, 100, 10).(graph.Tasks)
	graph.PageRank1(g)
	//fmt.Println(ws)
}

func BenchmarkPageRank1(b *testing.B) {
	g := tests.MakeTestTasksN(b.N, 200, 10).(graph.Tasks)
	b.ResetTimer()
	graph.PageRank1(g)
}
