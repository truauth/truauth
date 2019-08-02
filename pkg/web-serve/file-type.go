package webserve

import (
	"text/template"
	"net/http"
	"github.com/truauth/truauth/pkg/utilities"
	"strconv"
	"fmt"
)

// File file data
type File struct {
	Data *[]byte
	Template *template.Template
}

// Files map of files
type Files map[string]*File

// ErrorTemplate template for the error
type ErrorTemplate struct {
	Error bool
	ErrorDescription string
}

type Session struct {
	Error ErrorTemplate
}

func MarshalSession(req *http.Request) *Session {
	err, _ := strconv.ParseBool(utilities.GetParam(req, "error"))

	if len(req.URL.Query()) == 0 {
		return nil
	}

	return &Session{
		ErrorTemplate{
			Error: err,
			ErrorDescription: utilities.GetParam(req, "error_description"),
		},
	}
}

func FromSessionDetails(isError string, errorDescription string) string {
	return fmt.Sprintf("error=%s&error_description=%s", isError, errorDescription)
}

// func (session *Session) ToEncodedURI() string {
// 	return fmt.Sprintf("error=%s&")
// }