package server

import (
	"strconv"
	"strings"
)

type Request struct {
	Method        string
	URI           string
	Protocol      string
	Header        Header
	ContentLength uint64

	Body []byte
}

func parseHeaderValue(value string) []string {
	values := make([]string, 0)

	for _, item := range strings.Split(value, ",") {
		values = append(values, strings.Trim(item, " "))
	}

	return values
}

func ParseRequest(headers []string) *Request {
	requestLineInfo := strings.Split(headers[0], " ")

	parsedHeader := make(Header, 0)
	contentLength := uint64(0)

	for _, line := range headers[1:] {
		header := strings.SplitN(line, ":", 2)

		if len(header) != 2 {
			continue
		}

		key := header[0]
		value := header[1]

		parsedHeader[key] = parseHeaderValue(value)

		if key == "Content-Length" {
			contentLength, _ = strconv.ParseUint(strings.Trim(value, " "), 10, 0)
		}
	}

	request := Request{
		Method:        requestLineInfo[0],
		URI:           requestLineInfo[1],
		Protocol:      requestLineInfo[2],
		Header:        parsedHeader,
		ContentLength: contentLength}

	return &request
}
