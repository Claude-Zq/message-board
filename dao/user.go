package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"message-board/model"
)

func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE user SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("服务器出错")
			return
		}
	}()

	row := dB.QueryRow("SELECT id,password FROM user WHERE username = ?", username)
	if row.Err() != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("INSERT INTO user(username,password) values(?,?)", user.Username, user.Password)
	return err
}

var ErrNoQuestion = errors.New("dao: This user has no question")

func GetQuestionByUsername(username string) (string, error) {
	var questionNull sql.NullString
	row := dB.QueryRow("SELECT question FROM user WHERE username = ?", username)
	if row.Err() != nil {
		return questionNull.String, row.Err()
	}
	err := row.Scan(&questionNull)
	if err != nil {
		return questionNull.String, err
	}
	//密保问题不为空
	if questionNull.Valid {
		return questionNull.String, nil
	} else {
		return questionNull.String, ErrNoQuestion
	}
}

var ErrNoAnswer = errors.New("dao: This question has no answer")

func GetAnswerByUsername(username string) (string, error) {
	var answer string
	row := dB.QueryRow("SELECT COALESCE(answer,'') FROM user WHERE username = ?", username)
	if row.Err() != nil {
		return "", row.Err()
	}
	err := row.Scan(&answer)
	if err != nil {
		return "", err
	}
	if answer == "" {
		return "", ErrNoAnswer
	} else {
		return answer, nil
	}

}
