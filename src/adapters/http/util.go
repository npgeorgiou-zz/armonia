package http_adapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func readInput(request *http.Request, input interface{}) error {
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, input)
	if err != nil {
		return err
	}

	return nil
}
