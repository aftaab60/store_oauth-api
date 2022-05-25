package db

import (
	"github.com/aftaab60/store_oauth-api/src/domain/access_token"
	"github.com/aftaab60/store_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetAccessTokenById(id string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetAccessTokenById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
