package stackoverflow

type VoteType string

const (
	QUESTION_TYPE_VOTE VoteType = "question"
	ANSWER_TYPE_VOTE   VoteType = "answer"
)

type Votable interface {
	Vote(vote *Vote)
	GetVotes() []*Vote
}

type Vote struct {
	id         int64
	author     *User
	parentType VoteType
	parent     Votable
	value      int
}

var autoIncrementVoteId int64 = 0

func NewVote(user *User, t VoteType, p Votable, v int) *Vote {
	autoIncrementVoteId += 1
	return &Vote{
		id:         autoIncrementVoteId,
		author:     user,
		parentType: t,
		parent:     p,
		value:      v,
	}
}
