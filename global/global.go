package global

import (
	"MoneyBackward/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Conf *conf.ListOfConf
var Logger *logrus.Logger
var MysqlDB *gorm.DB
