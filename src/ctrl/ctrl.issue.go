package ctrl

import (
	"ledger-api/src/orm"

	"strconv"

	"github.com/gin-gonic/gin"
)

// Issue is controller
type Issue struct{}

// GetAllIssues is action
func (*Issue) GetAllIssues(c *gin.Context) {
	var issue orm.Issue
	issues, err := issue.GetAllIssues()
	if err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &issues)
}

// GetIssue is action
func (*Issue) GetIssue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		renderError(c, err)
		return
	}

	var ormObj orm.Issue
	issue, err := ormObj.GetIssue(id)
	if err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &issue)
}

// AddIssue is action
func (*Issue) AddIssue(c *gin.Context) {
	var ormObj orm.Issue
	if err := c.ShouldBind(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	var issue orm.Issue

	if count, err := issue.GetIssuesCount(); err != nil {
		renderError(c, err)
	} else {
		ormObj.ID = count + 1
		if err := issue.AddIssue(&ormObj); err != nil {
			renderError(c, err)
			return
		}
		renderJSON(c, &ormObj)
	}
}

// UpdateIssue is action
func (*Issue) UpdateIssue(c *gin.Context) {
	var ormObj orm.Issue
	if err := c.ShouldBind(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	var issue orm.Issue
	if err := issue.UpdateIssue(&ormObj); err != nil {
		renderError(c, err)
		return
	}

	renderJSON(c, &ormObj)
}

// DeleteIssue is action
func (*Issue) DeleteIssue(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		renderError(c, err)
		return
	}

	var issue orm.Issue
	if err := issue.DeleteIssue(id); err != nil {
		renderError(c, err)
		return
	}

	renderSuccessMessage(c, "done")
}

// GetStatusReport is action
func (*Issue) GetStatusReport(c *gin.Context) {
	var issue orm.Issue
	startedCount, err := issue.GetIssueCountByStatus(0)
	if err != nil {
		renderError(c, err)
	}

	inProgressCount, err := issue.GetIssueCountByStatus(1)
	if err != nil {
		renderError(c, err)
	}

	doneCount, err := issue.GetIssueCountByStatus(2)
	if err != nil {
		renderError(c, err)
	}

	var retStatus = []int{startedCount, inProgressCount, doneCount}
	renderJSON(c, &retStatus)
}

// GetNewReport is action
func (*Issue) GetNewReport(c *gin.Context) {
	var issue orm.Issue
	newCount, err := issue.GetIssueCountByIsNew(true)
	if err != nil {
		renderError(c, err)
	}

	oldCount, err := issue.GetIssueCountByIsNew(false)
	if err != nil {
		renderError(c, err)
	}

	var retStatus = []int{newCount, oldCount}
	renderJSON(c, &retStatus)
}
