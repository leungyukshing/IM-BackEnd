package database

import (
	"github.com/backend/conf"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var Server *gorm.DB

func InitMySQL() {
	log := logger()
	log.Info("InitMySQL start")

	dsn := conf.Conf.MysqlDsn
	Server, err := gorm.Open("mysql", dsn)

	if err == nil {
		log.Info("Open DB Success")
	} else {
		panic(err)
	}
	Server.SingularTable(true)
	log.Info("InitMySQL finish")
}

func logger() *logrus.Entry {
	fields := logrus.Fields{}
	return logrus.WithFields(fields)
}