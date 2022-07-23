package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ping 测试 api 是否还在运行
func (s *Server) ping(c *gin.Context) {
	time.Sleep(20 * time.Second)
	c.JSON(http.StatusOK, gin.H{"message": "Ping !"})
}
