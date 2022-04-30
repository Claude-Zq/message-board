package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

//按用户名修改密码
func ChangePassword(username, newPassword string) error {
	err := dao.UpdatePassword(username, newPassword)
	return err

}

//判断密码是否与数据库中的密码相符
func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.Password != password {
		return false, nil
	}
	return true, nil
}

func Register(user model.User) error {
	err := dao.InsertUser(user)
	return err
}

func IsRepeatUsername(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetQuestionByName(username string) (string, error) {
	return dao.GetQuestionByUsername(username)
}
