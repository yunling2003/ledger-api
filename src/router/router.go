package router

import (
	"ledger-api/src/ctrl"
	"ledger-api/src/middleware"

	"github.com/gin-gonic/gin"
)

// Route contains all the api routes mappings
func Route() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Options)

	r.GET("/", ctrl.SayHello)

	var issue ctrl.Issue
	r.GET("/api/issues", issue.GetAllIssues)
	r.POST("/api/issue/add", issue.AddIssue)
	r.POST("/api/issue/update", issue.UpdateIssue)
	r.GET("/api/issue/status/get", issue.GetStatusReport)
	r.GET("/api/issue/new/get", issue.GetNewReport)
	r.GET("/api/issues/:id", issue.GetIssue)
	r.GET("/api/issue/delete/:id", issue.DeleteIssue)

	return r
}
