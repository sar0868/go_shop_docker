package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go_shop_docker/docs"
	"go_shop_docker/internal/app"
	"go_shop_docker/internal/handler"
)

// @title HW15: shop
// @version 1
// @description API Server

// @host 0.0.0.0:8080/

func main() {
	conf := app.Init()

	IP := conf.HTTP.Host
	PORT := conf.HTTP.Port

	router := gin.Default()

	handle := handler.New()

	PATH := fmt.Sprintf("%s:%d", IP, PORT)
	url := ginSwagger.URL("http://" + PATH + "/swagger/doc.json")
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	handler.InitHandler(router, handle)

	if errRouter := router.Run(PATH); errRouter != nil {
		log.Panicln(errRouter.Error())
	}
}
