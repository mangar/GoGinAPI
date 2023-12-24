package main

import (
	"GoGinAPI/actions"
	"GoGinAPI/db"
	"GoGinAPI/util"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func LogarDuracaoMiddleware(c *gin.Context) {
	inicio := time.Now()
	url := c.Request.URL.Path
	request_id := util.GenerateRequestId(url)
	util.Elastic( util.Document{RequestID: request_id, EntryPoint: url, Step: "start"} )

    c.Next() // Processa o request

    duracao := time.Since(inicio)
	util.Elastic( util.Document{RequestID: request_id, EntryPoint: url, Step: "end", Duracao: int(duracao.Seconds())  } )
}



func main() {
    r := gin.Default()
	r.Use(LogarDuracaoMiddleware)


    r.GET("/", func(c *gin.Context) {
        c.String(200, "Olá, Gin!")
    })

	r.GET("/version", func(c *gin.Context) {		
		version := "v.0.0.1"
        c.String(200, version)
    })

	r.GET("/ContratoEquipamentos", func(c *gin.Context) {
		var con = db.InitDB()
		registros , err := db.QueryNotificacoes(con)

		if err != nil {
			fmt.Println(err)
		}
		con.Close()

		c.JSON(http.StatusOK, registros)
    })

	
	r.GET("/Logs", func(c *gin.Context) {
		result, _ := actions.LogsAction(c)
		c.JSON(http.StatusOK, result)
    })


	r.GET("/AccountsGenerate", func(c *gin.Context) {
		result, _ := actions.AccountsGenerateAction(c)
		c.JSON(http.StatusOK, result)
    })


    r.Run() // por padrão na porta 8080
}



