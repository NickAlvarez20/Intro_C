# Learning Path

Progress dashboard for this repo. Search the codebase for `LEARNING STEP` to jump to each task.

**Stack:** Go Fiber, GORM + SQLite, JWT, Docker, Postman — same tools as Applied Math Mastery, Matrix Library, and EndPointDB.

---

## Track A: Go Bank Backend (`go/`)

Work steps **in order**. Each step builds on the previous.

| Step | Topic | File(s) | Verify |
|------|-------|---------|--------|
| 1 | Account domain + CLI demo | [`go/internal/account/account.go`](go/internal/account/account.go), [`go/cmd/cli/main.go`](go/cmd/cli/main.go) | `cd go && go run ./cmd/cli` |
| 2 | Tests + refactor duplicates | [`go/internal/account/account_test.go`](go/internal/account/account_test.go) | `cd go && go test ./internal/account -cover` |
| 3 | In-memory BankService | [`go/internal/service/bank_service.go`](go/internal/service/bank_service.go) | `cd go && go test ./internal/service -v` |
| 4 | GORM model + SQLite | [`go/internal/models/account.go`](go/internal/models/account.go) | `cd go && go test ./internal/models -v` |
| 5 | Repository pattern | [`go/internal/repository/`](go/internal/repository/) | `cd go && go test ./internal/repository ./internal/service -v` |
| 6 | Fiber REST API | [`go/cmd/api/main.go`](go/cmd/api/main.go), [`go/internal/handlers/account.go`](go/internal/handlers/account.go) | `cd go && go run ./cmd/api` |
| 7 | Postman collection | [`go/postman/bank-api.json`](go/postman/bank-api.json) | All requests green in Postman |
| 8 | JWT auth middleware | [`go/internal/middleware/auth.go`](go/internal/middleware/auth.go) | 401 without token; 200 with Bearer |
| 9 | Docker | [`go/docker/Dockerfile`](go/docker/Dockerfile) | `docker build -f go/docker/Dockerfile -t henloworld-bank go` |
| 10 | Polish (optional) | handlers, service, README | structured logging, `.env`, httptest |

### Checklist

- [x] **Step 1** — CLI demo runs (`go run ./cmd/cli`)
- [x] **Step 2** — Table-driven tests; merge `BankAccount` / `BankAccountImpl`
- [ ] **Step 3** — `BankService` with typed errors; service tests pass
- [ ] **Step 4** — GORM `Account` model; AutoMigrate works
- [ ] **Step 5** — `AccountRepository` + SQLite impl; refactor service to use repo
- [ ] **Step 6** — Four REST endpoints on Fiber (`POST/GET /accounts...`)
- [ ] **Step 7** — Postman collection with `baseUrl` variable
- [ ] **Step 8** — `POST /auth/login` + JWT on `/accounts/*`
- [ ] **Step 9** — Multi-stage Dockerfile; API runs in container
- [ ] **Step 10** — Logging, env config, handler tests (bonus)

### End-state architecture

```
Postman → Fiber → JWT middleware → handlers → BankService → AccountRepository → SQLite
                                              ↘ BankAccountInterface
```

### API reference (implement in Step 6)

| Method | Path | Body |
|--------|------|------|
| POST | `/accounts` | `{ "initial_balance": 1000 }` |
| GET | `/accounts/:id/balance` | — |
| POST | `/accounts/:id/deposit` | `{ "amount": 500 }` |
| POST | `/accounts/:id/withdraw` | `{ "amount": 200 }` |

After Step 8 add `POST /auth/login` with `{ "username": "admin", "password": "secret" }`.

### Dependencies (add as you reach each step)

```powershell
cd go
go get github.com/gofiber/fiber/v2          # Step 6
go get gorm.io/gorm gorm.io/driver/sqlite   # Step 4
go get github.com/golang-jwt/jwt/v5         # Step 8
```

---

## Track B: C / K&R (optional, parallel)

Root C programs — does not block Go progress.

| Exercise | File | Task |
|----------|------|------|
| 1.6 | [`copy_input_to_output.c`](copy_input_to_output.c) | Done — EOF as 0/1 |
| 1.7 | [`copy_input_to_output.c`](copy_input_to_output.c) | Count spaces, tabs, newlines |
| 1.8 | new file or extend copy | Collapse multiple blanks to one |
| 1.9 | new file or extend copy | Collapse multiple newlines to one |

**Verify (C):**

```powershell
cl /nologo copy_input_to_output.c
echo "hello  world" | .\copy_input_to_output.exe
```

---

## Reference files (do not extend until Step 10+)

- [`car_shop.go`](car_shop.go) — OOP reference (`go run car_shop.go`); later reuse Fiber/GORM pattern for a car shop API.
- [`bank.go`](bank.go) — Pointer to `go/cmd/cli`; original logic lives in `go/internal/account/`.

---

## Quick commands

```powershell
# Go backend
cd go
go run ./cmd/cli
go test ./...
go run ./cmd/api

# C
cl /nologo henloworldtwo.c
cl /nologo copy_input_to_output.c

# Docker (C program at repo root)
docker build -t copy-input-to-output .
echo hello | docker run -i --rm copy-input-to-output

# Docker (Go API — after Step 9)
docker build -f go/docker/Dockerfile -t henloworld-bank go
docker run -p 8080:8080 henloworld-bank
```
