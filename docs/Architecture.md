# AskWise Architecture

## Overview

AskWise is structured like a production-grade SaaS app, meant to showcase realistic engineering decisions and practical cloud-ready development.

---

## Layered Architecture

### Frontend (React + Vite)

- Fast iteration with Vite
- API communication via `/api/*` proxy
- Modular components for upload, chat, analytics

### Backend (Go + Fiber)

- RESTful API for organisations, users, documents, and chats
- Connected to PostgreSQL with retry logic
- Designed for performance and concurrency

### Database (PostgreSQL)

- Containerized via Docker Compose
- Will store organisations, users, documents, and chat logs

### AI Layer (LangChain + Python)

- Separate microservice to handle doc ingestion, embeddings, and Q&A
- Will integrate with LangChain + OpenAI for now, later support Anthropic or self-hosted

### Infrastructure

- Dockerized services for consistent builds
- GitHub Actions CI/CD deploys to GCP Cloud Run
- `.env`-based config for local vs prod separation

---

## Design Rationale

| Decision           | Rationale                                            |
| ------------------ | ---------------------------------------------------- |
| Go backend         | Fast, efficient, and realistic for SaaS-scale APIs   |
| Postgres           | Reliable RDBMS with strong local + cloud support     |
| React + Vite       | Fast dev cycle, modern SPA setup                     |
| Docker Compose     | Easy service orchestration for dev                   |
| Cloud Run          | Simple, cost-effective managed deploy for containers |
| LangChain AI Layer | Best-in-class orchestration for RAG-based AI flows   |

---

## 🔧 Local Development Tooling

### Backend Hot Reload: `air`

We use [`air`](https://github.com/air-verse/air) for live-reloading the Go backend during development.

**Why this matters:**

- Running the backend inside Docker alone results in sluggish rebuild times
- `air` runs Go outside Docker and watches files for changes
- Rebuild + restart is near-instant, leading to faster feedback and iteration

**Alternative considered:**

- Docker with live-reload volume mounts + `reflex`/`air` inside container
  - Rejected for being slower, more complex to configure cross-platform
- Fully containerized dev environment
  - Does advantage minimal setup for new developers
  - Considered unnecessary for solo/local dev; can revisit if scaling contributors

**How it works:**

- Run `docker compose up postgres` to start DB
- Run backend with `air` (hot reload)
- Environment variables loaded via `.env` file using `godotenv`
