package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shaninalex/financial-analyzer/internal/db"
	"github.com/shaninalex/financial-analyzer/internal/typedefs"
)

const (
	DatabaseDsn           = "host=localhost user=user password=password dbname=fin port=5432 sslmode=disable"
	DatabaseType          = "psql"
	DatabaseDefaultSchema = "test"
)

func TestDatabaseInitialization(t *testing.T) {
	db, err := db.InitDatabase(DatabaseDsn, DatabaseDefaultSchema, DatabaseType)
	if err != nil {
		t.Errorf("Unable to initialize db. Error: %v", err)
	}
	if db == nil {
		t.Error("DB should not be nil.")
	}
}

func TestCreateReport(t *testing.T) {
	db := getDB()
	defer clear(db)
	report := &typedefs.Report{
		UserId: uuid.NewString(),
		Ticker: "AAPL",
		Link:   "https://report.pdf",
		Status: true,
	}
	err := db.ReportCreate(report)
	if err != nil {
		t.Error(err)
	}

	if report.ID < 0 {
		t.Errorf("report id should be greater than 0. Got %d", report.ID)
	}
}

func TestReportGet(t *testing.T) {
	db := getDB()
	defer clear(db)

	report := &typedefs.Report{
		UserId: uuid.NewString(),
		Ticker: "AAPL",
		Link:   "https://report.pdf",
		Status: true,
	}
	db.ReportCreate(report)

	r, err := db.ReportGet(report.ID)
	if err != nil {
		t.Errorf("Unable to get report. Err: %v", err)
	}
	if r.UserId != report.UserId {
		t.Errorf("Wrong report. Want user id: %s, Got: %s", report.UserId, r.UserId)
	}
	if r.Ticker != report.Ticker {
		t.Errorf("Want ticker: %s, Got: %s", report.Ticker, r.Ticker)
	}
	if r.Link != report.Link {
		t.Errorf("Want link: %s, Got: %s", report.Link, r.Link)
	}
	if r.UserId != report.UserId {
		t.Errorf("Want status: %t, Got: %t", report.Status, r.Status)
	}
}

func TestUpdateReport(t *testing.T) {
	db := getDB()
	defer clear(db)

	report := &typedefs.Report{
		UserId: uuid.NewString(),
		Ticker: "AAPL",
		Link:   "https://report.pdf",
		Status: true,
	}
	db.ReportCreate(report)
	payload := map[string]interface{}{
		"link":   "http://new_link.pdf",
		"ticker": "TEST",
		"status": false,
	}
	err := db.ReportUpdate(report.ID, payload)
	if err != nil {
		t.Errorf("Unable to get report. Err: %v", err)
	}

	r, _ := db.ReportGet(report.ID)
	if r.Link != payload["link"].(string) {
		t.Errorf("Wrong report. Want link: %s, Got: %s", payload["link"].(string), r.Link)
	}
	if r.Ticker != payload["ticker"].(string) {
		t.Errorf("Wrong report. Want ticker: %s, Got: %s", payload["ticker"].(string), r.Ticker)
	}
	if r.Status != payload["status"].(bool) {
		t.Errorf("Wrong report. Want status: %t, Got: %t", payload["status"].(bool), r.Status)
	}
}

func TestCreateIssue(t *testing.T) {
	db := getDB()
	defer clear(db)
	report := &typedefs.Report{
		UserId: uuid.NewString(),
		Ticker: "AAPL",
		Link:   "https://report.pdf",
		Status: true,
	}
	db.ReportCreate(report)
	issue := &typedefs.Issue{
		Title:       "new issue",
		Description: "description",
		ReportID:    report.ID,
	}
	err := db.IssueCreate(issue)
	if err != nil {
		t.Errorf("unable to create issue. Err: %v", err)
	}

	if issue.ID < 0 {
		t.Errorf("Issue should have id more than 0. Got: %d", issue.ID)
	}

	dbReport, _ := db.ReportGet(report.ID)
	if len(dbReport.Issues) != 1 {
		t.Errorf("Target report does not contain created issue. Want 1, Got: %d", len(dbReport.Issues))
	}
}

func TestIssueGet(t *testing.T) {
	db := getDB()
	defer clear(db)
	report := &typedefs.Report{
		UserId: uuid.NewString(),
		Ticker: "AAPL",
		Link:   "https://report.pdf",
		Status: true,
	}
	db.ReportCreate(report)
	issue := &typedefs.Issue{
		Title:       "new issue",
		Description: "description",
		ReportID:    report.ID,
	}
	db.IssueCreate(issue)

	dbIssue, err := db.IssueGet(issue.ID)
	if err != nil {
		t.Errorf("unable to get issue. Err: %v", err)
	}

	if dbIssue.ID != issue.ID {
		t.Errorf("Wrong issue. Want id: %d, Got: %d", issue.ID, dbIssue.ID)
	}
	if dbIssue.Title != issue.Title {
		t.Errorf("Wrong issue. Want title: %s, Got: %s", issue.Title, dbIssue.Title)
	}
	if dbIssue.Description != issue.Description {
		t.Errorf("Wrong issue. Want description: %s, Got: %s", issue.Description, dbIssue.Description)
	}
	if dbIssue.ReportID != issue.ReportID {
		t.Errorf("Wrong issue. Want report id: %d, Got: %d", issue.ReportID, dbIssue.ReportID)
	}
}

func TestUpdateIssue(t *testing.T) {
	db := getDB()
	defer clear(db)
	report := &typedefs.Report{
		UserId: uuid.NewString(),
		Ticker: "AAPL",
		Link:   "https://report.pdf",
		Status: true,
	}
	db.ReportCreate(report)
	issue := &typedefs.Issue{
		Title:       "new issue",
		Description: "description",
		ReportID:    report.ID,
	}
	db.IssueCreate(issue)

	payload := map[string]interface{}{
		"title":       "changed title",
		"description": "changed description",
	}

	db.IssueUpdate(issue.ID, payload)
	dbIssue, _ := db.IssueGet(issue.ID)
	if dbIssue.Title != payload["title"].(string) {
		t.Errorf("Wrong issue. Want title: %s, Got: %s", payload["title"].(string), dbIssue.Title)
	}

	if dbIssue.Description != payload["description"].(string) {
		t.Errorf("Wrong issue. Want description: %s, Got: %s", payload["description"].(string), dbIssue.Description)
	}
}

func clear(d db.IDatabaseRepository) {
	d.Raw("DELETE FROM issues")
	d.Raw("DELETE FROM reports")
}

func getDB() db.IDatabaseRepository {
	d, _ := db.InitDatabase(DatabaseDsn, DatabaseDefaultSchema, DatabaseType)
	// hide gorm warnings and live only test messages
	// d.DB.Logger = logger.Default.LogMode(logger.Silent)
	return d
}
