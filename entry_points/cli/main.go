package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	util "lendme/golang-seeder-backend/entry_points"
	"lendme/golang-seeder-backend/entry_points/cli/commands"
	"os"
)

func main() {
	commandName := os.Args[1]

	// A cli command can be executed directly here...
	if commandName == "create-db" {
		persistence := util.CreatePersistence()
		persistence.CreateDB()

		println("Database logs created.")
		return
	}

	// ...Or be delegated to a dedicated Command...
	args := os.Args[2:]
	di := util.CreateDefaultDi()

	switch commandName {
	case "create-schema":
		println(commands.CreateSchema(di, args))
		break
	case "create-log":
		println(commands.CreateLog(di, args))
		break
	case "hello":
		println(commands.Hello(di, args))
		break
	default:
		println("Unknown command " + commandName)
	}
}
