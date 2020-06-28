package comments

import (
	"github.com/RobertMaulana/x-comment-service/utils/errors"
)

type ApiGeneralResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type ApiCreationResponse struct {
	Status int `json:"status"`
	Data CommentRequest `json:"data"`
	Message string `json:"message"`
}

type CommentRequest struct {
	Id int64 `json:"id"`
	Comment string `json:"comment"`
	OrganizationId int64 `json:"organization_id"`
}

type ApiListResponse struct {
	Status int `json:"status"`
	Data []Comments `json:"data"`
	Message string `json:"message"`
}

type Comments struct {
	Id int64 `json:"id"`
	Comment string `json:"comment"`
	CreatedAt string `json:"created_at"`
}

type Organization struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

func (comment *CommentRequest) Validate() *errors.RestErr {
	if comment.Comment == "" {
		return errors.BadRequest("comment body is empty")
	}
	return nil
}
