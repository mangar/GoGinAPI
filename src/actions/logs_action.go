package actions

import (
	"GoGinAPI/db"
	"GoGinAPI/dbsqlc"
	"context"
	"fmt"
)


func LogsAction() []Log {

	ctx := context.Background()
	var con = db.InitDB()

	queries := dbsqlc.New(con)
	logs, err := queries.GetAllLogs(ctx)
	if err != nil {
		fmt.Println(err)
	}
	
	con.Close()


	return nil

}