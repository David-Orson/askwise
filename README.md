# AskWise

**AskWise** is an AI-powered knowledge base SaaS platform that allows teams to upload internal documents and query them using natural language. It simulates a production-grade system at scale (100k users) and is designed to demonstrate full-stack engineering, DevOps, AI integration, and data engineering capability.

> 💡 This is a personal portfolio project intended to look and behave like a real product.

---

## ✨ Key Features

- Upload docs (PDF, Markdown, HTML)
- Ask questions in natural language via AI chatbot
- Admin dashboard with usage analytics and AI feedback
- Simulated user base for metrics like churn, retention, LTV, etc.
- Cloud-native CI/CD pipeline with GitHub Actions + GCP Cloud Run

---

## 🧠 Tech Stack

| Layer       | Stack                                 |
| ----------- | ------------------------------------- |
| Frontend    | React + Vite + SWR                    |
| Backend     | Go (Fiber)                            |
| Database    | PostgreSQL (Docker)                   |
| AI Layer    | Python (LangChain, LangGraph, OpenAI) |
| Data Engine | Python, dbt Core, Jinja               |
| Analytics   | React Dashboards (Simulated BI)       |
| DevOps      | Docker, GitHub Actions, GCP Cloud Run |

---

## 💻 Local Development

### Requirements

- Docker + Docker Compose
- Go ≥ 1.24.5
- Node.js ≥ 18
- PostgreSQL client (optional for manual queries)

### Setup

```bash
# 1. Clone the repo
git clone https://github.com/david-orson/askwise.git
cd askwise

# 2. Start Postgres
docker compose up postgres

# 3. Run backend with hot reload
cd backend
air

# 4. Run frontend with hot reload
cd frontend
npm install
npm run dev
```
