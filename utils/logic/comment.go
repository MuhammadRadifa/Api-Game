package logic

import (
	"database/sql"
	"game-api/utils/structs"
)

func InsertComment(db *sql.DB, Comment structs.Comment) (err error) {
	sqlStatement := "INSERT INTO Comment (text,users_id,game_id) VALUES ($1,$2,$3);"

	errs := db.QueryRow(sqlStatement, Comment.Text, Comment.Users_id, Comment.Game_id)

	return errs.Err()
}

func UpdateComment(db *sql.DB, Comment structs.Comment) (err error) {
	sqlStatement := "UPDATE Comment SET text=$1 WHERE id=$2;"

	errs := db.QueryRow(sqlStatement, Comment.Text, Comment.Id)

	return errs.Err()
}

func DeletedComment(db *sql.DB, Comment structs.Comment) (err error) {
	sqlStatement := "DELETE FROM Comment WHERE id=$1;"

	errs := db.QueryRow(sqlStatement, Comment.Id)

	return errs.Err()

}
