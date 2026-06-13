// Track: Go Backend | Step 2 of 10
package account

import "testing"

// LEARNING STEP 2: Table-driven account tests
// Goal: Test deposit/withdraw validation with table-driven tests.
// Done when: Tests cover zero, negative, valid deposit, valid withdraw, and overdraft.
// Verify: cd go && go test ./internal/account -cover
// Reinforces: Go testing idioms, table-driven tests
//
// Tasks:
//   1. Test Deposit rejects amount <= 0; accepts valid amounts.
//   2. Test Withdraw rejects overdraft and amount <= 0.
//   3. Add compile-time check: var _ BankAccountInterface = (*BankAccountImpl)(nil)
//   4. After refactor: tests should target the single merged type.
//
// Example table-driven shape:
//   tests := []struct {
//       name           string
//       startBalance   float64
//       amount         float64
//       op             string // "deposit" or "withdraw"
//       wantOK         bool
//       wantBalance    float64
//   }{ ... }


// Complete Task # 1: Test Deposit rejects amount <= 0; accepts valid amounts.
func TestDeposit(t *testing.T) {
	tests := []struct {
		name        string
		amount      float64
		wantOK      bool
		wantBalance float64
	}{
		{name: "Zero deposit", amount: 0.00, wantOK: false, wantBalance: 0.00},
		{name: "Negative deposit", amount: -100.00, wantOK: false, wantBalance: 0.00},
		{name: "Positive deposit", amount: 100.00, wantOK: true, wantBalance: 100.00},
		{name: "Large deposit", amount: 1000000.00, wantOK: true, wantBalance: 1000000.00},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			account := &BankAccount{balance: 0.00}
			gotOK := account.Deposit(test.amount)
			if gotOK != test.wantOK {
				t.Errorf("Deposit(%f) = %v, want %v", test.amount, gotOK, test.wantOK)
			}
			if account.GetBalance() != test.wantBalance {
				t.Errorf("GetBalance() = %f, want %f", account.GetBalance(), test.wantBalance)
			}
		})
	}
}

// Complete Task # 2: Test Withdraw rejects overdraft and amount <= 0.
func TestWithdraw(t *testing.T) {
	tests := []struct{
		name string
		startBalance float64
		amount float64
		wantOK bool
		wantBalance float64
	}{
		// CHANGE: each row needs startBalance; balance after failed ops must stay unchanged
        {name: "Zero withdraw", startBalance: 100, amount: 0, wantOK: false, wantBalance: 100},
        {name: "Negative withdraw", startBalance: 100, amount: -100, wantOK: false, wantBalance: 100},
        {name: "Overdraft", startBalance: 100, amount: 150, wantOK: false, wantBalance: 100},
        {name: "Valid withdraw", startBalance: 100, amount: 50, wantOK: true, wantBalance: 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			account := &BankAccount{balance: test.startBalance}
			gotOK := account.Withdraw(test.amount)
			if gotOK != test.wantOK {
				t.Errorf("Withdraw(%f) = %v, want %v", test.amount, gotOK, test.wantOK)
			}
			if account.GetBalance() != test.wantBalance {
				t.Errorf("GetBalance() = %f, want %f", account.GetBalance(), test.wantBalance)
			}
		})
		
	}
}


var _ BankAccountInterface = (*BankAccount)(nil)
