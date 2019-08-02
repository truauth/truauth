package webserve

import (
	"fmt"
	"io/ioutil"
	"text/template"
	
	"github.com/truauth/truauth/pkg/minify"
)

// Init creates byte array of all serveable files
func Init(fileNames ...string) Files {
	files := Files{}

	for _, fileName := range fileNames {
		path := fmt.Sprintf("./public/%s.html", fileName)
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
			Data: &byteData,
			Template: template,
		}
	}

	return files
}

// RetrieveFile retrieves a byte array from memory
func (files Files) RetrieveFile(name string) *File {
	return files[name]
}
