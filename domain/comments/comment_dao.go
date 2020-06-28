package comments

import (
	"database/sql"
	"encoding/json"
	"github.com/RobertMaulana/x-comment-service/datasource/postgre/comments_db"
	"github.com/RobertMaulana/x-comment-service/logger"
	"github.com/RobertMaulana/x-comment-service/proto/common"
	"github.com/RobertMaulana/x-comment-service/query"
	"github.com/RobertMaulana/x-comment-service/utils/errors"
	"net/http"
	"strings"
	"time"
)

func createOrganization(organizationName string) (*Organization, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.CreateOrganization)
	if err != nil {
		logger.Error("error when trying to prepare create organization statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()

	var organizationResp Organization
	err = stmt.QueryRow(strings.ToLower(organizationName)).Scan(&organizationResp.Id)
	if err != nil {
		logger.Error("error when trying to create organization statement", err)
		return nil, errors.InternalServerError("database error")
	}
	res := &Organization{
		Id: organizationResp.Id,
		Name: organizationName,
	}
	return res, nil
}

func (organization *Organization) GetOrganizationId() (*Organization, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.GetOrganizationId)
	if err != nil {
		logger.Error("error when trying to prepare get organization id statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(organization.Name)
	if err := result.Scan(&organization.Id, &organization.Name); err != nil {
		if err == sql.ErrNoRows {
			newOrganization, saveErr := createOrganization(organization.Name)
			if saveErr != nil {
				logger.Error("error when trying to create organization", err)
				return nil, errors.InternalServerError("database error")
			}
			newOrganizationResp := &Organization{
				Id: newOrganization.Id,
				Name: newOrganization.Name,
			}
			return newOrganizationResp, nil
		}
		logger.Error("error when trying to execute get organization id", err)
		return nil, errors.InternalServerError("database error")
	}
	res := &Organization{
		Id: organization.Id,
		Name: organization.Name,
	}
	return res, nil
}

func (comment *CommentRequest) Save() (*ApiCreationResponse, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.CreateComment)
	if err != nil {
		logger.Error("error when trying to prepare save comment statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()

	var commentResp CommentRequest
	saveErr := stmt.QueryRow(comment.Comment, comment.OrganizationId).Scan(&commentResp.Id)
	if saveErr != nil {
		logger.Error("error when trying to save comment statement", saveErr)
		return nil, errors.InternalServerError("database error")
	}
	res := &ApiCreationResponse{
		Status: http.StatusCreated,
		Data: CommentRequest{
			Id: commentResp.Id,
			Comment: comment.Comment,
			OrganizationId: comment.OrganizationId,
		},
		Message: "comment is successful created",
	}
	return res, nil
}

func (comment *CommentRequest) GetComments() (*ApiListResponse, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.GetAllComments)
	if err != nil {
		logger.Error("error when trying to prepare get comments info statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(comment.OrganizationId)
	if err != nil {
		logger.Error("error when trying to get comments info statement", err)
		return nil, errors.InternalServerError("database error")
	}
	var results []Comments
	for rows.Next() {
		var res Comments
		if err := rows.Scan(&res.Id, &res.Comment, &res.CreatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.NotFound("data not found")
			}
			logger.Error("error when trying to execute investor info", err)
			return nil, errors.InternalServerError("database error")
		}
		results = append(results, res)
	}
	if results == nil {
		results = []Comments{}
	}
	res := &ApiListResponse{
		Status: http.StatusOK,
		Data: results,
		Message: "success",
	}
	return res, nil
}

func (comment *CommentRequest) DeleteComments() (*ApiGeneralResponse, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.DeleteComments)
	if err != nil {
		logger.Error("error when trying to prepare delete comments statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()

	now := time.Now()
	if _, err := stmt.Exec(now, comment.OrganizationId); err != nil {
		logger.Error("error when trying to delete comments statement", err)
		return nil, errors.InternalServerError("database error")
	}
	res := &ApiGeneralResponse{
		Status: http.StatusOK,
		Message: "all comments are successful removed",
	}
	return res, nil
}

func (organization *Organization) GetOrganizationDataGrpc() (*common.Response, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.GetOrganizationIdByName)
	if err != nil {
		logger.Error("error when trying to prepare get organization id statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow("%"+organization.Name+"%")
	if err := result.Scan(&organization.Id, &organization.Name); err != nil {
		if err == sql.ErrNoRows {
			byteResp, _ := json.Marshal(organization)
			res := &common.Response{
				Status: http.StatusOK,
				Data: byteResp,
				Message: "success",
			}
			return res, nil
		}
		logger.Error("error when trying to execute get organization id", err)
		return nil, errors.InternalServerError("database error")
	}
	byteResp, _ := json.Marshal(organization)
	res := &common.Response{
		Status: http.StatusOK,
		Data: byteResp,
		Message: "success",
	}
	return res, nil
}