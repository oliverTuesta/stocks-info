# Stocks info

Stocks info is a web application that shows top stocks from top companies and analysis from brokerage firms. It also surfaces AI Analysis produced by the Google Gemini LLM.

Author: Oliver Tuesta — https://spigis.com/

## Project layout

-   `backend/` — Golang backend (server + fetcher). Handles data import, storage, and exposes the API.
-   `frontend/` — Vue 3 + Tailwind CSS 4 SPA (uses TanStack tables) that consumes the backend API.

## Quick start

1. Backend preparation

    - Copy `backend/.env.sample` to `backend/.env` and fill the values (see `backend/README.md`).
    - Place `companies.csv` into `backend/data/companies.csv`. You can obtain a suitable CSV from the Kaggle dataset:
      https://www.kaggle.com/datasets/marketahead/all-us-stocks-tickers-company-info-logos
    - Ensure a PostgreSQL database is available and set `DATABASE_URL` in `backend/.env` accordingly.
    - Add your Google Gemini API key to `backend/.env` as `GEMINI_API_KEY`.

2. Run the fetcher to import companies and analyses, then run the server (backend):

```bash
cd backend
make run-fetcher   # imports companies from companies.csv and fetches analyses
make run-server    # starts the API server
```

3. Frontend (development):

```bash
cd frontend
npm install
npm run dev
```

## Important links

-   Author blog: https://spigis.com/
-   Kaggle companies CSV: https://www.kaggle.com/datasets/marketahead/all-us-stocks-tickers-company-info-logos
-   Logos repo: https://github.com/ahmeterenodaci/Nasdaq-Stock-Exchange-including-Symbols-and-Logos

## Notes and caveats

-   The fetcher expects a `companies.csv` file in `backend/data/`. Without it the importer will fail.
-   The repository references company logos via an external base URL. The sample `.env.sample` provides a `LOGO_BASE_URL` value, but the fetcher code reads `LOGOS_BASE_URL` (plural). Make sure the variable you set in `backend/.env` matches what the running program expects (`LOGOS_BASE_URL`) or update the code to use your preferred name.
-   Gemini: get an API key from Google Gemini and add it to `backend/.env` as `GEMINI_API_KEY` — the backend uses this for AI analysis.
-   Database: the backend expects a PostgreSQL-compatible database. Put the connection string in `backend/.env` as `DATABASE_URL`.
