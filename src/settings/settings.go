package settings

import (
	"fmt"
	"gin-cli/src/utils"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

var Conf = new(AppConfig)

func WriteConf() {
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
	}
}

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

	//读取配置反序列化
	WriteConf()

	viper.WatchConfig() // 监听配置更新
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file updated...")
		WriteConf()
	})
	return nil
}
