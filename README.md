# henloworld

Small C programs while working through *The C Programming Language* (K&R), a Go backend learning path, and Docker setups.

## Go learning path

Bank account OOP code grows into a full REST API in [`go/`](go/). Follow step-by-step tasks in **[LEARNING.md](LEARNING.md)** — search the repo for `LEARNING STEP` comments.

| File | Description |
|------|-------------|
| `go/cmd/cli` | Bank account demo (Step 1) |
| `go/cmd/api` | Fiber REST API (Steps 6–8) |
| `go/internal/account` | Domain model + interfaces |
| `bank.go` | Redirect — run `cd go && go run ./cmd/cli` |
| `car_shop.go` | Standalone OOP reference |

```powershell
cd go
go run ./cmd/cli          # Step 1 demo
go test ./...             # after Step 2+
go run ./cmd/api          # after Step 6
```

Stack: **Fiber**, **GORM + SQLite**, **JWT**, **Docker**, **Postman**.

## Programs (C)

| File | Description |
|------|-------------|
| `henloworldtwo.c` | Fahrenheit–Celsius table (exercises 1.2–1.5) |
| `copy_input_to_output.c` | Copy stdin to stdout; verifies `getchar() != EOF` is 0 or 1 (exercise 1.6) |
| `henloworld.cpp` | Hello-world C++ sample |

## Build and run (Windows / MSVC)

Requires Visual Studio with the C++ build tools and `cl.exe` on your PATH (or a Developer PowerShell terminal).

```powershell
cl /nologo henloworldtwo.c
.\henloworldtwo.exe

cl /nologo copy_input_to_output.c
.\copy_input_to_output.exe
```

`copy_input_to_output` reads from the keyboard until EOF (**Ctrl+Z**, then **Enter** on Windows). Or pipe input:

```powershell
echo hello | .\copy_input_to_output.exe
```

## Docker

A multi-stage `Dockerfile` builds and runs `copy_input_to_output.c` with **gcc** on Linux.

**Build:**

```bash
docker build -t copy-input-to-output .
```

**Run with piped input:**

```bash
echo hello | docker run -i --rm copy-input-to-output
```

**Run interactively** (type input, then **Ctrl+D** for EOF):

```bash
docker run -i --rm copy-input-to-output
```

Use `-i` so the program can read from stdin via `getchar()`.

## Docker (Go API)

After completing Step 9 in [LEARNING.md](LEARNING.md):

```bash
docker build -f go/docker/Dockerfile -t henloworld-bank go
docker run -p 8080:8080 henloworld-bank
```

## Editor setup

`.vscode/` includes build tasks and launch configs for **VS Code** (debug with `cl.exe`) and **Cursor** (build and run).
