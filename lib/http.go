package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Alias type so that we can add new functions.
type Response http.Response

// Function to easily fetch response body as string (So that we dont have to do all this juggling in each test case).
func (response *Response) ReadBody() string {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyBytes)
}
