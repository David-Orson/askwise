# 🧾 Project Journal — AskWise

## 📅 2025-07-14 — First Cloud Run Deployment Success

**Goal:** Deploy the backend service to GCP Cloud Run via GitHub Actions.

---

### ✅ What Worked

- Backend container built successfully via GitHub Actions.
- Docker image pushed to **Artifact Registry** (`europe-west2-docker.pkg.dev`).
- Service deployed to Cloud Run (`askwise-backend`).

---

### 🧱 Key Steps

- Created a new **GCP project** and **Artifact Registry** in `europe-west2`.
- Generated a **service account JSON key** and added it to GitHub secrets:
  - `GCP_SA_KEY`
  - `GCP_PROJECT_ID`
- Set up:

  ```bash
  gcloud auth configure-docker europe-west2-docker.pkg.dev
  ```

- Updated the GitHub Actions workflow to:
  - Use Artifact Registry region-specific path
  - Reference the new project ID and service account

---

## 🐛 Issues Faced

- ❌ **Missing artifact registry permission**  
  Resolved by granting `roles/artifactregistry.writer` to the service account.

- ❌ **Cloud Run deploy failed**  
  **Error:** `PORT=8080` timeout.  
  **Root Cause:** Backend attempted to connect to a Postgres DB that doesn’t exist in the cloud environment.

  **Fix:** Temporarily commented out DB connection in Go code.

- ❌ **`iam.serviceAccountUser` permission missing**  
  Resolved via:

  ```bash
  gcloud iam service-accounts add-iam-policy-binding \
    [COMPUTE_SA_EMAIL] \
    --member="serviceAccount:[YOUR_SA_EMAIL]" \
    --role="roles/iam.serviceAccountUser"
  ```

---

## 📅 2025-07-16 — Frontend Redeployed with TypeScript + Firebase

**Goal:** Rebuild the frontend with TypeScript and deploy it to Firebase Hosting.

---

### ✅ What Worked

- Replaced previous JS frontend with Vite + React + TypeScript setup
- Connected frontend to backend via `VITE_API_URL` env var
- Updated backend to include CORS headers (allow Firebase-hosted frontend)
- Deployed static frontend to Firebase

---

### ⚙️ Tech Decisions

- Used `.env` for local dev, `.env.production` for Firebase builds
- Used `fetch()` with `${import.meta.env.VITE_API_URL}` instead of hardcoding base URLs

---

### 🔗 Result

Frontend app now **successfully fetches backend `/api/ping`** via Cloud Run from Firebase.

---

## 📅 2025-07-25 — Frontend Auth + Dashboard Bootstrapped

**Goal:** Add Google Sign-In and authenticated user flow

---

### ✅ What Worked

- Added `next-auth` and configured Google Provider via `.env.local`
- Created `lib/auth.ts` with reusable `authOptions`
- Protected authenticated routes (`/dashboard`, `/project/[id]`) using `getServerSession()`
- Added SessionProvider to `RootLayout` for auth-aware context
- Verified JWT and tested login in both local and deployed Vercel environment

---

### 🧱 Features Implemented

- 🌐 Project list with fake sample data on dashboard
- 📄 Individual project page with mock document list
- ❓ Question input to simulate AI chat interaction
- ⬅️ Back button from project → dashboard (using `ArrowLeft` icon)

---

### 🐛 Gotchas

- ❌ `params.id` used synchronously in server component (triggered Next.js warning)
  ✅ Fixed by `await`ing `params` destructure properly
- ❌ Login flow failed with `redirect_uri_mismatch`
  ✅ Resolved by updating OAuth Client to match Vercel domain

---

### 🧩 Next Goals

- Design DB schema to store projects, users, documents
- Add `/api/projects` and `/api/projects/:id/upload`
- Begin wiring frontend actions to backend Go API

---

## 📅 2025-07-26 — Backend API + Auth Sync Complete

**Goal:** Connect frontend Google login to actual backend user persistence and enable real project creation.

---

### ✅ What Worked

- Created core Go backend with Fiber + GORM + Postgres (Neon)
- Defined clean REST API grouping (`/auth`, `/api`, `/public`)
- Implemented robust user sync endpoint: `POST /auth/sync`
- Used JWT Bearer tokens (via NextAuth session) to pass user identity securely
- Created middleware to parse JWT, extract `email`, and inject context
- Hooked up frontend React Query `useSyncUser()` to fire after login
- Confirmed real user is persisted (UUID + Google `sub`) in DB

---

### 🧱 Features Implemented

- 👤 Dynamic user sync (creates user on first login)
- 🗂️ Created project API (`/api/projects`)
- 🔒 Projects tied to specific user via backend-side filtering
- 🔐 JWT middleware verifies signature using shared `NEXTAUTH_SECRET`

---

### 🐛 Gotchas

- ❌ Using `userID` from Google directly → caused UUID mismatch in DB
  ✅ Resolved by introducing `GoogleID` field in user model

- ❌ CORS rejected authorized frontend request
  ✅ Fixed by replacing `*` with real frontend origin(s)

- ❌ Infinite React Query sync loop after error
  ✅ Prevented with `useRef` + `onError` fallback

---

### 🚀 Next Goals

- Hook up real file upload and document metadata endpoint
- Add `useCreateProject()` and connect "New Project" UI
- Begin chunking + embedding pipeline for uploaded documents
- Store embeddings in Neon with PGVector
