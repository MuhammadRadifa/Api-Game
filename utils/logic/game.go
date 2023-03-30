package logic

import (
	"database/sql"
	"game-api/utils/structs"
)

func GetAllGame(db *sql.DB, param int) (result []structs.Game, err error) {
	sqlStatement := "SELECT * FROM Game JOIN Category ON game.category_id = category.id WHERE game.id=($1);"

	rows, err := db.Query(sqlStatement, param)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var Game = structs.Game{}
		var Category = structs.Category{}

		err = rows.Scan(&Game.Id, &Game.Name, &Game.Description, &Game.Category_id, &Category.Id, &Category.Name)

		if err != nil {
			panic(err)
		}

		Game.Category = append(Game.Category, Category)

		result = append(result, Game)
	}

	return
}

func InsertGame(db *sql.DB, Game structs.Game) (err error) {
	sqlStatement := "INSERT INTO Game (name,description,category_id) VALUES ($1,$2,$3);"

	errs := db.QueryRow(sqlStatement, Game.Name, Game.Description, Game.Category_id)

	return errs.Err()
}

func UpdateGame(db *sql.DB, Game structs.Game) (err error) {
	sqlStatement := "UPDATE Game SET name=$1, description=$2 ,category_id=$3 WHERE id=$4;"

	errs := db.QueryRow(sqlStatement, Game.Name, Game.Description, Game.Category_id, Game.Id)

	return errs.Err()
}

func DeletedGame(db *sql.DB, Game structs.Game) (err error) {
	sqlStatement := "DELETE FROM Game WHERE id=$1;"

	errs := db.QueryRow(sqlStatement, Game.Id)

	return errs.Err()

}
