package taskmanagement

import "time"

type TaskPriority int

const (
	P0 TaskPriority = 0
	P1 TaskPriority = 1
	P2 TaskPriority = 2
	P3 TaskPriority = 3
)

type TaskStatus string

const (
	PENDING     TaskStatus = "pending"
	IN_PROGRESS TaskStatus = "inprogress"
	COMPLETED   TaskStatus = "completed"
)

type Task struct {
	id           int64
	title        string
	description  string
	dueDate      time.Time
	priority     TaskPriority
	status       TaskStatus
	assignedUser *User
}

var autoIncrementTaskId int64 = 0

func NewTask(title string, description string, dueDate time.Time, priority TaskPriority, user *User) *Task {
	autoIncrementTaskId++
	return &Task{
		id:           autoIncrementTaskId,
		title:        title,
		description:  description,
		dueDate:      dueDate,
		priority:     priority,
		status:       IN_PROGRESS,
		assignedUser: user,
	}
}

func (t *Task) GetId() int64 {
	return t.id
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetDueDate() time.Time {
	return t.dueDate
}

func (t *Task) GetPriority() TaskPriority {
	return t.priority
}

func (t *Task) GetStatus() TaskStatus {
	return t.status
}

func (t *Task) GetAssignedUser() *User {
	return t.assignedUser
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) SetDescription(desc string) {
	t.description = desc
}

func (t *Task) SetDueDate(date time.Time) {
	t.dueDate = date
}

func (t *Task) SetPriority(p TaskPriority) {
	t.priority = p
}

func (t *Task) SetStatus(status TaskStatus) {
	t.status = status
}

func (t *Task) SetAssignedUser(user *User) {
	t.assignedUser = user
}
