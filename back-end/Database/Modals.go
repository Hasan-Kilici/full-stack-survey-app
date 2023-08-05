package Database

type User struct{
	Token		string
	Username	string
	Email		string
	Password	string
	Perm		string
}

type Surveys struct{
	Token			string
	Title			string
	Description		string
}

type Options struct{
	Token			string
	Title			string
	Votes			int 
	SurveyToken		string
}

type Votes struct{
	Token			string
	OptionToken		string
	UserToken		string
	SurveyToken		string
}