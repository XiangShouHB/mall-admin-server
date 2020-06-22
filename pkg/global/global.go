package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"imooc.com/shop-admin/mall-admin-server/config"
)

var (
	DB     *gorm.DB
	Redis  *redis.Client
	Config config.Server
	VP     *viper.Viper
	LOG    *oplogging.Logger
)
