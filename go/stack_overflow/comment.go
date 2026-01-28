package stackoverflow

type CommentableType string

const (
	QUESTION_TYPE_COMMENT CommentableType = "question"
	ANSWER_TYPE_COMMENT   CommentableType = "answer"
)

type Commentable interface {
	AddComment(comment *Comment) error
	GetComments() []*Comment
}

type Comment struct {
	id         int64
	content    string
	author     *User
	parentType CommentableType
	parent     Commentable
}

var autoIncrementCommentId int64 = 0

func NewComment(content string, user *User, t CommentableType, p Commentable) *Comment {
	autoIncrementCommentId += 1

	return &Comment{
		id:         autoIncrementCommentId,
		content:    content,
		author:     user,
		parentType: t,
		parent:     p,
	}
}
