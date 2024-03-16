package db

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/shaninalex/financial-analyzer/internal/typedefs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabaseRepository interface {
	// CreateRequest(userId, ticker string)
}

type Database struct {
	ctx context.Context
	db  *gorm.DB
}

func InitDatabase(dsn, scheme string) (IDatabaseRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	tx.Exec(fmt.Sprintf("SET SEARCH_PATH TO %s", scheme))
	tx.Commit()

	db.AutoMigrate(
		&typedefs.Report{},
		&typedefs.Issue{},
	)

	return &Database{
		ctx: context.TODO(),
		db:  db,
	}, nil
}
