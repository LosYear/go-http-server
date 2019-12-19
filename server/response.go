package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Status     string
	StatusCode int
	Protocol   string

	Header Header
	Body   []byte
}

func NewResponse(code int, status string) *Response {
	return &Response{
		Status:     status,
		StatusCode: code,
		Header:     make(Header, 0)}
}

func NewResponseWithBody(code int, status string, body []byte) *Response {
	response := NewResponse(code, status)
	response.Body = body

	return response
}

func NotFound() *Response {
	return NewResponse(404, "Not found")
}

func InternalServerError(message string) *Response {
	return NewResponse(500, "Internal server error: "+message)
}

func (response Response) setServerHeaders() Response {
	response.Protocol = "HTTP/1.1"
	response.Header["Date"] = []string{time.Now().String()}
	response.Header["Server"] = []string{"EduGoServer/0.1"}
	response.Header["Connection"] = []string{"Closed"}
	response.Header["Access-Control-Allow-Origin"] = []string{"*"}
	response.Header["Content-Length"] = []string{strconv.Itoa(len(response.Body))}

	return response
}

func (response Response) toString() string {
	stringRepresentation := fmt.Sprintf("%s %d %s\r\n", response.Protocol, response.StatusCode, response.Status)

	for key, value := range response.Header {
		stringRepresentation += fmt.Sprintf("%s: %s\r\n", key, strings.Join(value, ", "))
	}

	if len(response.Body) > 0 {
		stringRepresentation += "\r\n"
		stringRepresentation += string(response.Body)
	}

	return stringRepresentation
}
