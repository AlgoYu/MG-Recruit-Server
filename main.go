package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"machine-geek.cn/recruit-server/database"
	"machine-geek.cn/recruit-server/route"
)

func main() {
	defer database.DB.Close()
	r := gin.Default()
	route.Route(r)
	r.Run(":8888")
}
