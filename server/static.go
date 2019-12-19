package server

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func detectContentType(filename string) string {
	// Detecting content type by file extension is bad :)
	extension := filepath.Ext(filename)

	switch extension {
	case ".html", ".htm":
		return "text/html"
	case ".svg":
		return "image/svg+xml"
	case ".js":
		return "text/javascript"
	case ".css":
		return "text/css"
	}

	return ""
}

func ServeStatic(basedir string) func(request Request) Response {
	return func(request Request) Response {
		filename := basedir + request.URI
		info, err := os.Stat(filename)
		if os.IsNotExist(err) || info.IsDir() {
			return *NotFound()

		}

		fileContent, err := ioutil.ReadFile(filename)

		if err != nil {
			return *InternalServerError(err.Error())
		}


		response := NewResponseWithBody(200, "OK", fileContent)

		response.Header["Content-Type"] = []string{detectContentType(filename)}

		return *response
	}
}
