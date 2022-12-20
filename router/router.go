package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zeimedee/slightlytechie/controllers"
)

func SetUpRouter() *gin.Engine {

	router := gin.Default()

	apiv1 := router.Group("/api/v1")
	{
		apiv1.POST("/add", controllers.AddPost)
		apiv1.GET("/read/:id", controllers.ReadPost)
		apiv1.GET("/readall", controllers.ReadAllPost)
		apiv1.PATCH("/update", controllers.UpdatePost)
		apiv1.DELETE("/delete", controllers.DeletePost)
	}

	return router
}
