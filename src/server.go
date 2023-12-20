package main

import (
	"GoGinAPI/db"

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

		db.QueryNotificacoes(con)

		con.Close()

		// var conn = db.initDB()

		// var d = db.initDB()
		// db.queryTable(d)


        c.String(200, "Contrato Equipamentos")
    })

    r.Run() // por padrão na porta 8080
}