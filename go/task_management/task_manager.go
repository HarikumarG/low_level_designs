package taskmanagement

import (
	"errors"
	"sync"
	"time"
)

type TaskManager struct {
	mu       sync.Mutex
	users    map[int64]*User
	tasks    map[int64]*Task
	userTask map[int64][]int64
}

var instance *TaskManager
var once sync.Once

func NewTaskManager() *TaskManager {
	once.Do(func() {
		instance = &TaskManager{
			users:    make(map[int64]*User),
			tasks:    make(map[int64]*Task),
			userTask: make(map[int64][]int64),
		}
	})
	return instance
}

func (taskManager *TaskManager) CreateUser(email, name string) int64 {
	taskManager.mu.Lock()
	defer taskManager.mu.Unlock()

	var user *User = NewUser(email, name)
	taskManager.users[user.GetId()] = user
	return user.GetId()
}

func (taskManager *TaskManager) CreateTask(title string, desc string, dueDate time.Time, priority TaskPriority, toBeAssignedUserId int64) int64 {
	taskManager.mu.Lock()
	defer taskManager.mu.Unlock()

	var toBeAssignedUser *User = taskManager.users[toBeAssignedUserId]
	var task *Task = NewTask(title, desc, dueDate, priority, nil)
	taskManager.tasks[task.GetId()] = task
	taskManager.assignTaskToUser(task, toBeAssignedUser)
	return task.GetId()
}

func (taskManager *TaskManager) UpdateTask(taskId int64, title string, desc string, dueDate *time.Time, priority *TaskPriority, status *TaskStatus, toBeAssignedUserId int64) error {
	taskManager.mu.Lock()
	defer taskManager.mu.Unlock()

	var task *Task = taskManager.tasks[taskId]
	if task == nil {
		return errors.New("Invalid Task Id")
	}
	var user *User = taskManager.users[toBeAssignedUserId]
	if toBeAssignedUserId > 0 && user == nil {
		return errors.New("Invalid User Id")
	}

	if title != "" && title != task.GetTitle() {
		task.SetTitle(title)
	}

	if desc != "" && desc != task.GetDescription() {
		task.SetDescription(desc)
	}

	if dueDate != nil && *dueDate != task.GetDueDate() {
		task.SetDueDate(*dueDate)
	}

	if priority != nil && *priority != task.GetPriority() {
		task.SetPriority(*priority)
	}

	if status != nil && *status != task.GetStatus() {
		task.SetStatus(*status)
	}

	if user != task.GetAssignedUser() {
		taskManager.unassignTaskToUser(task)
		taskManager.assignTaskToUser(task, user)
	}

	return nil
}

func (taskManager *TaskManager) DeleteTask(taskId int64) error {
	taskManager.mu.Lock()
	defer taskManager.mu.Unlock()

	var task *Task = taskManager.tasks[taskId]
	if task == nil {
		return errors.New("Invalid Task Id")
	}
	taskManager.unassignTaskToUser(task)
	delete(taskManager.tasks, taskId)
	return nil
}

func (taskManager *TaskManager) AssignTaskToUser(taskId int64, userId int64) error {
	taskManager.mu.Lock()
	defer taskManager.mu.Unlock()

	var task *Task = taskManager.tasks[taskId]
	if task == nil {
		return errors.New("Invalid Task Id")
	}
	var user *User = taskManager.users[userId]
	if userId > 0 && user == nil {
		return errors.New("Invalid User Id")
	}

	taskManager.unassignTaskToUser(task)
	taskManager.assignTaskToUser(task, user)
	return nil
}

func (taskManager *TaskManager) MarkTaskToCompleted(taskId int64) error {
	taskManager.mu.Lock()
	defer taskManager.mu.Unlock()

	var task *Task = taskManager.tasks[taskId]
	if task == nil {
		return errors.New("Invalid Task Id")
	}

	task.SetStatus(COMPLETED)
	return nil
}

func (taskManager *TaskManager) GetTaskHistory(userId int64) ([]*Task, error) {
	var user *User = taskManager.users[userId]
	if user == nil {
		return make([]*Task, 0), errors.New("Invalid User Id")
	}
	var tasks []*Task
	taskIds := taskManager.userTask[user.GetId()]
	for _, taskId := range taskIds {
		tasks = append(tasks, taskManager.tasks[taskId])
	}
	return tasks, nil
}

func (taskManager *TaskManager) SetReminder(taskId int64, date *time.Time) error {
	if date != nil && date.Before(time.Now()) {
		return errors.New("Invalid date")
	}
	var task *Task = taskManager.tasks[taskId]
	if task == nil {
		return errors.New("Invalid Task Id")
	}
	task.SetReminder(date)
	return nil
}

func (taskManager *TaskManager) unassignTaskToUser(task *Task) {
	var previousAssignedUser *User = task.GetAssignedUser()
	if previousAssignedUser != nil {
		tasks := taskManager.userTask[previousAssignedUser.GetId()]
		for i, id := range tasks {
			if id == task.GetId() {
				taskManager.userTask[previousAssignedUser.GetId()] = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
	}
}

func (taskManager *TaskManager) assignTaskToUser(task *Task, user *User) {
	task.SetAssignedUser(user)
	if user != nil {
		taskManager.userTask[user.GetId()] = append(taskManager.userTask[user.GetId()], task.GetId())
	}
}
