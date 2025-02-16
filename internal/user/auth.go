package user

import (
	"fmt"
	"merch-shop/internal/utils"
)

func (repo *UserDBRepostitory) GetAuthenticated(usr User) (*User, int, error) {
	exists := utils.CheckUser(repo.dtb, usr.Password)

	if exists {
		passwordHash, err := GetPasswordHash(repo.dtb, usr.Username)
		if err != nil {
			return nil, 500, err
		}
		check, err := CheckPassword(usr.Password, passwordHash)
		if err != nil {
			return nil, 500, err
		}
		if !check {
			return nil, 401, fmt.Errorf("password is incorrect")
		}

		thisUser := User{Username: usr.Username, Password: passwordHash}
		return &thisUser, 200, nil
	}

	thisUser, err := CreateUser(repo.dtb, usr)
	if err != nil {
		return nil, 500, err
	}

	return thisUser, 200, nil
}
