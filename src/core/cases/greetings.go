package cases

import (
	"lendme/golang-seeder-backend/src/di"
)

type Greetings struct {
	Di *di.Di
}

func (greetings *Greetings) Hello() string {
	return "Hi there, I log stuff!"
}

func (greetings *Greetings) Bye() string {
	return "Bye!"
}

