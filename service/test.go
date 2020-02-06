// This file is used to test API connection of this service
package service

import (
	"github.com/astaxie/beego/context"
	"github.com/sirupsen/logrus"
)

func (this *Handler) Test() {
	ctx := *(this.Ctx)
	log := logger(ctx)
	log.Info("Test start")
	this.Ctx.Output.Body([]byte("connection success"))
}

func logger(ctx context.Context) *logrus.Entry {
	fields := logrus.Fields{}
	return logrus.WithFields(fields)
}