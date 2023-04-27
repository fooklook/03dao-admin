package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormDB *gorm.DB

func InitMysql() {
	if GormDB == nil {
		dsn := "root:hx2602966@tcp(127.0.0.1:3306)/admin-03dao?charset=utf8mb4&parseTime=True&loc=Local"
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}
		GormDB = conn
	}
}
