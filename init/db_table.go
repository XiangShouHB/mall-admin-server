package boot

import (
	"imooc.com/shop-admin/mall-admin-server/app/model"
	"imooc.com/shop-admin/mall-admin-server/library/global"
)

// 注册数据库表专用
func DBTables() {
	db := global.DB
	db.AutoMigrate(model.SpUser{})
}
