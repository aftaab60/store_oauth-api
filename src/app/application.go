package app

import (
	"github.com/aftaab60/store_oauth-api/src/domain/access_token"
	"github.com/aftaab60/store_oauth-api/src/http"
	"github.com/aftaab60/store_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run("8080")
}
