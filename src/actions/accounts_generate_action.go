package actions

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"context"

	"github.com/gin-gonic/gin"
)


type AccountsActionResult struct {
	Result			Result				`json:"result"`
	AccountsCount	int					`json:"accountsCount"`
	Accounts		[]AccountsResult	`json:"accounts"`
}

type AccountsResult struct {
	Id		int			`json:"id"`
	Nome	string		`json:"nome"`
	Email	string		`json:"email"`
}


func AccountsGenerateAction(c *gin.Context) AccountsActionResult {

	result := AccountsActionResult{ Result: Result{StatusCode: "200", Messages: make([]string,0), IsOK: true} } 

	ctx := context.Background()
	var con = db.InitDB()

	// 
	generateAccounts(10, con)

	// 
	queries := dbsqlc.New(con)
	
	// accounts, err := queries.GetAllLogs(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	
	con.Close()

	// result.LogCont = len(logs)
	// result.Logs = convertDBLogToLog(logs)
	return result

}


func generateAccounts(count int) {

}





// func convertDBLogToLog(logs []dbsqlc.Log) []LogResult {
// 	logResults := make([]LogResult, 0, len(logs))

// 	for _, log := range logs {
// 		logResult := LogResult{Id: int(log.ID), Mensagem: log.Mensagem.String}
// 		logResults = append(logResults, logResult)
// 	}

// 	return logResults
// }
