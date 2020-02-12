package di

import (
	"lendme/golang-seeder-backend/src/di/services"
)

type Di struct {
	persistence services.Persistence
	keyValue    services.KeyValue
}

// "Constructor" pattern to allow instantiation of struct and keep struct members private,
// so they cant be set to something else.
func NewDi(persistence services.Persistence, keyValue services.KeyValue) Di {
	di := Di{
		persistence: persistence,
		keyValue:    keyValue,
	}
	return di
}

func (di *Di) GetPersistence() services.Persistence {
	return di.persistence
}

func (di *Di) GetKeyValue() services.KeyValue {
	return di.keyValue
}

