package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handlerV1 := handler.NewHandlerV1(store)

	r.POST("/category", handlerV1.CreateCategory)
	r.GET("/category/:id", handlerV1.GetCategoryById)
	r.GET("/category", handlerV1.GetCategoryList)
	r.PUT("/category/:id", handlerV1.UpdateCategory)
	r.DELETE("/category/:id", handlerV1.DeleteCategory)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
