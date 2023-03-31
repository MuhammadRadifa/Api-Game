package databases

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func Connection() {
	err = godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("failed Load file environment")
	}

	// psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	err = db.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Ready To Connection ...")
	}

	DbMigrate(db)
}
