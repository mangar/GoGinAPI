package actions

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"GoGinAPI/sqlc"
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
	Logs		[]LogResult		`json:"logs"`
}

type LogResult struct {
	Id			int		`json:"id"`
	Mensagem 	string 	`json:"mensagem"`
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
	result.Logs = convertDBLogToLog(logs)
	return result

}


func convertDBLogToLog(logs []sqlc.Log) []LogResult {
	logResults := make([]LogResult, 0, len(logs))

	for _, log := range logs {
		logResult := LogResult{Id: log.ID, Mensagem: log.Mensagem}
		logResults = append(logResults, logResult)
	}

	return logResults
}
