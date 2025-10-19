package main

import (
	"fmt"
	"githab/rtemkin/golangnewkurs092025/restApi/http"
	todo "githab/rtemkin/golangnewkurs092025/restApi/toDo"
)

func main() {

	todoList := todo.NewList()

	httpHanlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHanlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("faled to start http server", err)
	}
}
