package preload

import (
	"wechat-go/defines"
	"github.com/spf13/viper"
)

func Init()  {
	initConf()
}

func initConf() {
	confFile := defines.ProjectPath + "/conf/config.json"
	viper.SetConfigFile(confFile)
	if err := viper.ReadInConfig(); err != nil {
		// add log
		panic(err)
	}
}