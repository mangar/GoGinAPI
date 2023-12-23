package main

import (
	"GoGinAPI/actions"
	"GoGinAPI/db"
	"GoGinAPI/util"
	"encoding/json"
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
		
		util.Elastic( util.Document{API: "/version", Status: "start"} )
		version := "v.0.0.1"
		util.Elastic( util.Document{API: "/version", Status: "end", Info: "Version:" + version} )
        c.String(200, version)
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
		util.Elastic( util.Document{API: "/AccountsGenerate", Status: "start"} )
		result, _ := actions.AccountsGenerateAction(c)

		
		jsonData, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
	
		// Convertendo os dados JSON para string
		jsonString := string(jsonData)

		util.Elastic( util.Document{API: "/AccountsGenerate", Status: "end", Info: jsonString } )
		c.JSON(http.StatusOK, result)
    })


    r.Run() // por padrão na porta 8080
}