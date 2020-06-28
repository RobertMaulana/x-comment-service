package app

import (
	"github.com/RobertMaulana/x-comment-service/controllers/comments"
	"github.com/RobertMaulana/x-comment-service/controllers/ping"
)

func mapUrls() {
	// Minikube need to check health pod using this open public endpoint
	router.GET("/ping", ping.Ping)

	// Comments route
	xendit := router.Group("/orgs")
	{
		xendit.POST("/:organization_name/comments", comments.CreateComment)
		xendit.GET("/:organization_name/comments", comments.GetComments)
		xendit.DELETE("/:organization_name/comments", comments.DeleteComments)
	}
}