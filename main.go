package main

import (
	"03dao-admin/apps/user/manager"
	"03dao-admin/apps/user/service"
	"03dao-admin/apps/user/view"
	"03dao-admin/middleware/logger"
	"03dao-admin/middleware/mysql"
	"03dao-admin/middleware/route"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	mysql.InitMysql()
	logger.InitLogger()
	userManger := manager.NewUserManager()
	userService := service.NewUserService(userManger)
	user := view.NewUserView(userService)
	app.Handle("GET", "/index", route.JSON(user.GetUserList))
	app.Handle("GET", "/user/home", route.UserJSON(user.GetUserList))
	err := app.Run()
	if err != nil {
		println(err)
	}
}
