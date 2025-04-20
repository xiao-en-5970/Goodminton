package types

type GetUserInfoReq struct {
    Username string `json:"username" binding:"required,alphanum,min=4,max=20"`
}