package main

import (
	"./config"
	"./libraries"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	db := config.DBInit()
	defer db.Close()
	router := gin.Default()
	//router.Use(gin.Recovery())
	//router.Use(gin.Logger())
	router.Use(libraries.GinJwt)
	config.InitRouter(router, db)
	ginpprof.Wrap(router)
	err := router.Run(":3333")
	if err != nil {
		log.Print(err)
	}

}
