package webserve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/truauth/truauth/pkg/utilities"
)

// File file data
type File struct {
	Data     *[]byte
	Template *template.Template
}

// Files map of files
type Files map[string]*File

// OutTemplate template for the error
type OutTemplate struct {
	Error            bool
	ErrorDescription string
	Script           string
}

type Session struct {
	Error OutTemplate
}

func MarshalSession(req *http.Request) *Session {
	err, _ := strconv.ParseBool(utilities.GetParam(req, "error"))

	if len(req.URL.Query()) == 0 {
		return nil
	}

	return &Session{
		OutTemplate{
			Error:            err,
			ErrorDescription: utilities.GetParam(req, "error_description"),
		},
	}
}

func FromSessionDetails(isError string, errorDescription string) string {
	return fmt.Sprintf("error=%s&error_description=%s", isError, errorDescription)
}

// DevError generates a developer erorr in the javacsript browser console
func DevError(message string) string {
	return fmt.Sprintf("<script> console.error(`Error: Api Service: %s`) </script>", message)
}

// WriteSessionStorage writes the desired josn content to session storage
func WriteSessionStorage(key string, jsonContent interface{}) string {
	jsonStr, err := json.Marshal(jsonContent)
	if err != nil {
		return DevError("Web Serve Error: File Type: Error occured while writing to session storage")
	}

	return fmt.Sprintf("<script> sessionStorage.setItem('%s', '%s') </script> ", key, jsonStr)
}
