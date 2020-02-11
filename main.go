package main

import (
	"github.com/astaxie/beego"
	"github.com/backend/conf"
	"github.com/backend/database"
	_ "github.com/backend/routers"
	"github.com/gemnasium/logrus-graylog-hook"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
	initLogger()
	database.InitMySQL()
	beego.Run()
}

func initLogger() {
	addr := conf.Conf.GraylogAddr
	hook := graylog.NewGraylogHook(addr, map[string]interface{}{"service": "im-backend", "ENV": os.Getenv("ENV")})
	logrus.AddHook(hook)
	logrus.Info("initLogger")
}
