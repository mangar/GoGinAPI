package main

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)





func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Olá, Gin!")
    })

	r.GET("/version", func(c *gin.Context) {
        c.String(200, "v0.0.1")
    })

	r.GET("/ContratoEquipamentos", func(c *gin.Context) {
		var con = db.InitDB()
		registros , err := db.QueryNotificacoes(con)
		if err != nil {
			// log.Fatal(err)
		}
		con.Close()
		c.JSON(http.StatusOK, registros)
    })

	
	r.GET("/Logs", func(c *gin.Context) {
		ctx := context.Background()
		var con = db.InitDB()

		queries := dbsqlc.New(con)
		logs, err := queries.GetAllLogs(ctx)
		if err != nil {
			fmt.Println(err)
		}
		
		con.Close()
		c.JSON(http.StatusOK, logs)
    })

    r.Run() // por padrão na porta 8080
}