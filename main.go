package main

import (
	"MoneyBackward/global"
	"MoneyBackward/setup"
	"fmt"
	"log"
)

func main() {
	//init config
	const filepath = "conf.yaml"
	confList, err := setup.InitConfFrom(filepath)
	if err != nil {
		panic(fmt.Errorf("server crashed: %s", err))
	} else {
		global.Conf = confList
		log.Printf("init config from %v success, conf now at %v\n with content %v\n", filepath, &global.Conf, global.Conf)
	}

	//init logger
	global.Logger = setup.InitLogger()
	global.Logger.Info(fmt.Sprintf("init logger success, logger now at %v\n", &global.Logger))

	//init gorm
	global.MysqlDB, err = setup.InitGorm()
	if err != nil {
		panic(fmt.Errorf("server crashed: %s", err))
	}
	global.Logger.Info(fmt.Sprintf("init mysql success, mysqlDB now at %v\n", &global.MysqlDB))

	//init router
	router := setup.InitRouter()
	if err := router.Run(global.Conf.Server.Addr()); err != nil {
		panic(fmt.Errorf("server crashed: %s", err))
	} else {
		global.Logger.Info(fmt.Sprintf("init router success, router now at %v\n", &router))
	}
}
