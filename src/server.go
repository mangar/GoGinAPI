package main

import (
	"GoGinAPI/actions"
	"GoGinAPI/db"
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
			//log.Fatal(err)
		}
		con.Close()
		c.JSON(http.StatusOK, registros)
    })

	
	r.GET("/Logs", func(c *gin.Context) {
		c.JSON(http.StatusOK, actions.LogsAction(c))
    })


	r.GET("/AccountsGenerate", func(c *gin.Context) {
		c.JSON(http.StatusOK, actions.AccountsGenerateAction(c))
    })


    r.Run() // por padrão na porta 8080
}