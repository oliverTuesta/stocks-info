# Backend — Stocks info (Go)

The backend is written in Go and contains two primary runnable parts:

-   `server` — the HTTP API that serves the frontend and exposes endpoints for companies and analyses.
-   `fetcher` — an importer that reads `data/companies.csv` and calls a private source API to retrieve stock analyses.

Project structure highlights

-   `cmd/server` — main for the API server.
-   `cmd/fetcher` — main for the fetcher/importer.
-   `internal/infrastructure/db` — DB connection and migrations.
-   `internal/infrastructure/external` — clients for external/private APIs (analysis provider).
-   `internal/infrastructure/repository` — database repositories for companies and analyses.
-   `internal/usecase` — application usecases (importers and synchronizers).

Prerequisites

-   Go (the project uses Go modules; see `go.mod`).
-   PostgreSQL (or compatible) database.
-   A Google Gemini API key for AI analysis (store it in the `.env`, see below).

.env and configuration

Copy the sample file and fill the values:

```bash
cp .env.sample .env
# then edit .env and add your values
```

Important environment variables (from `.env.sample`):

-   `LOGO_BASE_URL` or `LOGOS_BASE_URL` — base URL used to compose company logo URLs. Note: the code references `LOGOS_BASE_URL` in `cmd/fetcher/main.go`. Make sure you set the matching key.
-   `SOURCE_BASE_URL` — base URL for the private analysis source API.
-   `SOURCE_BEARER` — bearer token for the private analysis source API.
-   `DATABASE_URL` — PostgreSQL connection string (example: `postgres://user:pass@host:port/dbname?sslmode=disable`).
-   `GEMINI_API_KEY` — Google Gemini API key used for AI analysis (required for AI features).

Data file

The fetcher imports companies from `backend/data/companies.csv`. This file is not included in the repo and must be provided by you. A suitable source is the Kaggle dataset:

https://www.kaggle.com/datasets/marketahead/all-us-stocks-tickers-company-info-logos

Logos

Company logos are referenced from this repo (used as the logo base):

https://github.com/ahmeterenodaci/Nasdaq-Stock-Exchange-including-Symbols-and-Logos/tree/main

Running (Makefile)

The repository includes a `makefile` with convenient targets. From the `backend/` directory you can run:

-   `make build` — builds `bin/server` and `bin/fetcher` from the `cmd/` packages.
-   `make run-server` — runs the server using `go run ./cmd/server`.
-   `make run-fetcher` — runs the fetcher using `go run ./cmd/fetcher`.
-   `make test` — runs `go test ./...`.

Typical workflow

1. Provide `backend/data/companies.csv` and update `backend/.env`.
2. Run migrations and import companies + analyses:

```bash
cd backend
make run-fetcher
```

3. Start the API server:

```bash
make run-server
```

Notes and troubleshooting

-   The fetcher calls a private source API to fetch stock analyses — set `SOURCE_BASE_URL` and `SOURCE_BEARER` in the `.env` file so the `internal/infrastructure/external` client can authenticate.
-   The fetcher reads `backend/data/companies.csv` using a relative path; run it from the `backend/` folder or ensure the working directory contains `data/companies.csv`.
-   Gemini: obtain your API key from Google Gemini (Cloud Console / API keys) and set `GEMINI_API_KEY` in the `.env` file.

If you'd like, I can add example `.env` values (redacted), sample CSV columns expected by the importer, or a minimal Postgres docker-compose file for local testing.
