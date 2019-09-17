package webserve

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/truauth/truauth/pkg/minify"
)

// Init creates byte array of all serveable files
func Init(fileNames ...string) Files {
	files := Files{}

	// todo: remove template brackets on bytes init.
	for _, fileName := range fileNames {
		path := fmt.Sprintf("static/%s.html", fileName)
		data, err := ioutil.ReadFile(path)

		if err != nil {
			return nil
		}

		minifiedData := minify.FromBytes(&data)

		template := template.New(fileName)
		template, err = template.Parse(minifiedData)

		if err != nil {
			panic(err)
		}

		byteData := []byte(minifiedData)

		files[fileName] = &File{
			Data:     &byteData,
			Template: template,
		}
	}

	return files
}

// RetrieveFile retrieves a byte array from memory
func (files Files) RetrieveFile(name string) *File {
	return files[name]
}

// ExecuteBytes exectutes the file bytes to the response writer
func (file File) ExecuteBytes(w http.ResponseWriter) {
	w.Write(*file.Data)
	w.WriteHeader(http.StatusFound)
}
