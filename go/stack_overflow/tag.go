package stackoverflow

type Tag struct {
	id   int64
	name string
}

var autoIncrementTagId int64 = 0

func NewTag(name string) *Tag {
	autoIncrementTagId += 1
	return &Tag{
		id:   autoIncrementTagId,
		name: name,
	}
}
