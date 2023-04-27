package route

import (
	"03dao-admin/common"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	Method   string      //Method is one of the following: GET,PUT,POST,DELETE. required
	Path     string      //Path contains a path pattern. required
	Handler  Handler     //the func this API calls. you must set this field or ResourceFunc, if you set both, ResourceFunc will be used
	Response interface{} // Response
	Request  interface{} //Request
	Name     string      //Name
}

const (
	ResponseOk          = 0
	ResponseSystemError = 500
	ResponseInvalidAuth = 401
)

type HTTPJSONResponse struct {
	RetCode int         `json:"ret_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Prompt  string      `json:"prompt"` // 操作成功，但前端进行友善提示内容。
}

type Handler func(ctx context.Context) (interface{}, error)

func FriendlyPromptJSON(handle func(ctx *gin.Context) (interface{}, string, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, prompt, err := handle(ctx)
		if err != nil {
			err1, ok := err.(common.ErrorCode)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, HTTPJSONResponse{
					RetCode: ResponseSystemError,
					Message: "System Error",
				})
				return
			}
			ctx.JSON(http.StatusOK, HTTPJSONResponse{
				RetCode: err1.Code,
				Message: err1.Message,
			})
			return
		}
		ctx.JSON(http.StatusOK, HTTPJSONResponse{
			RetCode: ResponseOk,
			Data:    data,
			Prompt:  prompt,
		})
	}
}

func JSON(handle func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := handle(ctx)
		if err != nil {
			err1, ok := err.(common.ErrorCode)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, HTTPJSONResponse{
					RetCode: ResponseSystemError,
					Message: "System Error",
				})
				return
			}
			ctx.JSON(http.StatusOK, HTTPJSONResponse{
				RetCode: err1.Code,
				Message: err1.Message,
			})
			return
		}
		ctx.JSON(http.StatusOK, HTTPJSONResponse{
			RetCode: ResponseOk,
			Data:    data,
		})
	}
}

func UserJSON(handle func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 验证用户权限
		data, err := handle(ctx)
		if err != nil {
			err1, ok := err.(common.ErrorCode)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, HTTPJSONResponse{
					RetCode: ResponseSystemError,
					Message: "System Error",
				})
				return
			}
			ctx.JSON(http.StatusOK, HTTPJSONResponse{
				RetCode: err1.Code,
				Message: err1.Message,
			})
			return
		}
		ctx.JSON(http.StatusOK, HTTPJSONResponse{
			RetCode: ResponseOk,
			Data:    data,
		})
	}
}
