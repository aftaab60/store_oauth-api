package http

import (
	"github.com/aftaab60/store_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

//this is handler file, same as controller in MVC pattern

type AccessTokenHandler interface {
	Create(c *gin.Context)
	GetById(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var accessToken access_token.AccessToken
	if err := c.ShouldBindJSON(&accessToken); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := h.service.Create(accessToken); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := h.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	//TODO implement
}
