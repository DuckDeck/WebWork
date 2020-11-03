package initialize

import (
	"WebWork/global"
	"WebWork/model"
)

func DBTables() {
	db := global.G_DB
	db.SingularTable(true)
	db.AutoMigrate(model.FiveCode{})
	global.G_LOG.Debug("register table success")
}
