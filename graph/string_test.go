package graph

import (
	"fmt"
)

func ExampleTasks() {
	fmt.Println(makeTestTasks())
	// Output:
	// 	[{Id:A Start:0 End:0 Completed:false Dependencies:[B C]}
	// {Id:B Start:0 End:0 Completed:false Dependencies:[D E]}
	// {Id:C Start:0 End:0 Completed:false Dependencies:[E F]}
	// {Id:D Start:0 End:0 Completed:false Dependencies:[]}
	// {Id:E Start:0 End:0 Completed:false Dependencies:[]}
	// {Id:F Start:0 End:0 Completed:false Dependencies:[]}]
}
