package errormessage

const (
	SUCCESS= 200
	ERROR = 500
	//CODE =1000 用户模块的错误
	ERROR_USERNAME_USE =1001
	ERROR_PASSWORD_WRONG=1002
	ERROR_USER_NOTEXIST=1003
	ERROR_TOKEN_EXISTS=1004
	ERROR_TOKEN_RUNTIME=1005
	ERROR_TOKEN_WRONG=1006
	ERROR_TOKEN_TYPR_ERONG=1007
	//2000开头的错误 文章模块的错误
	//3000开头的错误  分类模块的错误
)


var codemsg=map[int]string{
	SUCCESS :"OK",
	ERROR:"FALSE",
	ERROR_USERNAME_USE :"用户名已经存在！",
	ERROR_PASSWORD_WRONG :"密码错误",
	ERROR_USER_NOTEXIST :"用户不存在",
	ERROR_TOKEN_EXISTS :"token不存在",
	ERROR_TOKEN_RUNTIME: "token 超时了",
	ERROR_TOKEN_WRONG : "token错误",
	ERROR_TOKEN_TYPR_ERONG: "token格式错误",
}

func GetErrorMessage(code int )string{
	return  codemsg[code]
}