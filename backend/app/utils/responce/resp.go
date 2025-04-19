package responce

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/code"
)



func Success(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":code.CodeAllSuccess.Code,
		"msg":code.CodeAllSuccess.Msg,
		"data":"null",
	})
}
func SuccessWithCode(c *gin.Context,code code.CodeMsg){
	c.JSON(http.StatusOK,gin.H{
		"code":code.Code,
		"msg":code.Msg,
		"data":"null",
	})
}

func SuccessWithMsg(c *gin.Context,msg string){
	c.JSON(http.StatusOK,gin.H{
		"code":code.CodeAllSuccess.Code,
		"msg":msg,
		"data":"null",
	})
}

func SuccessWithData(c *gin.Context,data interface{}){
	c.JSON(http.StatusOK,gin.H{
		"code":code.CodeAllSuccess.Code,
		"msg":code.CodeAllSuccess.Msg,
		"data":data,
	})
}

func ErrorStatusBadRequest(c *gin.Context,err error){
	c.JSON(http.StatusBadRequest,gin.H{
		"code":code.CodeAllRequestFormatError.Code,
		"msg":code.CodeAllRequestFormatError.Msg,
		"err":err.Error(),
	})
}

func ErrorStatusBadGateway(c *gin.Context,err error){
	c.JSON(http.StatusBadGateway,gin.H{
		"code":code.CodeAllBadGateway.Code,
		"msg":code.CodeAllBadGateway.Msg,
		"err":err.Error(),
	})
}

func ErrorStatusInternalServerError(c *gin.Context,err error){
	c.JSON(http.StatusInternalServerError,gin.H{
		"code":code.CodeAllIntervalError.Code,
		"msg":code.CodeAllIntervalError.Msg,
		"err":err.Error(),
	})
}
