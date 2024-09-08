package main

import (
	"todolist-api/database"
	"todolist-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New() // Creating a new instance of Fiber.

	dbErr := database.InitDB() // Database connection initialization

    if dbErr != nil {
        panic(dbErr)
    }


    list := app.Group("/tasks")

    list.Get("/", routes.GetAllTasks) //Get endpoint for fetching all the tasks.

    list.Get("/:id", routes.GetTask) //Get endpoint for fetching a single task.

    list.Post("/", routes.AddTask) //Post endpoint for add a new task.

    list.Delete("/:id", routes.DeleteTask) //Delete endpoint for removing an existing task.

    list.Put("/:id", routes.UpdateTask) //Patch endpoint for updating an existing task.

	app.Listen(":8081")
}