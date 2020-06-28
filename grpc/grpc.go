package grpc

import (
	"context"
	"github.com/RobertMaulana/x-comment-service/controllers/comments"
	"github.com/RobertMaulana/x-comment-service/proto/comment"
	"github.com/RobertMaulana/x-comment-service/proto/common"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type CommentsServer struct {
}

func (CommentsServer) Run() {
	port := os.Getenv("GRPC_PORT")
	log.Println("Starting RPC server at", port)

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unfortunately could not listen to %s: %v", port, err)
	}

	srv := grpc.NewServer()
	var commentServer CommentsServer
	comment.RegisterCommentsServer(srv, commentServer)

	log.Fatal(srv.Serve(l))
}

func (CommentsServer) GetOrganizationMembers(_ context.Context, request *comment.OrganizationNameRequest) (*common.Response, error) {
	result := comments.GetOrganizationIdGrpc(*request)
	return result, nil
}
