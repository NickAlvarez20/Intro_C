// Track: Go Backend | Step 1 of 10
//
// LEARNING STEP 1: Run the migrated bank demo
// Goal: Confirm account package works from a separate cmd/ entry point.
// Done when: Output matches your original bank.go run.
// Verify: cd go && go run ./cmd/cli
package main

import "henloworld/internal/account"

func main() {
	account.RunDemo()
}
