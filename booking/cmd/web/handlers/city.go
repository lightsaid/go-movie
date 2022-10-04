package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getCity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "City"})
}
