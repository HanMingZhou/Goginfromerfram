package bootstrap

import (
	"Goginfromerfram/global"
	"Goginfromerfram/routes"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 前端项目静态资源
	router.StaticFile("/", "./statis/dist/index.html")
	router.Static("/assets", "./statis/dist/assets")
	router.StaticFile("/favicon.ico", "./statis/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./storage/app/public")
	router.Static("/storage", "./storage/app/public")

	// 注册api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router

}

// RunServer 启动服务
func RunServer() {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen : %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shotdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("server exiting")

	
}
