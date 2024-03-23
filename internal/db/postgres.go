package db

import (
	"context"
	"fmt"
	"log"

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

func (db *PSQLDatabase) ReportGet(reportId uint) (*typedefs.Report, error) {
	report := &typedefs.Report{}
	db.DB.Preload("Issues").First(report, reportId)
	if db.DB.Error != nil {
		return nil, db.DB.Error
	}
	return report, nil
}

func (db *PSQLDatabase) ReportCreate(report *typedefs.Report) error {
	db.DB.Create(report)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}

func (db *PSQLDatabase) ReportUpdate(reportId uint, report map[string]interface{}) error {
	db.DB.Model(typedefs.Report{}).Where("id = ?", reportId).Updates(report)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}

func (db *PSQLDatabase) ReportDataCreate(data *typedefs.ReportData) error {
	db.DB.Create(data)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}

func (db *PSQLDatabase) CheckReportStatus(reportId uint, dataType string) {
	var dTypes []string
	if err := db.DB.Model(&typedefs.ReportData{}).Where("report_id = ?", reportId).Pluck("type", &dTypes).Error; err != nil {
		log.Println("Error querying database:", err)
		return
	}

	required := map[string]bool{
		"summary":    true,
		"financials": true,
		"dividend":   true,
		"price":      true,
		"keyratios":  true,
	}

	for _, dType := range dTypes {
		delete(required, dType)
	}

	if len(required) > 0 {
		return
	}

	if err := db.ReportUpdate(reportId, map[string]interface{}{
		"status": true,
	}); err != nil {
		log.Println("Error updating report status:", err)
	}
}

func (db *PSQLDatabase) IssueGet(issueId uint) (*typedefs.Issue, error) {
	issue := &typedefs.Issue{}
	db.DB.First(issue, issueId)
	if db.DB.Error != nil {
		return nil, db.DB.Error
	}
	return issue, nil
}

func (db *PSQLDatabase) IssueCreate(issue *typedefs.Issue) error {
	db.DB.Create(issue)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}

func (db *PSQLDatabase) IssueUpdate(issueId uint, issue map[string]interface{}) error {
	db.DB.Model(typedefs.Issue{}).Where("id = ?", issueId).Updates(issue)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}

func (db *PSQLDatabase) IssueDelete(issueId uint) error {
	db.DB.Delete(typedefs.Issue{}, issueId)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}

func (db *PSQLDatabase) Raw(query string) error {
	db.DB.Exec(query)
	if db.DB.Error != nil {
		return db.DB.Error
	}
	return nil
}
