package main

import (
	"GoGinAPI/actions"
	"GoGinAPI/db"
	"GoGinAPI/util"
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
		util.Elastic( util.Document{API: "/ContratoEquipamentos", Status: "start"} )

		var con = db.InitDB()
		registros , err := db.QueryNotificacoes(con)
		jsonContent, _ := util.Convert(registros)

		if err != nil {
			fmt.Println(err)
		}
		con.Close()

		util.Elastic( util.Document{API: "/ContratoEquipamentos", Status: "end", Info: jsonContent} )
		c.JSON(http.StatusOK, registros)
    })

	
	r.GET("/Logs", func(c *gin.Context) {
		c.JSON(http.StatusOK, actions.LogsAction(c))
    })


	r.GET("/AccountsGenerate", func(c *gin.Context) {
		util.Elastic( util.Document{API: "/AccountsGenerate", Status: "start"} )
		result, _ := actions.AccountsGenerateAction(c)
		jsonContent, _ := util.Convert(result)
		util.Elastic( util.Document{API: "/AccountsGenerate", Status: "end", Info: jsonContent } )
		c.JSON(http.StatusOK, result)
    })


    r.Run() // por padrão na porta 8080
}