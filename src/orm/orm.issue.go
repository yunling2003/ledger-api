package orm

import (
	"time"
)

// Issue Object
type Issue struct {
	ID          uint64    `bson:"id"`
	Name        string    `bson:"name"`
	Status      int       `bson:"status"`
	IsNew       bool      `bson:"isNew"`
	CreatedDate time.Time `bson:"createdDate"`
	CreatedBy   string    `bson:"createdBy"`
}

// GetAllIssues will get all issues
func (*Issue) GetAllIssues(issues *[]Issue) error {
	return db.C("issue").Find(nil).All(&issues)
}

// AddIssue insert new issue
func (*Issue) AddIssue(obj *Issue) error {
	return db.C("issue").Insert(obj)
}
