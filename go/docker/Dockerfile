# Track: Go Backend | Step 9 of 10
#
# LEARNING STEP 9: Dockerize the bank API
# Goal: Multi-stage build that runs cmd/api on port 8080.
# Done when: docker build && docker run serves /health (and /accounts after Step 6).
# Verify:
#   docker build -f go/docker/Dockerfile -t henloworld-bank go
#   docker run -p 8080:8080 henloworld-bank
# Reinforces: multi-stage Docker (extends root C Dockerfile pattern)
#
# Tasks:
#   1. Stage 1: golang:1.22-alpine — COPY go.mod, go.sum, source; go build -o /bank-api ./cmd/api
#   2. Stage 2: alpine — COPY binary; EXPOSE 8080; CMD ["./bank-api"]
#   3. ENV PORT=8080 DB_PATH=/data/henloworld.db
#   4. Optional: VOLUME /data for SQLite persistence
#
# LEARNING STEP 9 (bonus): docker-compose.yml at go/docker/ with volume mount for DB.

FROM alpine:3.19
CMD ["echo", "LEARNING STEP 9: Replace this stub with a multi-stage Go build."]
