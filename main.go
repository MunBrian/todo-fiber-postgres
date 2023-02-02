package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	fmt.Print(user)

	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=5432", user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to the database.")
	}

	fmt.Println("Successfully connected!")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		todos := new(Todo)

		db.Find(&todos)

		return c.Status(fiber.StatusOK).JSON(todos)

	})

	log.Fatal(app.Listen(":3000"))

}
