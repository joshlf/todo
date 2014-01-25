package main

import (
	"testing"
	"time"
	"fmt"
	"github.com/joshlf13/todo/tests"
	_"github.com/joshlf13/todo/tests/impl"
	"github.com/joshlf13/todo/graph"
)


func TestPageRank1(t *testing.T) {
	i := tests.MakeTestTasksN(100, 100, 10).(graph.Tasks)
	t1 := time.Now()
	graph.PageRank1(i)
	t2 := time.Now()
	//fmt.Println(ws)
	fmt.Println()
	fmt.Println("Took", t2.Sub(t1))
}
