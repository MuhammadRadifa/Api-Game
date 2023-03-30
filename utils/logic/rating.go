package logic

import (
	"database/sql"
	"game-api/utils/structs"
)

func InsertRating(db *sql.DB, Rating structs.Rating) (err error) {
	sqlStatement := "INSERT INTO Rating (rate,users_id,game_id) VALUES ($1,$2,$3);"

	errs := db.QueryRow(sqlStatement, Rating.Rate, Rating.Users_id, Rating.Game_id)

	return errs.Err()
}

func UpdateRating(db *sql.DB, Rating structs.Rating) (err error) {
	sqlStatement := "UPDATE Rating SET text=$1 WHERE id=$2;"

	errs := db.QueryRow(sqlStatement, Rating.Rate, Rating.Id)

	return errs.Err()
}

func DeletedRating(db *sql.DB, Rating structs.Rating) (err error) {
	sqlStatement := "DELETE FROM Rating WHERE id=$1;"

	errs := db.QueryRow(sqlStatement, Rating.Id)

	return errs.Err()

}
