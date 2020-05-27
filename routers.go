package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(apmgin.Middleware(router))
	router.GET("/customerinfo", getCustomer)
	router.POST("/customerinfo", createCustomer)
	router.GET("/customerinfo/:cifno", getCustomerDetail)
	router.PUT("/customerinfo/:cifno", editCustomer)

	return router
}
