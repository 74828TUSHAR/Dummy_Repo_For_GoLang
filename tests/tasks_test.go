package tests

import (
    "testing"
    "todo-app-go/internal/tasks"
)

func TestAddTask(t *testing.T) {
    tasks.AddTask("Test Git with Go")
}
