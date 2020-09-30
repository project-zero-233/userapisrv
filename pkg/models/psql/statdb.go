package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/project-zero-233/userapisrv/pkg/config"
)

// 创建 statdb 连接
func NewStatDBConn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	StatDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大链接数
	StatDB.DB().SetMaxOpenConns(10)
}
