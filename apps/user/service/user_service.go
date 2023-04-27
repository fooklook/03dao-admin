package service

import (
	"03dao-admin/apps/user/manager"
	"03dao-admin/apps/user/model"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userManager *manager.UserManager
}

func NewUserService(userManager *manager.UserManager) *UserService {
	return &UserService{
		userManager: userManager,
	}
}

func (us *UserService) GetUserList(ctx *gin.Context) *model.UserTable {
	return us.userManager.GetUserList(ctx, 1)
}
