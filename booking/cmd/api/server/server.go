package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	db "lightsaid.com/go-movie/booking/db/sqlc"
	"lightsaid.com/go-movie/booking/utils"
)

type Server struct {
	config  *utils.ApiConfig
	querier db.Querier
	router  *gin.Engine
}

// NewServer 创建一个 api server
func NewServer(config *utils.ApiConfig, querier db.Querier) *Server {
	srv := &Server{
		config:  config,
		querier: querier,
	}
	gin.SetMode(config.RunMode)
	router := gin.New()
	if config.RunMode == "debug" {
		router.Use(gin.Logger())
	}
	router.Use(gin.Logger())

	router.Group("/api/v1/")
	{
		router.GET("/ping", srv.ping)
		router.GET("/cities", srv.getCity)
	}

	srv.router = router
	return srv
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:           s.config.ApiPort,
		Handler:        s.router,
		IdleTimeout:    time.Minute,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("srv.ListenAndServe err: %v", err)
		}
	}()

	// 优雅关机
	quit := make(chan os.Signal, 1)

	// 捕获信号, ctrl+c、kill pid （kill -9 pid 强制退出，无法捕获）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quit

	zap.S().Info("Shuting down server signal: ", sig.String())

	// 创建一个20秒后超时上下文, 等待系统逻辑执行
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		// 可能超时了
		zap.S().Fatal("Server shutdown error: ", err)
	}

	// 所有逻辑处理完成
	zap.S().Info("Stopped Api Server Success.")

}
