// Track: Go Backend | Step 8 of 10
package middleware

// LEARNING STEP 8: JWT authentication middleware
// Goal: Login endpoint issues JWT; protected routes require Bearer token.
// Done when: /accounts/* returns 401 without token; 200 with valid token.
// Verify: cd go && go run ./cmd/api — test with Postman Authorization tab
// Reinforces: JWT middleware (EndPointDB pattern)
//
// Tasks:
//   1. go get github.com/golang-jwt/jwt/v5
//   2. Hardcoded credentials for learning: admin / secret (move to env in Step 10).
//   3. LoginHandler: POST body { "username": "admin", "password": "secret" }
//      returns { "token": "<jwt>" } on success, 401 on failure.
//   4. JWTMiddleware: read Authorization: Bearer <token>, validate, call c.Next().
//   5. Sign claims with JWT_SECRET env var (default "dev-secret" for local only).
//
// Wire in cmd/api/main.go:
//   app.Post("/auth/login", LoginHandler)
//   app.Get("/health", ...)           // public
//   accounts := app.Group("/accounts", JWTMiddleware)
