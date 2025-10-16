package http


type HTTPServer struct{
	httpHandlers *HTTPHandlers
}

//конструктор
func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer{
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

//запуск сервера
func(s *HTTPServer)StartServer()error{
	
}