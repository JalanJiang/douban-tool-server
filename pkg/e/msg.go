package e

// MsgFlags 错误码详解
var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	NOT_FOUND:      "资源不存在",
}
