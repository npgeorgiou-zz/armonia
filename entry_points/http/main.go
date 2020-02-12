package main

import (
	util "lendme/golang-seeder-backend/entry_points"
	"log"
	"net/http"
)

func main() {
	di := util.CreateDefaultDi()
	defer di.GetPersistence().CloseConnection()

	router := util.CreateRouter(di)
	log.Fatal(http.ListenAndServe(":80", router))
}
