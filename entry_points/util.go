package util

import (
	"github.com/go-redis/redis/v7"
	http_adapter "lendme/golang-seeder-backend/src/adapters/http"
	"lendme/golang-seeder-backend/src/core/cases"
	"lendme/golang-seeder-backend/src/di"
	"lendme/golang-seeder-backend/src/di/services"
	"net/http"
	"os"
)

func CreateDefaultDi() di.Di {
	persistence := CreatePersistence()
	keyValue := CreateKeyValue()

	di := di.NewDi(
		persistence,
		keyValue,
	)

	return di
}

func CreatePersistence() services.Persistence {
	persistence := services.GormPersistence{}
	persistence.OpenConnection(os.Getenv("DB_DRIVER"), os.Getenv("DB_URI")+os.Getenv("DB_NAME"))
	return &persistence
}

func CreateKeyValue() services.KeyValue {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		// No password set
		Password: "",
		// Use default DB
		DB:       0,
	})

	service := services.NewRedisKeyValue(client)
	return &service
}

func CreateRouter(di di.Di) *http.ServeMux {
	router := &http.ServeMux{}

	// Inject Cases with Di.
	greetingsCases := &cases.Greetings{Di: &di}
	logsCases := &cases.Logs{Di: &di}

	// Register routes.
	http_adapter.RegisterGreetingsRoutes(router, greetingsCases)
	http_adapter.RegisterLogRoutes(router, logsCases)

	return router
}
