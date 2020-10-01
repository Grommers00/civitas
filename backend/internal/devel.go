package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// NotImplementedHandler writes to ResponseWriter lack of implementation for context
func NotImplementedHandler(context string, w http.ResponseWriter) {
	w.Write([]byte(context + " Not Implemented!\n"))
}

// UnwrapJSONData parses a jsonfile from declared src and unmarshalls into object
func UnwrapJSONData(src string, object interface{}) error {
	data, err := os.Open(src)

	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(data)
	defer data.Close()

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(byteValue), &object)

	if err != nil {
		return err
	}

	return nil
}
