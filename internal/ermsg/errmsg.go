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
	FROM_FORMAT_ERROR    = 708
	FILE_SAVE_ERROR      = 709

	HAVING_NO_PICTURE    = 710
	PICTURE_FORM_ERROR   = 711
	PICTURE_DNCODE_ERROR = 712

	// 个人主页模块错误

	// 用户详细信息错误
	PASSWORD_ERROR = 901
	// 发送转发模块错误

)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	// 用户登陆注册模块错误
	TOKEN_ERROR:        "token错误",
	TOKEN_NOT_FOUND:    "无token",
	TOKEN_FORMAT_ERROR: "token格式错误",
	TOKEN_NOT_VALID:    "token无效或已过期",

	ERROR_ACCOUNT_HAS_EXISTED: "账号已存在",
	ERROR_ACCOUNT_NOT_EXIST:   "账号不存在",
	ERROR_PASSWORD_ERROR:      "密码错误",

	// 主页模块错误
	POST_FORMAT_ERROR:      "帖子格式错误",
	POST_WRITE_ERROR:       "帖子写入数据库错误",
	UID_IS_NOT_EXIST:       "uid不存在",
	FEED_WRITE_ERROR:       "Feed表写入错误",
	RELATEDATA_WRITE_ERROR: "用户相关信息表写入错误",

	FORWARD_FORMAT_ERROR: "转发帖子格式错误",
	FORWARD_WRITE_ERROR:  "转发表写入错误",

	FROM_FORMAT_ERROR: "上传文件错误",
	FILE_SAVE_ERROR:   "文件写入错误",

	HAVING_NO_PICTURE:    "没有图片",
	PICTURE_FORM_ERROR:   "图片格式错误",
	PICTURE_DNCODE_ERROR: "图片解码错误",

	// 用户详细信息错误
	PASSWORD_ERROR: "原密码错误",
}

//通过错误码获取错误信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}
