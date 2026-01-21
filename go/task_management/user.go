package taskmanagement

type User struct {
	id    int64
	email string
	name  string
}

var autoIncrementUserId int64 = 0

func NewUser(email, name string) *User {
	autoIncrementUserId += 1
	return &User{
		id:    autoIncrementUserId,
		email: email,
		name:  name,
	}
}

func (u *User) GetId() int64 {
	return u.id
}
