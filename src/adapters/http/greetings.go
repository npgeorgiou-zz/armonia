package http_adapter

import (
	"fmt"
	"lendme/golang-seeder-backend/src/core/cases"
	"net/http"
)

func ForHelloCase(cases *cases.Greetings) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt I/O input to business Case input.
		// ################## Http Layer ##################

		// ############# Business Case Layer ##############
		result := cases.Hello()
		// ############# Business Case Layer ##############

		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt business Case output to I/O output.
		fmt.Fprintf(writer, result)
		// ################## Http Layer ##################
	}
}

func ForByeCase(cases *cases.Greetings) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt I/O input to business Case input.
		// ################## Http Layer ##################

		// ############# Business Case Layer ##############
		result := cases.Bye()
		// ############# Business Case Layer ##############

		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt business Case output to I/O output.
		fmt.Fprintf(writer, result)
		// ################## Http Layer ##################
	}
}

func RegisterGreetingsRoutes(router *http.ServeMux, cases *cases.Greetings) {
	router.HandleFunc("/", ForHelloCase(cases))
	router.HandleFunc("/bye", ForByeCase(cases))
}