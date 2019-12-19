package server

import (
	"bufio"
	"log"
	"net"
	"net/textproto"
	"strings"
)

type Server struct {
	Address string

	listener        *net.TCPListener
	routes          []Route
	NotFoundHandler Route
}

func (server Server) findRoute(request Request) *Route {
	for _, route := range server.routes {
		if route.method == request.Method && (route.path == request.URI || route.path == "*") {
			return &route
		}
	}

	return &server.NotFoundHandler

}

func (server Server) requestHandler(conn net.Conn) {
	r := textproto.NewReader(bufio.NewReader(conn))

	headers := make([]string, 0)

	for {
		str, err := r.ReadLine()

		checkError(err)

		if str == "" {
			// End of headers since headers and body (if exists are separated with empty line)
			break
		}

		headers = append(headers, str)
	}

	request := ParseRequest(headers)
	route := server.findRoute(*request)

	if request.Method == "PUT" || request.Method == "POST" {
		body := make([]byte, request.ContentLength, request.ContentLength)

		_, err := r.R.Read(body[0:request.ContentLength])
		checkError(err)
		request.Body = body
	}

	response := route.handler(*request)
	response = response.setServerHeaders()

	w := textproto.NewWriter(bufio.NewWriter(conn))
	err := w.PrintfLine("%s", response.toString())

	checkError(err)

	err = conn.Close()
	checkError(err)
}

func (server *Server) AddRoute(method string, url string, handler RequestHandler) {
	method = strings.ToUpper(method)

	if method != "GET" && method != "POST" && method != "PUT" {
		return
	}

	server.routes = append(server.routes, Route{method: method, path: url, handler: handler})
}

func (server Server) Run() {
	for {
		conn, err := server.listener.Accept()

		if err != nil {
			return
		}

		log.Println("Accepted new client", conn.RemoteAddr())

		go server.requestHandler(conn)
	}
}
