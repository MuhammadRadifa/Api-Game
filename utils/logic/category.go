package logic

import (
	"database/sql"
	"game-api/utils/structs"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
	sqlStatement := "SELECT * FROM Category;"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.Id, &category.Name)

		if err != nil {
			panic(err)
		}

		result = append(result, category)
	}

	return
}

func InsertCategory(db *sql.DB, Category structs.Category) (err error) {
	sqlStatement := "INSERT INTO Category (name) VALUES ($1);"

	errs := db.QueryRow(sqlStatement, Category.Name)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, Category structs.Category) (err error) {
	sqlStatement := "UPDATE Category SET name=$1 WHERE id=$2;"

	errs := db.QueryRow(sqlStatement, Category.Name, Category.Id)

	return errs.Err()
}

func DeletedCategory(db *sql.DB, Category structs.Category) (err error) {
	sqlStatement := "DELETE FROM Category WHERE id=$1;"

	errs := db.QueryRow(sqlStatement, Category.Id)

	return errs.Err()

}
