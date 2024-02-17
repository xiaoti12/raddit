package controller

type ResponseCode int

const (
	CodeSuccess ResponseCode = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeNeedLogin
	CodeInvalidToken
	CodeServerError
)

var codeMsgMap = map[ResponseCode]string{
	CodeSuccess:         "成功",
	CodeInvalidParams:   "参数信息不合法",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "密码错误",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的token",
	CodeServerError:     "服务出错",
}

func (c ResponseCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerError]
	}
	return msg
}
