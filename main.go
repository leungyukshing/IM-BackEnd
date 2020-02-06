package main

import (
	"github.com/astaxie/beego"
	"github.com/backend/database"
	_ "github.com/backend/routers"
	"github.com/gemnasium/logrus-graylog-hook"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	/*
		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   10,
		}).Info("A group of walrus emerges from the ocean")

		log.WithFields(log.Fields{
			"omg":    true,
			"number": 122,
		}).Warn("The group's number increased tremendously!")*/
	initLogger()
	database.InitMySQL()
	beego.Run()
}

func initLogger() {
	var addr string
	if env := os.Getenv("ENV"); env == "product" {
		addr = "47.100.58.6:12201"
	} else {
		addr = "127.0.0.1:12201"
	}
	hook := graylog.NewGraylogHook(addr, map[string]interface{}{"service": "im-backend", "ENV": os.Getenv("ENV")})
	logrus.AddHook(hook)
	logrus.Info("initLogger")
}