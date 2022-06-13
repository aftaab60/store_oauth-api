package app

import (
	"github.com/aftaab60/store_oauth-api/src/clients/cassandra"
	"github.com/aftaab60/store_oauth-api/src/domain/access_token"
	"github.com/aftaab60/store_oauth-api/src/http"
	"github.com/aftaab60/store_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	cassandra.GetSession()

	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8081")
}
