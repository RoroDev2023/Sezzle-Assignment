# Sezzle Calculator Assignment

Full-stack calculator application with a React TypeScript frontend and a Go REST API backend.

## Features

- Addition, subtraction, multiplication, division
- Optional operations: exponentiation, square root, percentage
- Frontend validation for missing or invalid inputs
- Backend validation for invalid JSON, unsupported operations, missing operands, division by zero, negative square root, and non-finite results
- Unit tests for backend calculation/API behavior and frontend UI/API behavior
- Docker Compose setup for running the full stack

## Project Structure

```text
.
├── backend/     # Go REST API
├── frontend/    # React + Vite + TypeScript app
└── docker-compose.yml
```

## Prerequisites

- Go 1.22+
- Node.js 22+ and npm
- Optional: Docker and Docker Compose

## Run Locally

### Backend

```bash
cd backend
go run ./cmd/server
```

The API runs on `http://localhost:8080`.

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend runs on `http://localhost:5173`.

To point the frontend at a different API URL:

```bash
VITE_API_BASE_URL=http://localhost:8080 npm run dev
```

## Run With Docker

```bash
docker compose up --build
```

- Frontend: `http://localhost:5173`
- Backend: `http://localhost:8080`

## API Usage

### Health Check

```bash
curl http://localhost:8080/health
```

Response:

```json
{
  "status": "ok"
}
```

### List Operations

```bash
curl http://localhost:8080/api/operations
```

Response:

```json
{
  "operations": ["add", "subtract", "multiply", "divide", "power", "sqrt", "percentage"]
}
```

### Calculate

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"multiply","a":6,"b":7}'
```

Response:

```json
{
  "operation": "multiply",
  "result": 42
}
```

Square root uses only `a`:

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"sqrt","a":81}'
```

Error example:

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","a":10,"b":0}'
```

```json
{
  "error": "division by zero"
}
```

## Test And Coverage

### Backend

```bash
cd backend
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### Frontend

```bash
cd frontend
npm install
npm run test
```

Vitest writes an HTML coverage report to `frontend/coverage/`.

## Design Decisions

- The backend owns all arithmetic logic so validation and edge-case behavior are consistent for every client.
- The API uses one `POST /api/calculate` endpoint with an explicit `operation` field. This keeps the REST surface small while still exposing all calculator operations.
- The Go backend uses only the standard library. That keeps the service easy to run, test, and audit.
- The React UI performs lightweight client-side validation for fast feedback, then relies on backend errors for authoritative validation.
- Square root is modeled as a unary operation. All other supported operations require operands `a` and `b`.
- Percentage is implemented as `a * b / 100`, for example `20` percent of `10` returns `2`.

## AI Tooling Disclosure

AI tooling was used to scaffold and implement the assignment. Prompts used:

```text
Document the README.md file for the Sezzle Junior Software Engineer assignment as a full-stack calculator app. Document the:
React TypeScript frontend, Go REST API backend, validation, error handling,
responsive UI, tests, setup instructions, API examples, design decisions,
coverage commands, and optional Docker support.

Moreover, install the add on packages for additional functionality
```

