package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	//ErrMsg string `json:"err_msg"`
}

const (
	SUCCESS             = 200
	ERROR               = 505
	ParamError          = 8000
	InternalServerError = http.StatusInternalServerError
)

const (
	SuccessMsg = "操作成功"
	FailMsg    = "操作失败"

	ParamErrorMsg          = "参数绑定失败, 请检查数据类型"
	InternalServerErrorMsg = "服务器内部错误"
)

var CustomError = map[int]string{
	SUCCESS:    SuccessMsg,
	ERROR:      FailMsg,
	ParamError: ParamErrorMsg,
}

func RspOK(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func RspFail(code int, data interface{}, msg string, c *gin.Context) {
	if msg == "" {
		c.JSON(http.StatusOK, Response{
			Code: code,
			Data: data,
			Msg:  CustomError[code],
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code: code,
			Data: data,
			Msg:  msg,
		})
	}
}

func OK(c *gin.Context) {
	RspOK(SUCCESS, map[string]interface{}{}, SuccessMsg, c)
}

func OKHasMsg(msg string, c *gin.Context) {
	RspOK(SUCCESS, map[string]interface{}{}, msg, c)
}

func OKHasData(data interface{}, c *gin.Context) {
	RspOK(SUCCESS, data, SuccessMsg, c)
}
func OKHasDetailed(data interface{}, msg string, c *gin.Context) {
	RspOK(SUCCESS, data, msg, c)
}

func Fail(c *gin.Context) {
	RspFail(ERROR, map[string]interface{}{}, FailMsg, c)
}

func FailHasMsg(code int, msg string, c *gin.Context) {
	RspFail(code, map[string]interface{}{}, msg, c)
}
func FailHasData(data interface{}, msg string, c *gin.Context) {
	RspFail(ERROR, data, msg, c)
}

func FailHasDetailed(code int, data interface{}, msg string, c *gin.Context) {
	RspFail(code, data, msg, c)
}
