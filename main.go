package main

import (
	"flag"
	"fmt"
	"github.com/project-zero-233/logging"
	"github.com/project-zero-233/userapisrv/pkg/config"
	"github.com/project-zero-233/userapisrv/pkg/define"
	"github.com/project-zero-233/userapisrv/pkg/models/psql"
	"github.com/project-zero-233/userapisrv/pkg/models/redis"
	"github.com/project-zero-233/userapisrv/pkg/router"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

var (
	tomlFilePath, mode string
)

func init() {
	flag.StringVar(&tomlFilePath, "config", "docs/local.toml", "服务配置文件")
	flag.StringVar(&mode, "mode", "release", "模型-debug还是release还是test")

	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 入口函数
func main() {
	defer func() {
		// 停止服务
		logging.Warnf("stop server!")
		err := recover()
		if err != nil {
			fmt.Printf("err: %v, debug.Stack():%v", err, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	// 解析配置文件
	conf, err := config.UnmarshalConfig(tomlFilePath)
	if err != nil {
		panic(err)
	}

	// 初始化engine
	engine, err := router.InitEngine(mode)
	if err != nil {
		panic(err)
	}

	// 初始化 DB
	initDB(conf)

	// 启动服务
	logging.Infof("run userapisrv at %v\n", conf.GetListenAddr())
	go engine.Run(conf.GetListenAddr())

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGKILL)
	<-signalChan
}

func initDB(conf *config.Config) {
	// 设置db连接
	psql.NewQipaiDBConn(define.QiPaiDB, conf)

	psql.NewStatDBConn(define.StatDB, conf)

	// 设置Redis连接
	redis.NewVipRedisConn(define.VipCache, conf)
}
