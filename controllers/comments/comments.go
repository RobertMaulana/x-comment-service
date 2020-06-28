package comments

import (
	"github.com/RobertMaulana/x-comment-service/domain/comments"
	"github.com/RobertMaulana/x-comment-service/proto/comment"
	"github.com/RobertMaulana/x-comment-service/proto/common"
	"github.com/RobertMaulana/x-comment-service/services"
	"github.com/RobertMaulana/x-comment-service/utils/errors"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrganizationId(organizationName string) (*comments.Organization, *errors.RestErr){
	if organizationName == "" {
		return nil, errors.BadRequest("organization name can not be empty")
	}
	res, err := services.CommentService.GetOrganizationId(organizationName)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateComment(c *gin.Context) {
	organizationName := c.Param("organization_name")
	organization, err := GetOrganizationId(organizationName)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	var comment comments.CommentRequest
	if err := c.ShouldBindJSON(&comment); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := services.CommentService.CreateComment(comment, organization.Id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetComments(c *gin.Context) {
	organizationName := c.Param("organization_name")
	organization, err := GetOrganizationId(organizationName)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	comments, err := services.CommentService.GetComments(organization.Id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, comments)
}

func DeleteComments(c *gin.Context) {
	organizationName := c.Param("organization_name")
	organization, err := GetOrganizationId(organizationName)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	res, err := services.CommentService.DeleteComments(organization.Id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetOrganizationIdGrpc(request comment.OrganizationNameRequest) *common.Response {
	resp, err := services.CommentService.GetOrganizationIdGrpc(request.Name)
	if err != nil {
		return resp
	}
	return resp
}