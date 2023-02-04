package controllers

import (
	"strconv"

	"github.com/MunBrian/todo-fiber-postgres/initializers"
	"github.com/MunBrian/todo-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
)

// create a slice of Todo struct
var todos = make([]models.Todo, 0)

// get all tasks
func GetAllTasks(c *fiber.Ctx) error {
	initializers.DB.Find(&todos)
	return c.Status(fiber.StatusOK).JSON(todos)
}

// get task by id
func GetTaskById(c *fiber.Ctx) error {

	//create struct todo of type Todo from model
	var task models.Todo

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
	initializers.DB.First(&task, paramIDUint)

	if task.ID == 0 {
		return c.Status(fiber.StatusOK).SendString("Task not found.")
	}

	//return task
	return c.Status(fiber.StatusOK).JSON(task)
}

// create task
func CreateTask(c *fiber.Ctx) error {

	//create struct todo of type Todo from model
	var todo models.Todo

	//get data from body
	if err := c.BodyParser(&todo); err != nil {
		return c.JSON(err.Error())
	}

	//Add todo data in the database
	initializers.DB.Create(&todo)

	return c.Status(fiber.StatusOK).SendString("Task created successfully.")
}

// update task
func UpdateTask(c *fiber.Ctx) error {
	var task models.Todo

	var upDatedTask models.Todo

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
	initializers.DB.First(&task, paramIDUint)

	//check if task exist
	if task.ID == 0 {
		//send message to user
		return c.Status(fiber.StatusOK).SendString("Task not found.")
	}

	//update task name
	task.Name = upDatedTask.Name

	//save updated value in initializers.DB
	initializers.DB.Save(&task)

	return c.Status(fiber.StatusOK).SendString("Task Updated successfully.")
}

// delete task
func DeleteTask(c *fiber.Ctx) error {
	var task models.Todo

	//get id from params
	id := c.Params("id")

	//convert id from param to int
	paramID, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	//convert paramId to uint
	paramIDUint := uint(paramID)

	initializers.DB.First(&task, paramIDUint)

	//check if task exist
	if task.ID == 0 {
		//send message to user
		return c.Status(fiber.StatusOK).SendString("Task not found.")
	} else {
		//delete task from database
		initializers.DB.Delete(&models.Todo{}, paramIDUint)
	}

	//return c.Status(fiber.StatusOK).SendString("Task deleted successfully.")
	return c.Status(fiber.StatusOK).JSON("Task deleted succesfully")
}
