package stackoverflow

type Answer struct {
	id         int64
	content    string
	author     *User
	question   *Question
	isAccepted bool
}

var autoIncrementAnswerId int64 = 0

func NewAnswer(content string, user *User, q *Question) *Answer {
	return &Answer{
		id:         autoIncrementAnswerId,
		content:    content,
		author:     user,
		question:   q,
		isAccepted: false,
	}
}

func (a *Answer) MarkAsAccepted() error {
	a.isAccepted = true
	return nil
}
