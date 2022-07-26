package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ping 测试 api 是否还在运行
func (s *Server) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping !"})
}
