package settings

import (
	"fmt"
	"gin-cli/src/utils"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile(utils.GetDirPath() + "/config.yaml")
	//viper.SetConfigName("config") // 指定配置文件名称
	//viper.SetConfigType("yaml")   // 指定文件类型(获取远程配置时使用)
	//viper.AddConfigPath(".") // 指定查找配置文件的相对路径
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return
	}
	viper.WatchConfig() // 监听配置更新
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file updated...")
	})
	return nil
}
