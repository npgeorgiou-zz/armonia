package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"lendme/golang-seeder-backend/src/core/entities"
	"os"
)

// Error(s) thrown by GormPersistence.
type NotFoundError struct {
}
func (error NotFoundError) Error() string {
	return "Entity not found"
}

type Persistence interface {
	SaveLog(entity *entities.Log)
	FindLog(id int) (entities.Log, error)
	FindLogBy(property string, value interface{}) (entities.Log, error)
	FindFirstLog() (entities.Log, error)
	DeleteLog(id int)
	OpenConnection(driver string, connectionString string)
	CloseConnection()
	DeleteSchema()
	MigrateSchema()
	CreateDB()
}

type GormPersistence struct {
	db *gorm.DB
}

func (gp *GormPersistence) SaveLog(entity *entities.Log) {
	gp.db.Create(entity)
}

func (gp *GormPersistence) FindLog(id int) (entities.Log, error) {
	return gp.FindLogBy("id", id)
}

func (gp *GormPersistence) FindLogBy(property string, value interface{}) (entities.Log, error) {
	toFill := entities.Log{}
	gp.db.Find(&toFill, property + " = ?", value)

	if toFill.ID == 0 {
		return toFill, new(NotFoundError)
	}

	return toFill, nil
}

func (gp *GormPersistence) FindFirstLog() (entities.Log, error) {
	toFill := entities.Log{}
	gp.db.First(&toFill)

	if toFill.ID == 0 {
		return toFill, new(NotFoundError)
	}

	return toFill, nil
}

func (gp *GormPersistence) DeleteLog(id int) {
	entity, err := gp.FindLog(id)
	if err != nil {
		panic(err)
	}

	gp.db.Delete(&entity)
}

func (gp *GormPersistence) OpenConnection(driver string, connectionString string) {
	db, err := gorm.Open(driver, connectionString)
	if err != nil {
		fmt.Printf("%v", err)
		panic("failed to connect to database")
	}

	gp.db = db
}

func (gp *GormPersistence) CloseConnection() {
	gp.db.Close()
}

func (gp *GormPersistence) DeleteSchema() {
	gp.db.DropTableIfExists(&entities.Log{})
}

func (gp *GormPersistence) MigrateSchema() {
	gp.db.AutoMigrate(&entities.Log{})
}

func (gp *GormPersistence) CreateDB() {
	gp.OpenConnection(os.Getenv("DB_DRIVER"), os.Getenv("DB_URI"))
	gp.db.Exec("CREATE DATABASE IF NOT EXISTS logs")
}
