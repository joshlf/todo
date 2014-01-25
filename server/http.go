package server

import (
    _ "github.com/joshlf13/todo/graph"
    "github.com/joshlf13/todo/middleman"
    _ "fmt"
    _ "net"
    _ "net/http"
    _ "os"
    "github.com/emicklei/go-restful"
)

func StartServer(todo middleman.Middleman, port int, noRestart bool) {
    api := APIHandler{
        todo: todo,
    }

    wsContainer := restful.NewContainer()
    api.registerTasks(wsContainer)
}
