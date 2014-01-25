package graph

import (
	"testing"
	"fmt"
	//"github.com/joshlf13/todo/tests"
)


func TestPageRank1(t *testing.T) {
	ws := PageRank1(make(Tasks))
	fmt.Println(ws)
}
