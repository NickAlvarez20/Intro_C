// Track: Go Backend | Step 5 of 10
package repository

// LEARNING STEP 5: AccountRepository interface
// Goal: Abstract persistence so BankService does not know about GORM.
// Done when: Interface is defined and sqlite.go implements it.
// Verify: cd go && go test ./internal/repository -v
// Reinforces: repository pattern, interface-based design (second spiral after account interface)
//
// Tasks:
//   1. Define AccountRepository interface:
//        Create(initialBalance float64) (uint, error)
//        GetByID(id uint) (balance float64, err error)
//        UpdateBalance(id uint, newBalance float64) error
//   2. Define domain errors: ErrAccountNotFound, etc.
//
// LEARNING STEP 5 (bonus): In-memory fake repository for fast service tests.
// Goal: Implement MemoryRepository satisfying AccountRepository without SQLite.
// Done when: service tests run without touching the filesystem.
