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

	return r
}
