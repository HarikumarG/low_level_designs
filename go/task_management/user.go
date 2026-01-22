package taskmanagement

type User struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

var autoIncrementUserId int64 = 0

func NewUser(email, name string) *User {
	autoIncrementUserId += 1
	return &User{
		Id:    autoIncrementUserId,
		Email: email,
		Name:  name,
	}
}

func (u *User) GetId() int64 {
	return u.Id
}
