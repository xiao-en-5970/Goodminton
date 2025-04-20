package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao-en-5970/Goodminton/backend/app/model"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)

func LogicGetUserInfo(c *gin.Context,req *types.GetUserInfoReq)(resp *types.UserInfo,code int,err error){
	user,_:=model.FindUserByName(req.Username)
	if user!=nil{
		//用户存在
		userinfo,err:=model.ConvertUserToUserInfo(user)
		if err!=nil{
			return &types.UserInfo{},codes.CodeAllIntervalError,err
		}
		return userinfo,codes.CodeUserLoginSuccess,nil
	}else{
		
		return &types.UserInfo{},codes.CodeUserNotExist,nil
	}
}