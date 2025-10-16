package tasks

import "fmt"

var taskList []string

func AddTask(task string) {
    taskList = append(taskList, task)
    fmt.Println("Task added:", task)
}

func ListTasks() {
    fmt.Println("Your Tasks:")
    for _, task := range taskList {
        fmt.Println("-", task)
    }
}
