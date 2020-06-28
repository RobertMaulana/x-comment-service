package app

import (
	"github.com/RobertMaulana/x-comment-service/grpc"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	grpcServer := grpc.CommentsServer{}
	go grpcServer.Run()

	mapUrls()
	router.Run(":8080")
}