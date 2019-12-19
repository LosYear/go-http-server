package main

import (
	"./app"
	"./server"
	"encoding/json"
	"fmt"
)

func main() {
	application := app.NewApp()
	httpServer := server.InitServer(":8080")


	httpServer.AddRoute("GET", "/items", func(request server.Request) server.Response {
		items := application.All()
		strRepr, _ := json.Marshal(items)

		response := server.NewResponseWithBody(200, "OK", strRepr)
		response.Header["Content-Type"] = []string{"application/json"}
		return *response
	})

	httpServer.AddRoute("POST", "/add_item", func(request server.Request) server.Response {
		application.Add(string(request.Body))

		response := server.NewResponse(200, "OK")
		response.Header["Content-Type"] = []string{"application/json"}
		return *response
	})

	httpServer.AddRoute("GET", "*", server.ServeStatic("./static/dist"))

	fmt.Println("Server listening on port", httpServer.Address)

	httpServer.Run()
}
