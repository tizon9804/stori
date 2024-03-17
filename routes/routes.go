package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"stori/docs"
	"stori/postgres"
	"stori/transactions"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	prefix := "api/stori"
	api := engine.Group(prefix)

	docs.SwaggerInfo.BasePath = prefix
	serveDocs(api)

	api.GET("health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "0.0.1",
		})
	})

	newPostgres, err := postgres.NewPostgres()
	if err != nil {
		panic(err)
	}
	storage := transactions.NewStorage(newPostgres)
	transactionService := transactions.NewTransaction(storage)
	api.POST("/transaction/upload", transactions.TransactionFileHandler(transactionService))

	return engine
}

// this function is the handler for the render documentation
func serveDocs(api *gin.RouterGroup) {
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL("doc.json")))
}
