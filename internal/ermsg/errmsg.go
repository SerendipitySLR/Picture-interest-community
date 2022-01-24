package errmsg

//按照自己的模块分类，从600用户注册、主页700、个人主页800、用户详细信息900、发送转发1000
const (
	//通用
	SUCCESS = 200
	ERROR   = 500

	// 用户注册模块错误
	TOKEN_ERROR               = 600
	TOKEN_NOT_FOUND           = 601
	TOKEN_FORMAT_ERROR        = 602
	TOKEN_NOT_VALID           = 603
	ERROR_ACCOUNT_HAS_EXISTED = 604
	ERROR_ACCOUNT_NOT_EXIST   = 605
	ERROR_PASSWORD_ERROR      = 606

	// 主页模块错误
	POST_FORMAT_ERROR      = 701
	POST_WRITE_ERROR       = 702
	UID_IS_NOT_EXIST       = 703
	FEED_WRITE_ERROR       = 704
	RELATEDATA_WRITE_ERROR = 705

	FORWARD_FORMAT_ERROR = 706
	FORWARD_WRITE_ERROR  = 707
	// 个人主页模块错误

	// 用户详细信息错误

	// 发送转发模块错误

)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	TOKEN_ERROR:        "token错误",
	TOKEN_NOT_FOUND:    "无token",
	TOKEN_FORMAT_ERROR: "token格式错误",
	TOKEN_NOT_VALID:    "token无效或已过期",
}

//通过错误码获取错误信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}
