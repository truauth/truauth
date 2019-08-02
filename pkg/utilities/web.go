package utilities

import (
	"net/http"
	"encoding/json"
)

// GetParam gets a parameter from the desired request
func GetParam(req *http.Request, value string) string {
	if len(req.URL.Query()[value]) > 0 {
		return req.URL.Query()[value][0]
	}

	return ""
}

// DecodeBody decodes a body given a reqest and an interface
func DecodeBody(req *http.Request, model interface{}) {
	decoder := json.NewDecoder(req.Body)

	decoder.Decode(model)
}