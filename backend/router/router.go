package router

import (
	"github.com/gin-gonic/gin"
	"github.com/JvWandermurem/cleanCode-golang/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/JvWandermurem/cleanCode-golang/docs"
)

func Setup(figurinhaHandler *handler.FigurinhaHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		figurinhas := v1.Group("/figurinhas")
		{
			figurinhas.POST("", figurinhaHandler.Create)       
			figurinhas.GET("", figurinhaHandler.List)         
			figurinhas.GET("/:id", figurinhaHandler.GetByID)   
			figurinhas.PATCH("/:id", figurinhaHandler.Update)  
			figurinhas.DELETE("/:id", figurinhaHandler.Delete) 
		}
	}

	return r
}