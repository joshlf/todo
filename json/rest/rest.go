package rest

import (
	"encoding/json"
	"github.com/joshlf13/todo/graph"
	myJson "github.com/joshlf13/todo/json"
)

type Request struct {
	Task graph.Task
}

func (r Request) toLocalRequest() request { return request{myJson.FromGraphTask(r.Task)} }

type request struct {
	Task myJson.Task `json:"task"`
}

func (r request) toGraphRequest() Request { return Request{r.Task.ToGraphTask()} }

type Response struct {
	Info  graph.Task
	Tasks []graph.Task
}

func (r Response) toLocalResponse() response {
	rr := response{myJson.FromGraphTask(r.Info), make([]myJson.Task, 0)}
	for _, t := range r.Tasks {
		rr.Tasks = append(rr.Tasks, myJson.FromGraphTask(t))
	}
	return rr
}

type response struct {
	Info  myJson.Task   `json:"info"`
	Tasks []myJson.Task `json:"tasks"`
}

func (r response) toGraphResponse() Response {
	rr := Response{r.Info.ToGraphTask(), make([]graph.Task, 0)}
	for _, t := range r.Tasks {
		rr.Tasks = append(rr.Tasks, t.ToGraphTask())
	}
	return rr
}

func MarshalRequest(r Request) ([]byte, error) {
	return json.Marshal(r.toLocalRequest())
}

func UnmarshalRequest(data []byte) (Request, error) {
	r := request{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return Request{}, err
	}
	return r.toGraphRequest(), nil
}

func MarshalResponse(r Response) ([]byte, error) {
	return json.Marshal(r.toLocalResponse())
}

func UnmarshalResponse(data []byte) (Response, error) {
	r := response{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return Response{}, err
	}
	return r.toGraphResponse(), nil
}
