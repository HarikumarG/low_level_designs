package stackoverflow

type Question struct {
	id          int64
	title       string
	description string
	author      *User
}

var autoIncrementQuestionId int64 = 0

func NewQuestion(title, desc string, user *User) *Question {
	autoIncrementQuestionId += 1
	return &Question{
		id:          autoIncrementQuestionId,
		title:       title,
		description: desc,
		author:      user,
	}
}
