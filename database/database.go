package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
    gorm.Model
    Name   string `json:"name"`
    Status string `json:"status"`
}

const (
    host     = "localhost"
    port     = 5432
    user     = "rejakucing"
    password = "rejakucing"
    dbname   = "todolistapirakamin"
)

var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
    host, port, user, password, dbname)

func InitDB() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		
	if err != nil {
		return err
		
	}
	
	db.AutoMigrate(&Task{})
	
	return nil
}

// CreateTask: function creates a new task in the database
func CreateTask(name string, status string) (Task, error) {
    var newTask = Task{Name: name, Status: status}

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return newTask, err
    }
    db.Create(&Task{Name: name, Status: status})

    return newTask, nil
}

// GetAllTasks: function fetches all the tasks from the database
func GetallTasks() ([]Task, error) {
    var tasks []Task

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return tasks, err
    }

    db.Find(&tasks)

    return tasks, nil
}

// GetTask: function fetches a single task from the database
func GetTask(id string) (Task, error) {
    var task Task

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return task, err
    }

    db.Where("ID = ?", id).First(&task)
    return task, nil
}

// DeleteTask: function removes an existing task from the database
func DeleteTask(id string) error {
    var task Task

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return err
    }

    db.Where("ID = ?", id).Delete(&task)
    return nil

}

// UpdateTask: function updates an existing task in the database
func UpdateTask(name string, status string, id string) (Task, error) {
    var newTask = Task{Name: name, Status: status}

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return newTask, err
    }

    db.Where("ID = ?", id).Updates(&Task{Name: newTask.Name, Status: newTask.Status})
    return newTask, nil
}

