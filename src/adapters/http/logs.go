package http_adapter

import (
	"fmt"
	"lendme/golang-seeder-backend/src/core/cases"
	"lendme/golang-seeder-backend/src/core/entities"
	"net/http"
	"time"
)

type CreateLogInput struct {
	Project string
	Level   string
	Message string
	SentAt int64
}

func ForCreateLogCase(cases *cases.Logs) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt I/O input to business Case input.
		input := CreateLogInput{}
		readInput(request, &input)

		log := entities.Log{
			Project: input.Project,
			Level: input.Level,
			Message: input.Message,
			SentAt: time.Unix(input.SentAt, 0),
		}

		// ################## Http Layer ##################

		// ############# Business Case Layer ##############
		cases.CreateLog(&log)
		// ############# Business Case Layer ##############

		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt business Case output to I/O output.
		fmt.Fprintf(writer, "OK")
		// ################## Http Layer ##################
	}
}

type DeleteLogInput struct {
	Id int
}
func ForDeleteLogCase(cases *cases.Logs) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt I/O input to business Case input.
		input := DeleteLogInput{}
		readInput(request, &input)

		// ################## Http Layer ##################

		// ############# Business Case Layer ##############
		cases.DeleteLog(input.Id)
		// ############# Business Case Layer ##############

		// ################## Http Layer ##################
		// This is where we do I/O layer-specific things, and adapt business Case output to I/O output.
		fmt.Fprintf(writer, "OK")
		// ################## Http Layer ##################
	}
}

func RegisterLogRoutes(router *http.ServeMux, cases *cases.Logs) {
	router.HandleFunc("/create", ForCreateLogCase(cases))
	router.HandleFunc("/delete", ForDeleteLogCase(cases))
}
