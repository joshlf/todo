package server

import (
	"fmt"
	"github.com/emicklei/go-restful"
	_ "github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/middleman"
	"log"
	_ "net"
	"net/http"
	_ "os"
)

func StartServer(todo middleman.Middleman, port string, noRestart bool) {
	api := APIHandler{
		todo: todo,
	}

	wsContainer := restful.NewContainer()
	api.registerTasks(wsContainer)

	server := &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
