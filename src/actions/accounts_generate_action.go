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
	Result        Result           `json:"result"`
	AccountsCount int              `json:"accountsCount"`
	Accounts      []AccountResult  `json:"accounts"`
}

type AccountResult struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func AccountsGenerateAction(c *gin.Context) (AccountsActionResult, error) {

	result := AccountsActionResult{Result: Result{StatusCode: "200", Messages: make([]string, 0), IsOK: true}}

	ctx := context.Background()
	var con = db.InitDB()

	
	queries := dbsqlc.New(con)

	//
	// 


	// generateAccounts(10, queries, &ctx)

	account := dbsqlc.Account{Name: sql.NullString{String:"aaaa"}, Email: sql.NullString{String:"bbbb"}}

	// _, err := queries.CreateAccount(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// }


	// 
	// 



	accounts, err := queries.GetAccounts(ctx)
	if err != nil {
		fmt.Println(err)
	}

	con.Close()

	result.AccountsCount = len(accounts)
	result.Accounts = convertDB(accounts)
	return result, err

}

// func generateAccounts(count int, queries *dbsqlc.Queries, ctx *context.Context) error {

// 	// account := dbsqlc.Account{
// 	// 	Name: sql.NullString{String:"aaaa"},
// 	// 	Email: sql.NullString{String:"bbbb"}}

// 	err := queries.CreateAccount(*ctx)

// 	// result, err := queries.InsertAccount(*ctx, account)

// 	return err
// }

func convertDB(accounts []dbsqlc.Account) []AccountResult {
	accountsResult := make([]AccountResult, 0, len(accounts))

	for _, r := range accounts {
		accResult := AccountResult{Id: int(r.ID), Nome: r.Name.String, Email: r.Email.String}
		accountsResult = append(accountsResult, accResult)
	}

	return accountsResult
}
