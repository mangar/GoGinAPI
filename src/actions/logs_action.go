package actions

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)


type Result struct {
	StatusCode 	string		`json:"statusCode"`
	Messages	[]string	`json:"messages"`	
	IsOK		bool		`json:"isOK"`
}

type LogActionResult struct {
	Result		Result			`json:"result"`
	LogCont		int				`json:"logCount"`
	Logs		[]dbsqlc.Log	`json:"logs"`
}





func LogsAction(c *gin.Context) LogActionResult {

	result := LogActionResult{ Result: Result{StatusCode: "200", Messages: make([]string,0), IsOK: true} } 

	ctx := context.Background()
	var con = db.InitDB()

	queries := dbsqlc.New(con)
	
	logs, err := queries.GetAllLogs(ctx)
	if err != nil {
		fmt.Println(err)
	}
	
	con.Close()

	result.LogCont = len(logs)
	result.Logs = logs
	return result

}