package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() {
	file, err := os.OpenFile("log/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		return
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(file)
	logrus.SetReportCaller(false)
}
