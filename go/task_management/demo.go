package taskmanagement

import (
	"encoding/json"
	"fmt"
	"time"
)

func TaskManagementDemo() {
	tm := NewTaskManager()

	userId1 := tm.CreateUser("sample1@abc.com", "sample1")
	userId2 := tm.CreateUser("sample2@abc.com", "sample2")

	task1 := tm.CreateTask("Work1", "Desc1", time.Now(), P0, -1)
	task2 := tm.CreateTask("Work2", "Desc2", time.Now(), P1, -1)

	tm.AssignTaskToUser(task1, userId1)
	tm.AssignTaskToUser(task2, userId2)

	tasks, _ := tm.GetTaskHistory(userId1)
	printTasksJSON("User 1 Tasks:", tasks)
	tasks, _ = tm.GetTaskHistory(userId2)
	printTasksJSON("User 2 Tasks:", tasks)

	var p2 TaskPriority = P2
	var status TaskStatus = IN_PROGRESS
	tm.UpdateTask(task1, "Work11", "Desc11", nil, &p2, &status, userId2)
	tasks, _ = tm.GetTaskHistory(userId1)
	printTasksJSON("User 1 Tasks (After Update):", tasks)

	tm.DeleteTask(task2)
	tasks, _ = tm.GetTaskHistory(userId2)
	printTasksJSON("User 2 Tasks (After Delete):", tasks)

	tm.MarkTaskToCompleted(task1)
	tasks, _ = tm.GetTaskHistory(userId2)
	printTasksJSON("User 2 Tasks (After Completion):", tasks)
}

func printTasksJSON(label string, tasks []*Task) {
	fmt.Println("\n" + label)
	jsonBytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
