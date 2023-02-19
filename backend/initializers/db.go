package initializers

import (
	"fmt"
	"os"

	"github.com/MunBrian/todo-fiber-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to db
func ConnectToDB() {

	//get env variables
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v", user, password, dbname, port)

	//connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//handle err
	if err != nil {
		panic("failed to connect to the database.")
	} else {
		fmt.Println("Successfully connected!")
	}

}

// migrate database
func SyncDB() {
	//make migrations to database if not already created
	DB.AutoMigrate(&models.Todo{})
}
