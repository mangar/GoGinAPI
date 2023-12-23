package actions


type Result struct {
	StatusCode 	string		`json:"statusCode"`
	Messages	[]string	`json:"messages"`	
	IsOK		bool		`json:"isOK"`
}

