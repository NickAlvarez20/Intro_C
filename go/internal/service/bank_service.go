// Track: Go Backend | Step 3 of 10
package service

// LEARNING STEP 3: In-memory BankService
// Goal: Manage multiple accounts by ID; delegate to account.BankAccountInterface.
// Done when: CreateAccount, GetBalance, Deposit, Withdraw work via service tests.
// Verify: cd go && go test ./internal/service -v
// Reinforces: service layer pattern (Matrix Library / EndPointDB)
// Hint: map[string]account.BankAccountInterface
//
// Tasks:
//   1. Define BankService struct with an in-memory account map.
//   2. Implement:
//        CreateAccount(id string, initialBalance float64) error
//        GetBalance(id string) (float64, error)
//        Deposit(id string, amount float64) error
//        Withdraw(id string, amount float64) error
//   3. Return typed errors (not just bool):
//        var ErrAccountNotFound = errors.New("account not found")
//        var ErrInvalidAmount = errors.New("invalid amount")
//   4. Reject duplicate CreateAccount IDs.
//
// LEARNING STEP 5 (refactor): Replace the in-memory map with repository.AccountRepository.
// Done when: BankService uses the repository interface; in-memory tests still pass via a fake repo.
// Verify: cd go && go test ./internal/repository ./internal/service -v
