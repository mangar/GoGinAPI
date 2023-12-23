package actions

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"context"
	"database/sql"
	"fmt"

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
	
	accounts, err := queries.GetAccounts(ctx)
	if err != nil {
		fmt.Println(err)
	}
	
	con.Close()

	result.AccountsCount = len(accounts)
	result.Accounts = convertDB(accounts)
	return result

}


func generateAccounts(count int, db *sql.DB) {
	
}





func convertDB(accounts []dbsqlc.Account) []AccountsResult {
	accountsResult := make([]AccountsResult, 0, len(accounts))

	for _, r := range accounts {
		accResult := AccountsResult{Id: int(r.ID), Nome: r.Name.String, Email: r.Email.String}
		accountsResult = append(accountsResult, accResult)
	}

	return accountsResult
}
