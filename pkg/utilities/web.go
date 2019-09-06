package utilities

import (
	"encoding/json"
	"net/http"
	"time"
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

// CreateCookie creates a http web cookie
func CreateCookie(name string, domain string, path string, expires time.Time, value string) *http.Cookie {
	return &http.Cookie{
		Name:    name,
		Domain:  domain,
		Path:    path,
		Expires: expires,
		Value:   value,
	}
}
