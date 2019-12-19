package server

import "net"

func notFoundHandler(request Request) Response {
	return *NotFound()
}

func InitServer(address string) *Server {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	server := Server{Address: address, listener: listener, NotFoundHandler: Route{handler: notFoundHandler}}

	return &server
}
