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
	Id           int64        `json:"id"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	DueDate      time.Time    `json:"dueDate"`
	Priority     TaskPriority `json:"priority"`
	Status       TaskStatus   `json:"status"`
	AssignedUser *User        `json:"assignedUser"`
	Reminder     *time.Time   `json:"reminder"`
}

var autoIncrementTaskId int64 = 0

func NewTask(title string, description string, dueDate time.Time, priority TaskPriority, user *User) *Task {
	autoIncrementTaskId++
	return &Task{
		Id:           autoIncrementTaskId,
		Title:        title,
		Description:  description,
		DueDate:      dueDate,
		Priority:     priority,
		Status:       PENDING,
		AssignedUser: user,
	}
}

func (t *Task) GetId() int64 {
	return t.Id
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetDueDate() time.Time {
	return t.DueDate
}

func (t *Task) GetPriority() TaskPriority {
	return t.Priority
}

func (t *Task) GetStatus() TaskStatus {
	return t.Status
}

func (t *Task) GetAssignedUser() *User {
	return t.AssignedUser
}

func (t *Task) SetTitle(title string) {
	t.Title = title
}

func (t *Task) SetDescription(desc string) {
	t.Description = desc
}

func (t *Task) SetDueDate(date time.Time) {
	t.DueDate = date
}

func (t *Task) SetPriority(p TaskPriority) {
	t.Priority = p
}

func (t *Task) SetStatus(status TaskStatus) {
	t.Status = status
}

func (t *Task) SetAssignedUser(user *User) {
	t.AssignedUser = user
}

func (t *Task) SetReminder(date *time.Time) {
	t.Reminder = date
}
