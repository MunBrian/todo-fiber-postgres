package main

import (
	//"fmt"
	"log"
	//"strconv"

	"github.com/MunBrian/todo-fiber-postgres/controllers"
	"github.com/MunBrian/todo-fiber-postgres/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()

}

func main() {

	app := fiber.New()

	//get all todos
	app.Get("/", controllers.GetAllTasks)

	//create todo
	app.Post("/", controllers.CreateTask)

	//get task by id
	app.Get("/task/:id", controllers.GetTaskById)

	//update task
	app.Patch("/task/:id", controllers.UpdateTask)

	//delete task
	app.Delete("/task/:id", controllers.DeleteTask)

	//listen to port
	log.Fatal(app.Listen(":3000"))

}
