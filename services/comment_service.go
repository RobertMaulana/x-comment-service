package services

import (
	"github.com/RobertMaulana/x-comment-service/domain/comments"
	"github.com/RobertMaulana/x-comment-service/proto/common"
	"github.com/RobertMaulana/x-comment-service/utils/errors"
)

var (
	CommentService commentServiceInterface = &commentService{}
)
type commentService struct {
}

type commentServiceInterface interface {
	GetOrganizationId(string) (*comments.Organization, *errors.RestErr)
	CreateComment(comments.CommentRequest, int64) (*comments.ApiCreationResponse, *errors.RestErr)
	GetComments(int64) (*comments.ApiListResponse, *errors.RestErr)
	DeleteComments(int64) (*comments.ApiGeneralResponse, *errors.RestErr)
	GetOrganizationIdGrpc(string) (*common.Response, *errors.RestErr)
}

func (s *commentService) GetOrganizationId(organizationName string) (*comments.Organization, *errors.RestErr) {
	dao := &comments.Organization{
		Name: organizationName,
	}
	res, err := dao.GetOrganizationId()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *commentService) CreateComment(comment comments.CommentRequest, organizationId int64) (*comments.ApiCreationResponse, *errors.RestErr) {
	if err := comment.Validate(); err != nil {
		return nil, err
	}
	dao := &comments.CommentRequest{
		Comment: comment.Comment,
		OrganizationId: organizationId,
	}
	result, err := dao.Save()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s * commentService) GetComments(organizationId int64) (*comments.ApiListResponse, *errors.RestErr) {
	dao := &comments.CommentRequest{
		OrganizationId: organizationId,
	}
	result, err := dao.GetComments()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *commentService) DeleteComments(organizationId int64) (*comments.ApiGeneralResponse, *errors.RestErr) {
	dao := &comments.CommentRequest{
		OrganizationId: organizationId,
	}
	result, err := dao.DeleteComments()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *commentService) GetOrganizationIdGrpc(organizationName string) (*common.Response, *errors.RestErr) {
	dao := &comments.Organization{
		Name: organizationName,
	}
	result, err := dao.GetOrganizationDataGrpc()
	if err != nil {
		return nil, err
	}
	return result, nil
}
