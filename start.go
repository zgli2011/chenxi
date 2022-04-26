package main

import (
	"chenxi/initialize"
	"chenxi/web/router"
	"flag"
)

// init函数先于main函数自动执行
func init_source(config_path string) {
	initialize.LoadConfiguration(config_path) // 加载配置文件
	initialize.NewDBConnection()              // 初始化数据库连接
	initialize.NewRedisClient()               // 初始化redis
	initialize.NewLog()                       // 初始化
}

func main() {
	var config_file string
	flag.StringVar(&config_file, "config", "./config.yaml", "配置文件目录")
	flag.Parse()
	init_source(config_file)
	router.Router()
}
