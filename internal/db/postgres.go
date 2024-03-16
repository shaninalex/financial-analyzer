package db

import (
	"context"
	"fmt"

	"github.com/shaninalex/financial-analyzer/internal/typedefs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PSQLDatabase struct {
	ctx context.Context
	DB  *gorm.DB
}

func InitPSQL(dsn, scheme string) (*PSQLDatabase, error) {
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

	return &PSQLDatabase{
		ctx: context.TODO(),
		DB:  db,
	}, nil
}

func (db *PSQLDatabase) ReportCreate(report *typedefs.Report) (*typedefs.Report, error) {
	return nil, nil
}

func (db *PSQLDatabase) ReportUpdate(report *typedefs.Report) (*typedefs.Report, error) {
	return nil, nil
}

func (db *PSQLDatabase) IssueCreate(issue *typedefs.Issue) (*typedefs.Issue, error) {
	return nil, nil
}

func (db *PSQLDatabase) IssueUpdate(issue *typedefs.Issue) (*typedefs.Issue, error) {
	return nil, nil
}

func (db *PSQLDatabase) IssueDelete(issue *typedefs.Issue) (*typedefs.Issue, error) {
	return nil, nil
}

func (db *PSQLDatabase) Raw(query string) error {
	return nil
}
