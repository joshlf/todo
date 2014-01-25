package server

import (
	_ "fmt"
	"github.com/emicklei/go-restful"
	_ "github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/middleman"
	_ "net"
	_ "net/http"
	_ "os"
)

func StartServer(todo middleman.Middleman, port int, noRestart bool) {
	api := APIHandler{
		todo: todo,
	}

	wsContainer := restful.NewContainer()
	api.registerTasks(wsContainer)
}
