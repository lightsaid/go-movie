package main

import (
	"database/sql"
	"log"
	"time"

	"lightsaid.com/go-movie/booking/cmd/api/server"
	db "lightsaid.com/go-movie/booking/db/sqlc"
	"lightsaid.com/go-movie/booking/utils"

	_ "github.com/lib/pq"
)

func init() {
	time.LoadLocation("Asia/Shanghai")
}

func main() {
	// 加载配置
	var config utils.ApiConfig
	if err := utils.LoadConfig(".", "ApiConfig", &config); err != nil {
		log.Fatalf("loading config file error: %v", err)
	}

	// 加载数据库
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("connect database error: %v", err)
	}
	defer conn.Close()
	q := db.New(conn)

	// 加载日志配置, 输出到终端和api.log文件
	paths := []string{"./api.log", "stderr"}
	logger, err := utils.NewLogger(paths...)
	if err != nil {
		log.Fatal("init zap logger error: ", err)
	}
	defer logger.Sync()
	// zap.S().Error("测试日志～")

	// 加载Server
	srv := server.NewServer(&config, q)

	// 启动服务
	srv.Run()
}
