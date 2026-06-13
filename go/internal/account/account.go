// Track: Go Backend | Steps 1–2 of 10
//
// LEARNING STEP 1: Account domain (complete — migrated from root bank.go)
// Goal: Encapsulated balance, validation, and BankAccountInterface.
// Done when: CLI demo runs and prints expected balances.
// Verify: cd go && go run ./cmd/cli
// Reinforces: structs, pointer receivers, interfaces, encapsulation
package account

import (
	"fmt"
	"math/rand"
)

type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(amount float64) bool {
	return b.addFunds(amount)
}

func (b *BankAccount) Withdraw(amount float64) bool {
	return b.subtractFunds(amount)
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) validateWithdraw(amount float64) bool {
	if amount <= 0 {
		return false
	}
	if amount > b.balance {
		return false
	}
	return true
}

func (b *BankAccount) validateDeposit(amount float64) bool {
	if amount <= 0 {
		return false
	}
	return true
}

func (b *BankAccount) addFunds(amount float64) bool {
	if !b.validateDeposit(amount) {
		return false
	}
	b.balance += amount
	return true
}

func (b *BankAccount) subtractFunds(amount float64) bool {
	if !b.validateWithdraw(amount) {
		return false
	}
	b.balance -= amount
	return true
}

// LEARNING STEP 2 (refactor): Merge BankAccount and BankAccountImpl
// Goal: One concrete type implements BankAccountInterface; remove duplicate methods.
// Done when: All tests pass and only one struct remains (keep the interface).
// Verify: cd go && go test ./internal/account -cover
// Reinforces: DRY, interface satisfaction, compile-time checks
// Hint: var _ BankAccountInterface = (*BankAccountImpl)(nil)

type BankAccountInterface interface {		
	Deposit(amount float64) bool
	Withdraw(amount float64) bool
	GetBalance() float64
}

func ProcessAccount(account BankAccountInterface) {
	account.Deposit(500.00)
	account.Withdraw(200.00)
	fmt.Printf("Balance: %.2f\n", account.GetBalance())
}

// RunDemo runs the original bank.go main() flow for CLI verification.
func RunDemo() {
	account := BankAccount{balance: 1000.00}
	amount := 0.00

	if account.Deposit(500.00) {
		amount = 500.00
		fmt.Printf("Amount deposited: %.2f\n", amount)
		fmt.Printf("Balance after deposit: %.2f\n", account.GetBalance())
	}

	if account.Withdraw(200.00) {
		amount = 200.00
		fmt.Printf("Amount withdrawn: %.2f\n", amount)
		fmt.Printf("Balance after withdraw: %.2f\n", account.GetBalance())
	}

	if account.subtractFunds(200.00) {
		amount = 200.00
		fmt.Printf("Amount subtracted: %.2f\n", amount)
		fmt.Printf("Balance after subtractFunds: %.2f\n", account.GetBalance())
	}

	if account.addFunds(500.00) {
		amount = 500.00
		fmt.Printf("Amount added: %.2f\n", amount)
		fmt.Printf("Balance after addFunds: %.2f\n", account.GetBalance())
	}

	if !account.Withdraw(9999.00) {
		fmt.Println("Rejected overdraft withdraw (expected)")
	}
	fmt.Printf("Final balance: %.2f\n", account.GetBalance())

	accountImpl := &BankAccount{balance: float64(rand.Intn(500000))}
	fmt.Printf("Starting Balance: %.2f\n", accountImpl.GetBalance())
	ProcessAccount(accountImpl)
	fmt.Printf("Balance after processAccount: %.2f\n", accountImpl.GetBalance())
	fmt.Printf("AccountImpl: %+v\n", accountImpl)
}
