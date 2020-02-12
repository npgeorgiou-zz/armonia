package cli_adapter

import (
	"fmt"
	"lendme/golang-seeder-backend/src/core/cases"
	"lendme/golang-seeder-backend/src/core/entities"
)

func ForCreateLogCase(cases *cases.Logs) CliAdapter {
	return func(Args []string) string {
		// ################## CLI Layer ##################
		// This is where we do I/O layer-specific things, and adapt I/O input to business Case input.
		if (len(Args) != 2) {
			return "Please input level and project"
		}

		log := entities.Log{
			Level: Args[0],
			Project: Args[1],
		}
		// ################## CLI Layer ##################

		// ############# Business Case Layer ##############
		cases.CreateLog(&log)
		// ############# Business Case Layer ##############

		// ################## CLI Layer ##################
		// ################## CLI Layer ##################
		// This is where we do I/O layer-specific things, and adapt business Case output to I/O output.
		return fmt.Sprintf("Log created, ID: %d", log.ID)
	}
}