package access_token

import (
	"github.com/aftaab60/store_oauth-api/src/utils/errors"
	"strings"
)

type Service interface {
	Create(token AccessToken) *errors.RestErr
	GetById(string) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(token *AccessToken) *errors.RestErr
}

//service requires any repository who has GetById implemented
type Repository interface {
	Create(token AccessToken) *errors.RestErr
	GetById(string) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(token *AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(token AccessToken) *errors.RestErr {
	if err := token.Validate(); err != nil {
		return err
	}
	token.AccessToken = strings.TrimSpace(token.AccessToken)

	if err := s.repository.Create(token); err != nil {
		return err
	}
	return nil
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) UpdateExpirationTime(token *AccessToken) *errors.RestErr {
	//TODO implement
	return nil
}
