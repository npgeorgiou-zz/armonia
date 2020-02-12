package commands

import (
	cli_adapter "lendme/golang-seeder-backend/src/adapters/cli"
	cases2 "lendme/golang-seeder-backend/src/core/cases"
	"lendme/golang-seeder-backend/src/di"
)

var Hello DiAwareCommand = func(di di.Di, Args []string) string {
	cases := &cases2.Greetings{Di: &di}
	return cli_adapter.ForHelloCase(cases)(Args)
}
