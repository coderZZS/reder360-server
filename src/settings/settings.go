package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称
	viper.SetConfigType("yaml")   // 指定文件类型
	viper.AddConfigPath(".")      // 指定查找配置文件的相对路径
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
