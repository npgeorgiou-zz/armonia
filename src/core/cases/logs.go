package cases

import (
	"lendme/golang-seeder-backend/src/core/entities"
	"lendme/golang-seeder-backend/src/di"
)

type Logs struct {
	Di *di.Di
}

func (logs *Logs) CreateLog(log *entities.Log) {
	// ############# Business Case Layer ##############
	// Do business logic stuff like search if exists, or other stuff.
	// ############# Business Case Layer ##############

	// ############# Delegating dirty jobs ############
	// Business case layer knows only the business rules. Only high-level stuff. Delegates implementation details
	// to services injected via Di. This way it is immune to changes. In other words:
	// Core Logic of the system changes only when business requirements change, not when developers play with different tools.
	logs.Di.GetPersistence().SaveLog(log)
	// ############# Delegating dirty jobs ############

	// ############# Business Case Layer ##############
	// Do business logic stuff, like send an email, or other stuff. Again by using the injected services ;)
	// ############# Business Case Layer ##############
}

func (logs *Logs) DeleteLog(id int) {
	// ############# Business Case Layer ##############
	// Do business logic stuff like search if exists, or other stuff.
	// ############# Business Case Layer ##############

	// ############# Delegating dirty jobs ############
	// Business case layer knows only the business rules. Only high-level stuff. Delegates implementation details
	// to services injected via Di. This way it is immune to changes. In other words:
	// Core Logic of the system changes only when business requirements change, not when developers play with different tools.
	logs.Di.GetPersistence().DeleteLog(id)
	// ############# Delegating dirty jobs ############

	// ############# Business Case Layer ##############
	// Do business logic stuff, like send an email, or other stuff. Again by using the injected services ;)
	// ############# Business Case Layer ##############
}