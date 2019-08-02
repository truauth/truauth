package settings

import (
	"encoding/json"
	"fmt"
	"os"
)

// Init initializes the provided settings model
func Init(model interface{}, name string) {
	path := fmt.Sprintf("../../%s.json", name)

	file, _ := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)

	err := decoder.Decode(model)

	if err != nil {
		panic(err.Error())
	}
}
