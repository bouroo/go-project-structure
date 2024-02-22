package datasources

import (
	"github.com/bouroo/go-project-structure/infrastructure"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	AppConfig    *viper.Viper
	DBConn       *gorm.DB
	RedisConn    *infrastructure.RedisConn
	UserGRPCConn *infrastructure.GRPCConnPool
)
