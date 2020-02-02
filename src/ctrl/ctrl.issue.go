package ctrl

import (
	"ledger-api/src/orm"

	"github.com/gin-gonic/gin"
)

// Issue is controller
type Issue struct{}

// GetAllIssues is action
func (*Issue) GetAllIssues(c *gin.Context) {
	var issues *[]orm.Issue
	var issue orm.Issue
	if err := issue.GetAllIssues(issues); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &issues)
}

// AddIssue is action
func (*Issue) AddIssue(c *gin.Context) {
	var ormObj orm.Issue
	if err := c.ShouldBind(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	var issue orm.Issue
	if err := issue.AddIssue(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &ormObj)
}
