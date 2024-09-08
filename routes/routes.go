package routes

import (
	"todolist-api/database"

	"github.com/gofiber/fiber/v2"
)

// GetTask: function fetches a single task from the database
func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")

    if id == "" {

        return c.Status(500).JSON(&fiber.Map{

            "message": "id cannot be empty",

        })
    }

    result, err := database.GetTask(id)
    if err != nil {
        return c.Status(500).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
    }

    return c.Status(200).JSON(&fiber.Map{
        "data":    result,
        "success": true,
        "message": "",
    })
}

// GetAllTasks: function fetches all the tasks from the database
func GetAllTasks(c *fiber.Ctx) error {
	result, err := database.GetallTasks()
    if err != nil {
        return c.Status(500).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
    }

    return c.Status(200).JSON(&fiber.Map{
        "data":    result,
        "success": true,
        "message": "All Tasks",
    })
}

// AddTask: function creates a new task in the database.
func AddTask(c *fiber.Ctx) error {
    newTask := new(database.Task)

    err := c.BodyParser(newTask)
    if err != nil {
        c.Status(400).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
        return err
    }

    result, err := database.CreateTask(newTask.Name, newTask.Status)
    if err != nil {
        c.Status(400).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
        return err
    }

    c.Status(200).JSON(&fiber.Map{
        "data":    result,
        "success": true,
        "message": "Task added!",
    })
    return nil
}


// DeleteTask: function removes an existing task from the database
func DeleteTask(c *fiber.Ctx) error {
    id := c.Params("id")

    if id == "" {

        return c.Status(500).JSON(&fiber.Map{

            "message": "id cannot be empty",
        })
    }

    err := database.DeleteTask(id)
    if err != nil {
        return c.Status(500).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
    }

    return c.Status(200).JSON(&fiber.Map{
        "data":    nil,
        "success": true,
        "message": "Task Deleted Successfully",
    })
}

// UpdateTask: function updates an existing task in the database
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

    if id == "" {

        return c.Status(500).JSON(&fiber.Map{

            "message": "id cannot be empty",
        })
    }

    newTask := new(database.Task)

    err := c.BodyParser(newTask)
    if err != nil {
        c.Status(400).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
        return err
    }

    result, err := database.UpdateTask(newTask.Name, newTask.Status, id)

    if err != nil {
        c.Status(400).JSON(&fiber.Map{
            "data":    nil,
            "success": false,
            "message": err,
        })
        return err
    }

    c.Status(200).JSON(&fiber.Map{
        "data":    result,
        "success": true,
        "message": "Task Updated!",
    })
    return nil
}
