package db

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type IDatabaseRepository interface {
	CreateRequest(userId, ticker string)
}

type Database struct {
	DB *sql.DB
}

func InitDatabase(dsn string) (IDatabaseRepository, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: conn,
	}, nil
}

func (d *Database) CreateRequest(userId, ticker string) {
	newRequestID := uuid.NewString()
	log.Println(newRequestID)
	sql := `
		INSERT INTO 
			requests (id, user_id, tocker)
		VALUES
			($1, $2, $3);
	`
	_, err := d.DB.Exec(sql, newRequestID, userId, ticker)
	if err != nil {
		log.Println(err)
	}

}
