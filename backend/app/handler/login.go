package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/logic"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/responce"
)
// @Summary 用户登录
// @Description 使用用户名密码登录系统
// @Tags 认证服务
// @Accept json
// @Produce json
// @Param login body model.LoginReq true "登录凭证"
// @Success 200 {object} model.LoginResp "登录成功"
// @Router /login [post]
func HandlerLogin(c *gin.Context){
	rawData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	req := &types.LoginReq{}
	err=json.Unmarshal(rawData,req)
	if err!=nil{
		responce.ErrorStatusBadRequest(c,err)
		return
	}
	resp,err:=logic.LogicLogin(c,req)
	if err!=nil{
		responce.ErrorStatusInternalServerError(c,err)
		return
	}
	responce.SuccessWithData(c,*resp)

}