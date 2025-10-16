package http

import (
	todo "githab/rtemkin/golangnewkurs092025/restApi/toDo"
	"net/http"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todolist *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todolist,
	}
}

/*
pattern: /tasks
metod: POST
info: JSON in http request body

succeed:
- status code: 201 Create
- response body: JSON represent created task

failed:
- status code: 400, 409, 500 ...
- response body: JSON with error + time now
*/

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
metod: GET
info: pattern

succeed:
- status code: 200 OK
- response body: JSON represent found task

failed:
- status code: 400, 409, 500 ...
- response body: JSON with error + time now
*/

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks
metod: GET
info: pattern

succeed:
- status code: 200 OK
- response body: JSON represent found task

failed:
- status code: 400, 500 ...
- response body: JSON with error + time now
*/

func (h *HTTPHandlers) HandleAllGetTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks?completed=true
metod: GET
info: query params

succeed:
- status code: 200 OK
- response body: JSON represent found task

failed:
- status code: 400, 500 ...
- response body: JSON with error + time now
*/
func (h *HTTPHandlers) HandleAllUncompletedTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
metod: PATCH
info: pattern + JSON in request body

succeed:
- status code: 200 OK
- response body: JSON represent found task

failed:
- status code: 400, 500 ...
- response body: JSON with error + time now
*/

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
metod: DELETE
info: pattern

succeed:
- status code: 204 No Content
- response body: -

failed:
- status code: 400, 404, 500 ...
- response body: JSON with error + time now
*/

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
