package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/handler"
)

func UserRouteInit(apiGroup *gin.RouterGroup) {
	r := apiGroup.Group("/user")
	r.POST("/login", handler.HandlerLogin)
	r.POST("/register",handler.HandlerRegister)
}
