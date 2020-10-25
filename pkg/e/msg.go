package e

// MsgFlags 错误码详解
var MsgFlags = map[string]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	NOT_FOUND:      "资源不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL: "身份校验失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 过期",
	LOGIN_FAIL: "登录失败，请稍后重试",
}
