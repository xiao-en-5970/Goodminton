package logic

import (
	"github.com/gin-gonic/gin"

	"github.com/xiao-en-5970/Goodminton/backend/app/model"
	"github.com/xiao-en-5970/Goodminton/backend/app/types"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/bcrypts"
	"github.com/xiao-en-5970/Goodminton/backend/app/utils/codes"
)

func LogicLogin(c *gin.Context,req *types.LoginReq)(resp *types.UserInfo,code int,err error){
	
	user,_:=model.FindUserByName(req.Username)
	
	if user!=nil{
		//用户存在
		if (bcrypts.CheckPasswordHash(user.PasswordHash,req.Password)){
			//账密正确
			userinfo,err:=model.ConvertUserToUserInfo(user)
			if err!=nil{
				return &types.UserInfo{},codes.CodeAllIntervalError,err
			}
			return userinfo,codes.CodeUserLoginSuccess,nil
		}else{
			//账密错误
			return &types.UserInfo{},codes.CodeUserLoginPasswordError,nil
		}
	}else{
		// 用户不存在
		code,err:=LogicLoginHFUT(&types.LoginHFUTReq{
			Username: req.Username,
			Password: req.Password,
		})
		return &types.UserInfo{},code,err
	}
}
