package migrations

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Migration used to package each individual migration.
type Migration struct {
	DB       *gorm.DB
	Executor func(*gorm.DB) error
	Key      string
}

// MigrationResult used to wrap the underlying error.
type MigrationResult struct {
	Error error
}

// Execute used to run each migration.
func (m *Migration) Execute() MigrationResult {
	var err error

	// Start transaction
	tx := m.DB.Begin()
	if tx.Error != nil {
		return MigrationResult{Error: err}
	}

	// Run migration logic
	err = m.Executor(tx)
	if err != nil {
		tx.Rollback()
		return MigrationResult{Error: err}
	}

	// Commit transaction
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return MigrationResult{Error: err}
	}

	// Return result
	return MigrationResult{Error: err}
}

// MigrateUp called from the postgres setup flow.
func MigrateUp(db *gorm.DB) error {
	// Ensure schema
	ensurePublicSchema(db)

	// Ensure migrations table exists
	ensureMigrationsTable(db)

	// Run migrations
	for _, m := range migrationsList(db) {
		// Skip if migration has been ran
		if migrationHasRan(db, m.Key) {
			fmt.Println("Skipping migration that has already ran: ", m.Key)
			continue
		}

		if result := m.Execute(); result.Error != nil {
			panic(result.Error)
		}

		// There was no error, so create a record for the migration
		createMigrationRecord(db, m.Key)
	}

	return nil
}

func migrationsList(db *gorm.DB) []Migration {
	// List migrations in order
	m := []Migration{
		Migration{DB: db, Executor: CreateArticlesTable, Key: "20190324_create_articles"},
	}

	return m
}

func ensurePublicSchema(db *gorm.DB) {
	err := db.Exec("CREATE SCHEMA IF NOT EXISTS public;").Error
	if err != nil {
		panic(fmt.Sprintf("Error creating public schema. Cannot continue: %s", err))
	}
}

func ensureMigrationsTable(db *gorm.DB) {
	err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			ran_at bigint,
			key text,
			CONSTRAINT migrations_key UNIQUE (key)
		)
																	`).Error
	if err != nil {
		panic(fmt.Sprintf("Error creating migrations table. Cannot continue: %s", err))
	}
}

func migrationHasRan(db *gorm.DB, key string) bool {
	var count int
	err := db.Table("migrations").Where("key = ?", key).Count(&count).Error
	if err != nil {
		panic(fmt.Sprintf("Error checking for executed migration. Cannot continue: %s", err))
	}

	return count > 0
}

func createMigrationRecord(db *gorm.DB, key string) {
	err := db.Exec(`INSERT INTO migrations (key, ran_at) VALUES (?, ?)`, key, time.Now().Unix()).Error
	if err != nil {
		panic(fmt.Sprintf("Error creating migration. Cannot continue: %s", err))
	}
}
