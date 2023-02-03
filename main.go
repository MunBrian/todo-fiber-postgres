package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	//"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// create a todo model
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

// declare global var
var db *gorm.DB
var err error

var todos = make([]Todo, 0)

func main() {
	//get env variables
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=%v", user, password, dbname, port)

	//connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//handle err
	if err != nil {
		panic("failed to connect to the database.")
	} else {
		fmt.Println("Successfully connected!")
	}

	//make migrations to database if not already created
	db.AutoMigrate(&Todo{})

	app := fiber.New()

	//get all todos
	app.Get("/", func(c *fiber.Ctx) error {

		db.Find(&todos)

		return c.Status(fiber.StatusOK).JSON(todos)

	})

	//create todo
	app.Post("/", func(c *fiber.Ctx) error {

		var todo Todo

		//get data from body
		if err := c.BodyParser(&todo); err != nil {
			return c.JSON(err.Error())
		}

		//Add todo data in the database
		db.Create(&todo)

		return c.Status(fiber.StatusOK).SendString("Task created successfully.")
	})

	//get task by id
	app.Get("/task/:id", func(c *fiber.Ctx) error {
		var task Todo

		//get id from params
		id := c.Params("id")

		//convert id from param to int
		paramID, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}

		//convert paramId to uint
		paramIDUint := uint(paramID)

		//get task from database
		db.First(&task, paramIDUint)

		if task.ID == 0 {
			return c.Status(fiber.StatusOK).SendString("Task not found.")
		}

		//return task
		return c.Status(fiber.StatusOK).JSON(task)
	})

	//update task
	app.Patch("/task/:id", func(c *fiber.Ctx) error {
		var task Todo

		var upDatedTask Todo

		//get id from params
		id := c.Params("id")

		//convert id from param to int
		paramID, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}

		//convert paramId to uint
		paramIDUint := uint(paramID)

		//get updated value from body and point to upDatedTask
		if err := c.BodyParser(&upDatedTask); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		//Fetch task from database with id same as param's id
		db.First(&task, paramIDUint)

		//check if task exist
		if task.ID == 0 {
			//send message to user
			return c.Status(fiber.StatusOK).SendString("Task not found.")
		}

		//update task name
		task.Name = upDatedTask.Name

		//save updated value in db
		db.Save(&task)

		return c.Status(fiber.StatusOK).SendString("Task Updated successfully.")
	})

	app.Delete("/task/:id", func(c *fiber.Ctx) error {

		var task Todo

		//get id from params
		id := c.Params("id")

		//convert id from param to int
		paramID, err := strconv.Atoi(id)

		if err != nil {
			panic(err)
		}

		//convert paramId to uint
		paramIDUint := uint(paramID)

		db.First(&task, paramIDUint)

		//check if task exist
		if task.ID == 0 {
			//send message to user
			return c.Status(fiber.StatusOK).SendString("Task not found.")
		} else {
			//delete task from database
			db.Delete(&Todo{}, paramIDUint)
		}

		//return c.Status(fiber.StatusOK).SendString("Task deleted successfully.")
		return c.Status(fiber.StatusOK).JSON("Task deleted succesfully")
	})

	//listen to port
	log.Fatal(app.Listen(":3000"))

}
