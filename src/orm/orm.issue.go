package orm

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Issue Object
type Issue struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Status      int       `json:"status"`
	IsNew       bool      `json:"isnew"`
	CreatedDate time.Time `json:"createddate"`
	CreatedBy   string    `json:"createdby"`
}

// GetAllIssues will get all issues
func (*Issue) GetAllIssues() ([]Issue, error) {
	issues := make([]Issue, 0, 10)
	err := db.C("issue").Find(nil).All(&issues)
	return issues, err
}

// GetIssue get issue by id
func (*Issue) GetIssue(id uint64) (Issue, error) {
	issue := Issue{}
	err := db.C("issue").Find(bson.M{"id": id}).One(&issue)
	return issue, err
}

// GetIssuesCount get all issues count
func (*Issue) GetIssuesCount() (int, error) {
	return db.C("issue").Find(nil).Count()
}

// AddIssue insert new issue
func (*Issue) AddIssue(obj *Issue) error {
	return db.C("issue").Insert(obj)
}

// UpdateIssue update issue
func (*Issue) UpdateIssue(obj *Issue) error {
	return db.C("issue").Update(bson.M{"id": obj.ID}, obj)
}

// DeleteIssue delete issue
func (*Issue) DeleteIssue(id uint64) error {
	return db.C("issue").Remove(bson.M{"id": id})
}

// GetIssueCountByStatus return issue count by status
func (*Issue) GetIssueCountByStatus(status uint64) (int, error) {
	return db.C("issue").Find(bson.M{"status": status}).Count()
}

// GetIssueCountByIsNew return issue count by isNew
func (*Issue) GetIssueCountByIsNew(isNew bool) (int, error) {
	return db.C("issue").Find(bson.M{"isnew": isNew}).Count()
}
