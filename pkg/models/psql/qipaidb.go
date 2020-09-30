package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/project-zero-233/userapisrv/pkg/config"
)

// 创建qipaidb连接
func NewQipaiDBConn(dbName string, tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(dbName)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", dbName))
	}

	QiPaiDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大链接数
	QiPaiDB.DB().SetMaxOpenConns(10)
}
