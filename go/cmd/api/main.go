// Track: Go Backend | Steps 6 & 8 of 10
package main

import "fmt"

// LEARNING STEP 6: Fiber REST API entry point
// Goal: Wire Fiber app, inject BankService, register account routes.
// Done when: All four account endpoints respond with correct JSON and status codes.
// Verify: cd go && go run ./cmd/api
// Reinforces: Fiber bootstrap (Applied Math Mastery)
// Hint: github.com/gofiber/fiber/v2 — go get github.com/gofiber/fiber/v2
//
// Endpoints to implement:
//   POST /accounts              body: { "initial_balance": 1000 }
//   GET  /accounts/:id/balance
//   POST /accounts/:id/deposit  body: { "amount": 500 }
//   POST /accounts/:id/withdraw body: { "amount": 200 }
//
// LEARNING STEP 6 (config): Read PORT from env, default "8080".
//
// LEARNING STEP 8 (auth wiring): After middleware/auth.go is done:
//   - Add POST /auth/login
//   - Protect /accounts/* with JWT middleware
//   - Keep GET /health public

func main() {
	fmt.Println("LEARNING STEP 6: Implement Fiber API in this file.")
	fmt.Println("Verify: cd go && go run ./cmd/api")
}
