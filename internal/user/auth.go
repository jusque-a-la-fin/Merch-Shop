package user

import (
	"fmt"
	"merch-shop/internal/utils"
)

func (repo *UserDBRepostitory) GetAuthenticated(usr User) (*User, bool, int, error) {
	exists, err := utils.CheckUser(repo.dtb, usr.Username)
	if err != nil {
		return nil, false, 500, err
	}

	if exists {
		passwordHash, err := GetPasswordHash(repo.dtb, usr.Username)
		if err != nil {
			return nil, false, 500, err
		}
		check, err := CheckPassword(usr.Password, passwordHash)
		if err != nil {
			return nil, false, 500, err
		}
		if !check {
			return nil, false, 401, fmt.Errorf("password is incorrect")
		}

		thisUser := User{Username: usr.Username, Password: passwordHash}
		return &thisUser, false, 200, nil
	}

	thisUser, err := CreateUser(repo.dtb, usr)
	if err != nil {
		return nil, false, 500, err
	}

	return thisUser, true, 200, nil
}
