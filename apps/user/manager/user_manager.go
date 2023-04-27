package manager

import (
	"03dao-admin/apps/user/model"
	"03dao-admin/middleware/mysql"
	"github.com/gin-gonic/gin"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (u *UserManager) GetUserList(ctx *gin.Context, id uint64) *model.UserTable {
	user := &model.UserTable{
		UserName: "test",
		NickName: "test",
	}
	mysql.GormDB.Create(user)
	return &model.UserTable{}
}
