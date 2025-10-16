package main

import (
    "fmt"
    "todo-app-go/internal/tasks"
)

func main() {
    fmt.Println("Welcome to To-Do App (Go Edition)!")
    tasks.AddTask("Learn Git")
    tasks.AddTask("Practice GoLang")
    tasks.ListTasks()
}
