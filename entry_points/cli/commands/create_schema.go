package commands

import (
	"lendme/golang-seeder-backend/src/di"
)

var CreateSchema DiAwareCommand = func (di di.Di, Args []string) string {
	di.GetPersistence().DeleteSchema()
	di.GetPersistence().MigrateSchema()

	return "schema created"
}
