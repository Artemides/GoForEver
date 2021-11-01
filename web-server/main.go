package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleWare(HandleHome, CheckAuth(), Logging()))
	server.Handle("GET", "/api/task", HandleTask)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Listen()
}
