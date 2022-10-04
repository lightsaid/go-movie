package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const pagesPath = "./templates"
const staticPath = "./templates/static"

func (srv *Server) routes(router *gin.Engine) error {

	// 静态文件资源
	router.StaticFS("/static", http.Dir(staticPath))

	router.LoadHTMLGlob(fmt.Sprintf("%s/*.html", pagesPath))

	// 路由
	// r := router.Group("/sys")
	// r.GET("/ping", srv.ping)

	// 匹配所有 .page.html（页面）
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pagesPath))
	if err != nil {
		zap.S().Errorf("templates 匹配 pages 错误：%s", err)
		return fmt.Errorf("templates 匹配 pages 错误：%s", err)
	}

	// 匹配 .layout.html (布局)
	// layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pagesPath))
	// if err != nil {
	// 	zap.S().Errorf("templates 匹配 layouts 错误：%s", err)
	// 	return fmt.Errorf("templates 匹配 layouts 错误：%s", err)
	// }

	for _, page := range pages {
		// files := append([]string{page}, layouts...)
		// router.LoadHTMLFiles(files...)

		name := filepath.Base(page)
		url := strings.ReplaceAll(name, ".page", "")

		fmt.Println(">> name: ", name)

		router.GET(fmt.Sprintf("sys/%s", url), func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, name, gin.H{"pageTitle": "登陆"})
		})
	}

	return nil
}
