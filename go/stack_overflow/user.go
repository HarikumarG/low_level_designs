package stackoverflow

type User struct {
	id              int64
	name            string
	email           string
	reputationScore int64
}

var autoIncrementUserId int64 = 0

func NewUser(name, email string) *User {
	autoIncrementUserId += 1
	return &User{
		id:              autoIncrementUserId,
		name:            name,
		email:           email,
		reputationScore: 0,
	}
}

func (u *User) UpdateReputation(value int) error {
	u.reputationScore += int64(value)
	return nil
}
