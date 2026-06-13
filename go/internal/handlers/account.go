// Track: Go Backend | Step 6 of 10
package handlers

// LEARNING STEP 6: Fiber account handlers
// Goal: HTTP layer that calls BankService and returns JSON.
// Done when: Handlers map service errors to HTTP status codes correctly.
// Verify: cd go && go run ./cmd/api (then curl or Postman)
// Reinforces: Fiber handlers, JSON request/response (Applied Math Mastery)
//
// Tasks:
//   1. type AccountHandler struct { svc *service.BankService }
//   2. Register routes on a Fiber app (or return a fiber.Router group):
//        POST /accounts
//        GET  /accounts/:id/balance
//        POST /accounts/:id/deposit
//        POST /accounts/:id/withdraw
//   3. Request/response structs with json tags, e.g.:
//        type CreateAccountRequest struct { InitialBalance float64 `json:"initial_balance"` }
//        type BalanceResponse struct { ID uint `json:"id"`; Balance float64 `json:"balance"` }
//   4. Error mapping:
//        ErrInvalidAmount  -> 400 Bad Request
//        ErrAccountNotFound -> 404 Not Found
//        other             -> 500 Internal Server Error
//
// LEARNING STEP 10 (polish): Structured logging and env config
// Goal: Load PORT, JWT_SECRET, DB_PATH from environment; use log/slog or Fiber logger.
// Done when: Config is documented in README; no hardcoded secrets in production paths.
// Verify: cd go && go run ./cmd/api with env vars set
//
// LEARNING STEP 10 (bonus): Table-driven httptest for handlers — see account_test note in this package.
