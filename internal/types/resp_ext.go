package types

// 统一成功返回
func Success(data ...any) *CommonResponse {
	var d any
	if len(data) > 0 {
		d = data[0]
	}
	return &CommonResponse{
		Code:    200,
		Message: "success",
		Data:    d,
		Success: true,
	}
}

// 统一失败返回（自定义错误码/信息）
func Fail(msg string) *CommonResponse {
	return &CommonResponse{
		Code:    500,
		Message: msg,
		Success: false,
	}
}
