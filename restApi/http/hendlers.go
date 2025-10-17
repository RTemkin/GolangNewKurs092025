package http

import (
	"encoding/json"
	"errors"
	"fmt"
	todo "githab/rtemkin/golangnewkurs092025/restApi/toDo"
	"net/http"
	"time"
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
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
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

