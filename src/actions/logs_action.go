package actions

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)


type Result struct {
	StatusCode 	string
	Message		string
	Result		[]dbsqlc.Log
}


func LogsAction(c *gin.Context) Result {

	result := Result{StatusCode: "", Message: ""}

	ctx := context.Background()
	var con = db.InitDB()

	queries := dbsqlc.New(con)
	
	logs, err := queries.GetAllLogs(ctx)
	if err != nil {
		fmt.Println(err)
	}
	
	con.Close()

	result.Result = logs
	return result

}