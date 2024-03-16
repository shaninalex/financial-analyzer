package tests

import (
	"testing"

	"github.com/shaninalex/financial-analyzer/internal/db"
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
	t.Error("not implemented")
}

func TestUpdateReport(t *testing.T) {
	db := getDB()
	defer clear(db)
	t.Error("not implemented")
}

func TestCreateIssue(t *testing.T) {
	db := getDB()
	defer clear(db)
	t.Error("not implemented")
}

func TestUpdateIssue(t *testing.T) {
	db := getDB()
	defer clear(db)
	t.Error("not implemented")
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
