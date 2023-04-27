package view

import (
	"03dao-admin/apps/user/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserView struct {
	userService *service.UserService
}

func NewUserView(userService *service.UserService) *UserView {
	return &UserView{
		userService: userService,
	}
}

func (u *UserView) GetUserList(ctx *gin.Context) (interface{}, error) {
	u.userService.GetUserList(ctx)
	logrus.Infof("test %+v", "GetUserList")
	return "user list", nil
}
