package main

import (
	"./config"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	db := config.DBInit()
	defer db.Close()
	router := gin.Default()
	config.InitRouter(router, db)
	ginpprof.Wrap(router)
	err := router.Run(":3333")
	if err != nil {
		log.Print(err)
	}

}
