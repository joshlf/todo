package server

import (
	"github.com/emicklei/go-restful"
	"github.com/joshlf13/todo/graph"
	"github.com/joshlf13/todo/middleman"
	"math"
)

type APIHandler struct {
	todo middleman.Middleman
}

func (api *APIHandler) registerTasks(container *restful.Container) {

	ws := new(restful.WebService)
	ws.
		Path("/todo/task").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(api.allWInfo).
		Doc("Query for all tasks and information").
		Writes(make(graph.Tasks)))

	ws.Route(ws.POST("/").To(api.create).
		Doc("Create new task. Returns ID.").
		Writes(graph.Task{}))

	ws.Route(ws.GET("/{ref}").To(api.someWInfo).
		Doc("Query for tasks and information").
		Param(ws.PathParameter("ref", "A reference to some task(s)").DataType("string")).
		Writes(make(graph.Tasks)))

	ws.Route(ws.PUT("/{ref}").To(api.update).
		Doc("Update task").
		Param(ws.PathParameter("ref", "A reference to some task(s)").DataType("string")))

	ws.Route(ws.DELETE("/{ref}").To(api.finish).
		Doc("Finish a task").
		Param(ws.PathParameter("ref", "A reference to some task(s)").DataType("string")))

	container.Add(ws)
}

func (api *APIHandler) allWInfo(request *restful.Request, response *restful.Response) {

}

func (api *APIHandler) create(request *restful.Request, response *restful.Response) {
	t := graph.Task{
		Id:           graph.GenerateID(),
		End:          math.MaxInt64,
		Dependencies: graph.MakeTaskIDSet(),
	}
	api.todo.AddTask(t)
}

func (api *APIHandler) update(request *restful.Request, response *restful.Response) {
	ref := request.PathParameter("ref")
	ids := []graph.TaskID{graph.TaskID(ref)}
	for _, _ = range ids {

	}
}

func (api *APIHandler) someWInfo(request *restful.Request, response *restful.Response) {

}

func (api *APIHandler) finish(request *restful.Request, response *restful.Response) {
	vs := request.Request.URL.Query()
	ref := graph.TaskID(request.PathParameter("ref"))
	obliterate := len(vs["obliterate"]) > 0 && vs["obliterate"][0] == "true"
	if len(vs["verify"]) > 0 && vs["verify"][0] == "true" {
		api.todo.MarkCompletedVerify(ref, obliterate)
	} else if len(vs["verify"]) > 0 && vs["recursive"][0] == "true" {
		api.todo.MarkCompletedRecursive(ref, obliterate)
	} else {
		api.todo.MarkCompleted(ref, obliterate)
	}
}

/*

GET  /todo/task          - Get all unblocked tasks and their information
POST /todo/task/         - Create task
PUT  /todo/task/{ref}    - Update task
GET  /todo/task/{ref}    - Get all tasks that lead to ref, and all information of ref
DELETE  /todo/task/{ref} - Mark as finished



*/
