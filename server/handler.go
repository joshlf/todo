package server

import (
	"github.com/emicklei/go-restful"
	"github.com/joshlf13/todo/graph"
	json "github.com/joshlf13/todo/json/rest"
	"github.com/joshlf13/todo/middleman"
	"math"
	"net/http"
	"time"
)

func parse(d string) time.Time {
	t, err := time.Parse("15:04:05 1/2/2006", d)
	if err != nil {
		panic("Fix me!")
	} else {
		return t
	}
}

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
		Writes(json.Response{}))

	ws.Route(ws.POST("/").To(api.create).
		Doc("Create new task. Returns ID.").
		Writes(json.Response{}))

	ws.Route(ws.GET("/{ref}").To(api.someWInfo).
		Doc("Query for tasks and information").
		Param(ws.PathParameter("ref", "A reference to some task(s)").DataType("string")).
		Writes(json.Response{}))

	ws.Route(ws.PUT("/{ref}").To(api.update).
		Doc("Update task").
		Param(ws.PathParameter("ref", "A reference to some task(s)").DataType("string")).
		Writes(json.Response{}))

	ws.Route(ws.DELETE("/{ref}").To(api.finish).
		Doc("Finish a task").
		Param(ws.PathParameter("ref", "A reference to some task(s)").DataType("string")))

	container.Add(ws)
}

func (api *APIHandler) allWInfo(request *restful.Request, response *restful.Response) {
	// Get available deps
	tasks, err := api.todo.GetUnblocked()
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// Get task
	var t graph.Task
	// Grab some task
	for _, v := range tasks {
		t = *v
		break
	}
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// Write back response
	response.WriteEntity(json.Response{
		Info:  t,
		Tasks: tasks.Values(),
	})
}

func (api *APIHandler) create(request *restful.Request, response *restful.Response) {
	var err error
	var tasks graph.Tasks
	t := graph.Task{
		End:          math.MaxInt64,
		Dependencies: graph.MakeTaskIDSet(),
	}
	request.ReadEntity(&t)
	t.Id, err = api.todo.AddTask(t)

	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	tasks, err = api.todo.GetUnblockedDependencies(t.Id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteEntity(json.Response{
			Info:  t,
			Tasks: tasks.Values(),
		})
	}
}

func (api *APIHandler) update(request *restful.Request, response *restful.Response) {
	ref := request.PathParameter("ref")
	id := graph.TaskID(ref) // TODO: Resolve ref once we have aliases and classes.
	t, err := api.todo.GetTask(id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	request.ReadEntity(&t)

	// Set everything to whatever the new values are
	for dep, _ := range t.Dependencies {
		api.todo.AddDependency(id, dep)
	}

	// err = api.todo.SetEndTime(id, t.End)
	// if err != nil {
	//     response.AddHeader("Content-Type", "text/plain")
	//     response.WriteErrorString(http.StatusInternalServerError, err.Error())
	//     return
	// }

	// err = api.todo.SetStartTime(id, t.Start)
	// if err != nil {
	//     response.AddHeader("Content-Type", "text/plain")
	//     response.WriteErrorString(http.StatusInternalServerError, err.Error())
	//     return
	// }

	err = api.todo.SetDescription(id, t.Description)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// api.todo.SetRunCmd(id, runcmd)
	// api.todo.SetWeight(id, weight)

	tasks, err := api.todo.GetUnblockedDependencies(id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteEntity(json.Response{
		Info:  t,
		Tasks: tasks.Values(),
	})
}

func (api *APIHandler) someWInfo(request *restful.Request, response *restful.Response) {
	ref := request.PathParameter("ref")
	id := graph.TaskID(ref)

	// Get task
	t, err := api.todo.GetTask(id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// Get available deps
	tasks, err := api.todo.GetUnblockedDependencies(id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	// Write back response
	response.WriteEntity(json.Response{
		Info:  t,
		Tasks: tasks.Values(),
	})
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
