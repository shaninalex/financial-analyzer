package db

import (
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/shaninalex/financial-analyzer/internal/typedefs"
)

type IDatabaseRepository interface {
	ReportGet(uint) (*typedefs.Report, error)
	ReportCreate(*typedefs.Report) error
	ReportUpdate(uint, map[string]interface{}) error
	ReportDataCreate(data *typedefs.ReportData) error
	CheckReportStatus(uint, string)
	IssueGet(uint) (*typedefs.Issue, error)
	IssueCreate(*typedefs.Issue) error
	IssueUpdate(uint, map[string]interface{}) error
	IssueDelete(uint) error

	// for test porpuses
	// Do not use in Production!
	Raw(string) error
}

func InitDatabase(dsn, scheme, dbType string) (IDatabaseRepository, error) {
	// can by multiple db implementation like Mysql or even MongoDB
	if dbType == "psql" {
		db, err := InitPSQL(dsn, scheme)
		if err != nil {
			return nil, err
		}

		tx := db.DB.Begin()
		tx.Exec(fmt.Sprintf("SET SEARCH_PATH TO %s", scheme))
		tx.Commit()

		db.DB.AutoMigrate(
			&typedefs.Report{},
			&typedefs.ReportData{},
			&typedefs.Issue{},
		)

		return db, nil
	}

	return nil, errors.New("db provider not implemented")

}
