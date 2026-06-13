// Track: Go Backend | Step 5 of 10
package repository

// LEARNING STEP 5: GORM SQLite implementation
// Goal: Implement AccountRepository using GORM + models.Account.
// Done when: Create/GetByID/UpdateBalance pass integration tests against henloworld.db.
// Verify: cd go && go test ./internal/repository -v
// Reinforces: GORM CRUD (Matrix Library)
//
// Tasks:
//   1. type SQLiteRepository struct { db *gorm.DB }
//   2. NewSQLiteRepository(db *gorm.DB) *SQLiteRepository
//   3. Implement all AccountRepository methods.
//   4. Use models.OpenDB + models.Migrate in test setup (t.TempDir() for test DB).
