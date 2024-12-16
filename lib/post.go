package lib

import (
	"database/sql"
	_ "database/sql/driver"
	"fmt"

	mod "FirstApp/model"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

func Con() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	/*	db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err)
		}
	*/db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func ConClose(db *sql.DB) {
	db.Close()

}
func InsDb(db *sql.DB, req string, client mod.CLient) sql.Result {
	deleteStmt := `INSERT INTO public.users (
		ph, chat_id, fio)
		VALUES ($1, $2, $3)`
	deleteStmt = req
	ans, _ := db.Exec(deleteStmt, client.Ph, client.Chat, client.Fio)

	fmt.Println("Successfully connected!")
	return ans
}
