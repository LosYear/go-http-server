package server

type RequestHandler func(request Request) Response

type Route struct {
	method  string
	path    string
	handler RequestHandler
}

