package commands

import (
	cli_adapter "lendme/golang-seeder-backend/src/adapters/cli"
	cases2 "lendme/golang-seeder-backend/src/core/cases"
	"lendme/golang-seeder-backend/src/di"
)

var CreateLog DiAwareCommand = func (di di.Di, Args []string) string {
	cases := &cases2.Logs{Di: &di}
	return cli_adapter.ForCreateLogCase(cases)(Args)
}
