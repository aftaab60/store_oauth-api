package db

import (
	"github.com/aftaab60/store_oauth-api/src/clients/cassandra"
	"github.com/aftaab60/store_oauth-api/src/domain/access_token"
	"github.com/aftaab60/store_oauth-api/src/utils/errors"
)

var (
	queryCreateAccessToken    = "insert into access_token(access_token_id, user_id, client_id, expires) values(?,?,?,?);"
	queryGetById              = "select access_token_id, user_id, client_id, expires from access_token where access_token_id=?;"
	queryUpdateExpirationTime = "update access_token set expires=? where access_token_id=?;"
)

type DbRepository interface {
	Create(token access_token.AccessToken) *errors.RestErr
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(token *access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()
	if err := session.Query(queryCreateAccessToken, token.AccessToken, token.UserId, token.ClientId, token.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()

	var accessToken access_token.AccessToken
	if err := session.Query(queryGetById, id).Scan(
		&accessToken.AccessToken,
		&accessToken.UserId,
		&accessToken.ClientId,
		&accessToken.Expires); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &accessToken, nil
}

func (r *dbRepository) UpdateExpirationTime(token *access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()

	if err := session.Query(queryUpdateExpirationTime, token.Expires, token.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
