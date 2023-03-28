package logic

import (
	"database/sql"
	"errors"
	"game-api/utils/structs"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isRegister(db *sql.DB, email string) bool {
	sqlStatement := "SELECT * FROM Users WHERE email=$1"

	res, errs := db.Exec(sqlStatement, email)

	if errs != nil {
		panic(errs)
	}

	count, errs := res.RowsAffected()

	if errs != nil {
		panic(errs)
	}

	if count > 0 {
		return true
	}
	return false

}

func Register(db *sql.DB, Users structs.Users) (err error) {

	if isRegister(db, Users.Email) {
		return errors.New("email sudah terdaftar")
	}

	sqlStatement := "INSERT INTO Users (id,name,email,password) VALUES ($1,$2,$3,$4);"

	Users.Id = uuid.New().String()
	hashPass, _ := HashPassword(Users.Password)

	errs := db.QueryRow(sqlStatement, Users.Id, Users.Name, Users.Email, hashPass)

	return errs.Err()
}

// func Login(db *sql.DB, Users structs.Users) (err error) {

// 	return errs.Err()
// }
