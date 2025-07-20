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
