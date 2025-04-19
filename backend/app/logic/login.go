package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
)

func LogicLogin(c *gin.Context,req *types.LoginReq)(resp *types.LoginResp,err error){
	return &types.LoginResp{
		UserId: 12,
		Username: "2022210826",
		Nickname: "xiao_en",
		},nil
}