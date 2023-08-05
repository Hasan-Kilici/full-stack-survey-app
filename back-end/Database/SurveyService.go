package Database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
	"survey/Utils"
	"fmt"
)
//Insert
func CreateSurvey(Title,Description,Options string) error {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO surveys (Token, Title, Description) VALUES ($1, $2, $3)"
	_, err := db.Exec(query,Token,Title,Description)
	if err != nil {
		fmt.Println(err)
		return err
	}

	CreateOptions(Options, Token)
	return nil
}

func CreateOptions(Options,SurveyToken string) error {
	Option := strings.Split(Options, ",")
	OptionCount := len(Option)

	for i := 0;i < OptionCount;i++ {
		Token := Utils.GenerateToken(31)
		query := "INSERT INTO options (Token, Title, Votes, SurveyToken) VALUES ($1, $2, $3, $4)"
		_, err := db.Exec(query,Token,Option[i],0,SurveyToken)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func GiveVote(UserToken,SurveyToken,OptionToken string) error {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO votes (Token, OptionToken, UserToken, SurveyToken) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query,Token,OptionToken,UserToken,SurveyToken)
	if err != nil {
		fmt.Println(err)
		return err
	}

	UpdateVoteCount(OptionToken)
	return nil
}
//Update
func UpdateVoteCount(OptionToken string) error {
	VoteCount := GetVoteCount(OptionToken)
	query := "UPDATE Options SET votes=$1 WHERE Token=$2"
	_, err := db.Exec(query,VoteCount,OptionToken)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
//Delete
func DeleteSurvey(SurveyToken string) error {
	query := "DELETE FROM surveys WHERE Token=$1"
	_, err := db.Exec(query,SurveyToken)
	if err != nil {
		fmt.Println(err)
		return err
	}

	DeleteAllOptions(SurveyToken)
	DeleteAllVotes(SurveyToken)
	return nil
}

func DeleteOption(Token string) error {
	query := "DELETE FROM options WHERE Token=$1"
	_, err := db.Exec(query,Token)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil	
}

func DeleteAllOptions(SurveyToken string) error {
	query := "DELETE FROM options WHERE SurveyToken=$1"
	_, err := db.Exec(query,SurveyToken)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteAllVotes(SurveyToken string) error {
	query := "DELETE FROM votes WHERE SurveyToken = $1"
	_, err := db.Exec(query,SurveyToken)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
//Find
func GetVoteCount(OptionToken string) int {
	var rowCount int
	query := "SELECT COUNT(*) FROM votes WHERE OptionToken=$1"
	err := db.QueryRow(query, OptionToken).Scan(&rowCount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return rowCount
}

func FindVote(UserToken, SurveyToken string) bool {
	query := "SELECT COUNT(*) FROM votes WHERE UserToken = $1 AND SurveyToken $2"
    row := db.QueryRow(query, UserToken, SurveyToken)

    var count int
    err := row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }
	
    return true
}

func FindVoteCount(UserToken, SurveyToken string) int {
	var rowCount int
	query := "SELECT COUNT(*) FROM votes WHERE SurveyToken=$1 AND UserToken=$2"
	err := db.QueryRow(query, SurveyToken,UserToken).Scan(&rowCount)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return rowCount
}

func ListSurveys() ([]Surveys, error) {
	query := "SELECT * FROM surveys"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    surveys := []Surveys{}
    for rows.Next() {
        var survey Surveys
        err := rows.Scan(&survey.Token, &survey.Title, &survey.Description)
        if err != nil {
            return nil, err
        }
        surveys = append(surveys, survey)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return surveys, nil	
}

func FindSurvey(Token string) (Surveys, error){
    query := "SELECT * FROM surveys WHERE token = $1"
    row := db.QueryRow(query, Token)

    var survey Surveys
    err := row.Scan(&survey.Token, &survey.Title, &survey.Description)
    if err != nil {
        if err == sql.ErrNoRows {
            return Surveys{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return Surveys{}, err
    }
    return survey, nil
}

func FindOptions(SurveyToken string) ([]Options, error) {
	query := "SELECT * FROM options WHERE SurveyToken = $1"
    rows, err := db.Query(query, SurveyToken)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    options := []Options{}
    for rows.Next() {
        var option Options
        err := rows.Scan(&option.Token, &option.Title, &option.Votes, &option.SurveyToken)
        if err != nil {
            return nil, err
        }
        options = append(options, option)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return options, nil
}