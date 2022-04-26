package initialize

import (
	"chenxi/utils"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBClient *gorm.DB

type MysqlOptions struct {
	Engine            string `yaml:"engine" json:"engine"`
	Host              string `yaml:"host" json:"host"`
	Port              string `yaml:"port" json:"port"`
	Username          string `yaml:"username" json:"username"`
	Password          string `yaml:"password" json:"password"`
	Database          string `yaml:"database" json:"database"`
	MaxIdleConn       int    `yaml:"max_idle_conn" json:"max_idle_conn"`
	MaxOpenConn       int    `yaml:"max_open_conn" json:"max_open_conn"`
	ConnectionTimeout int    `yaml:"connection_timeout" json:"connection_timeout"`
	SlowLog           int    `yaml:"slow_log" json:"slow_log"`
}

func NewDBConnection() {
	db_config := Config.Database
	sys_config := Config.System

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local&timeout=%ds",
		db_config.Username,
		utils.AesDecrypt(db_config.Password, sys_config.SecurityKey),
		db_config.Host,
		db_config.Port,
		db_config.Database,
		db_config.ConnectionTimeout,
	)
	log.Println("数据库初始化......")
	db_conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panicf("数据库初始化连接异常：%s", err)
	}
	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sql_db, err := db_conn.DB()
	if err != nil {
		log.Panicf(err.Error())
	}
	// defer sql_db.Close()
	sql_db.SetMaxIdleConns(db_config.MaxIdleConn)
	sql_db.SetMaxOpenConns(db_config.MaxOpenConn)

	if err := sql_db.Ping(); err != nil {
		sql_db.Close()
		log.Panicf("数据库初始化失败......")
	}
	log.Println("数据库初始化成功......")
	DBClient = db_conn
}
