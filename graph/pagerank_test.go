package graph

import (
	"fmt"
	"github.com/joshlf13/todo/tests"
//	_"github.com/joshlf13/todo/tests/impl"
)

func TestPageRank1(t *testing.T) {
	i := tests.MakeTestTasksN(100, 100, 10)
	ws := PageRank1(i)
	fmt.Println(ws)
}
