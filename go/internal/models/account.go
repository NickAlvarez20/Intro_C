// Track: Go Backend | Step 4 of 10
package models

// LEARNING STEP 4: GORM account model + SQLite
// Goal: Persist accounts with GORM; AutoMigrate in a test or helper.
// Done when: Account model migrates and you can create/read one row in SQLite.
// Verify: cd go && go test ./internal/models -v
// Reinforces: GORM tags, AutoMigrate, SQLite (Matrix Library)
//
// Tasks:
//   1. go get gorm.io/gorm gorm.io/driver/sqlite
//   2. Define Account struct:
//        ID        uint      `gorm:"primaryKey"`
//        Balance   float64
//        CreatedAt time.Time
//        UpdatedAt time.Time
//   3. Write OpenDB(path string) (*gorm.DB, error) opening henloworld.db
//   4. Write Migrate(db *gorm.DB) error calling db.AutoMigrate(&Account{})
//   5. Optional test: create account with Balance 1000, read it back.
//
// Hint: DB file path default "henloworld.db" in go/ (gitignored).
