package user_repositories

import (
	"crypto/md5"
	"delivery/Services/Auth/Domain/DTOs"
	shareddb "delivery/Services/Sahared/Infrastructure/DB"
	"errors"
	"fmt"
)

type AuthRepository struct {
}

func (obj *AuthRepository) Login(login string, password string) error {
	db := shareddb.NewDB()
	defer db.Close()
	dto := &DTOs.LoginDTO{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
	if dto.Failed() {
		return errors.New("invalid credential")
	}
	return nil
}

func (obj *AuthRepository) Register(
	userName string,
	fullName string,
	email string,
	password string,
	country string,
	city string,
	street string,
	house string,
	flat string) error {

	fmt.Println("fuck data: ", userName,
		fullName,
		email,
		password,
		country,
		city,
		street,
		house,
		flat)

	db := shareddb.NewDB()
	defer db.Close()

	hash := md5.Sum([]byte(password))
	hashedPass := fmt.Sprintf("%x", hash)

	statement, err := db.Prepare(`
				INSERT INTO users (user_name, full_name, email, password, country, city, street, house, flat) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(userName, fullName, email, hashedPass, country, city, street, house, flat)
	if err != nil {
		return errors.New("could not register a new account")
	}
	return nil
}
