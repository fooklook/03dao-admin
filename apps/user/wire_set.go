package user

import (
	"03dao-admin/apps/user/manager"
	"03dao-admin/apps/user/service"
	"03dao-admin/apps/user/view"
	"github.com/google/wire"
)

var ManagerSet = wire.NewSet(manager.NewUserManager)

var ServiceSet = wire.NewSet(service.NewUserService)

var ViewSet = wire.NewSet(view.NewUserView)

var UserSet = wire.NewSet(ManagerSet, ServiceSet, ViewSet)
