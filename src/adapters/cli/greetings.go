package cli_adapter

import (
	"lendme/golang-seeder-backend/src/core/cases"
)

func ForHelloCase(cases *cases.Greetings) CliAdapter {
	return func(Args []string) string {
		// ################## CLI Layer ##################
		// This is where we do I/O layer-specific things, and adapt I/O input to business Case input.
		if len(Args) != 1 {
			return "Please input your name"
		}

		// ################## CLI Layer ##################

		// ############# Business Case Layer ##############
		result := cases.Hello()
		// ############# Business Case Layer ##############

		// ################## CLI Layer ##################
		// This is where we do I/O layer-specific things, and adapt business Case output to I/O output.
		return result + Args[0]
		// ################## CLI Layer ##################
	}
}