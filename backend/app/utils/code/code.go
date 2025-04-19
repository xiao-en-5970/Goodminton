package code

type CodeMsg struct{
	Code int64
	Msg string
}

// 123举例，1代表模块名称，23代表模块下的错误码

var(
	//总体错误
	CodeAllSuccess = CodeMsg{Code: 101,Msg: "成功",}
	CodeAllIntervalError = CodeMsg{Code: 102,Msg: "服务器内部错误",}
	CodeAllRequestFormatError = CodeMsg{Code: 103,Msg: "请求格式错误",}
	CodeAllUnknownError = CodeMsg{Code: 104,Msg: "未知错误",}
	CodeAllBadGateway = CodeMsg{Code: 105,Msg: "错误网关",}
	//用户逻辑错误
	CodeUserLoginSuccess = CodeMsg{Code: 201,Msg: "登录成功",}
	CodeUserLoginSchoolAuthError = CodeMsg{Code: 202,Msg: "信息门户校验失败",}
	CodeUserLoginSchoolAuthSuccess = CodeMsg{Code: 203,Msg: "信息门户校验成功，跳转注册",}
	
)