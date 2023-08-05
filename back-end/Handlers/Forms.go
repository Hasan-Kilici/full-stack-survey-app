package Handlers

type RegisterForm struct {
	Username 	string 	`json:"username"`
	Email 		string 	`json:"email"`
	Pass 		string 	`json:"pass"`
}

type LoginForm struct {
	Email 	string 	`json:"email"`
	Pass 	string 	`json:"pass"`
}

type CreateSurveyForm struct {
	Title			string		`json:"title"`
	Description		string 		`json:"description"`
	Options			string		`json:"options"`
}
