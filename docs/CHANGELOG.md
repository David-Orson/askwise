## [0.1.0] - 2025-07-14

### Added

- First Cloud Run backend deployment
- Working `/api/ping` endpoint
- CI/CD via GitHub Actions to GCP Cloud Run
- Dockerized Go API
- Local development tools (Air)

---

## [0.1.1] - 2025-07-15

### Added

- Set up commit message linting with `commitlint` and `husky`
  - Enforced [Conventional Commit](https://www.conventionalcommits.org/) standard on commits via `commit-msg` hook
- Added `commitlint.config.js` with `@commitlint/config-conventional`
- Installed `husky` with `prepare` script and `.husky/commit-msg` hook

### Changed

- 🧹 Updated `.gitignore` to correctly ignore:
  - Root `node_modules/`
  - Editor & OS junk files
  - Frontend and backend build artifacts

---

## [0.1.2] - 2025-07-16

### Changed

- Rebuilt frontend using **Vite + TypeScript** for better DX and long-term maintainability
- Updated environment variable handling using `.env` and `.env.production`
- Refactored frontend `App.tsx` to pull API base URL from `VITE_API_URL`

### Fixed

- CORS headers enabled in backend to allow frontend → backend communication

### Deployed

- Frontend deployed to Firebase Hosting
- Working client-to-backend ping request

---

## [0.1.3] - 2025-07-20

### Fixed

- Updated backend CORS config to allow requests from deployed Firebase frontend

---

## [0.1.4] - 2025-07-20

### Changed

- Added Firebase deploy to CI/CD workflow
- Updated GitHub Actions to build & deploy frontend on push to `main`

---

## [0.1.5] - 2025-07-24

Changed

- Removed legacy React + Vite frontend
- Rebuilt frontend using Next.js App Router
- Integrated Tailwind CSS and configured PostCSS
- Added shadcn/ui, lucide-react, class-variance-authority, and tailwind-variants
- Setup layout structure with app/, components/, lib/, and styles/
- Introduced Zustand and React Query for state/data management

Added

- Created a full-featured landing page:
- Responsive hero section with CTA
- "How it works" walkthrough
- Feature highlights with reusable <FeatureCard />
- Fake testimonials and modern footer
- Set up Vercel as frontend hosting provider (CI/CD via GitHub + Vercel integration)
- Created deploy-backend.yml for backend-only Cloud Run deploy pipeline

Removed

- Removed Firebase Hosting deploy from CI/CD workflow

---

## [0.1.6] - 2025-07-25

### Added

- Integrated NextAuth.js with Google OAuth
- Protected `/dashboard` and `/project/[id]` routes behind authentication
- Created authenticated dashboard layout showing sample projects
- Created project page view with document list and AI query input
- Added "Back to dashboard" navigation with Lucide icon

### Changed

- Fixed server-side dynamic route usage with `await` on `params` for Next.js App Router
- Improved mobile layout + spacing consistency across homepage and dashboard

---

## [0.1.7] - 2025-07-26

### Added

- Built full Go backend with Fiber, GORM, and Neon (Postgres)
- Defined core data models: `User`, `Project`, `Document`, `Question`, `Answer`
- Created JWT-based auth middleware with claim extraction (via Google OAuth `sub` or `email`)
- Added `/auth/sync` endpoint to persist users after frontend login
- Connected frontend session (via NextAuth) to backend user creation
- Created `POST /api/projects` and `GET /api/projects` with per-user filtering
- Wrote middleware to verify and inject user claims into route context
- Used custom `GoogleID` field to map external identities to internal UUIDs

### Changed

- Replaced `DebugUserID` with dynamic user resolution via JWT/email
- Updated all project routes to respect authenticated user scope
- Improved error handling and response codes across backend:
  - `400` for bad input
  - `409` for existing users
  - `401` for unauthenticated access
  - `500` for server-side errors

### Fixed

- Infinite sync loop on login by guarding mutation with `useRef`
- CORS crash on authorized request (replaced `*` with specific allowed origins)

### DevOps

- Added new backend route groups: `/auth`, `/api`, `/public`
- Ready for production deploy via Cloud Run backend + Vercel frontend

---

## [0.1.8] - 2025-07-26

### Added

- Created modular document/ service using Hexagonal architecture
- Defined Document domain model with encapsulated fields and base metadata
- Added PostgresDocumentRepository implementing domain DocumentRepository port
- Built UploadDocument() application service to create + persist documents
- Emitted document.uploaded events via Redis Stream using EventBus abstraction
- Created Redis-backed EventBus with Publish() and Subscribe() support
- Connected HTTP upload route to backend logic using Fiber
- Defined DocumentResponse DTO + mapper in presentation layer
- Wrote unit tests for DocumentService.UploadDocument() using mocked ports

### Changed

- Refined domain encapsulation: made projectID, userID, fileName private with getters
- Split responsibilities cleanly across domain, application, handler, and adapters
- Enforced strict layering between core logic and infrastructure via interfaces

---

## [0.1.9] - 2025-07-27

### Added

- Introduced `DocumentRecord` (GORM model) and persistence mappers to bridge domain <-> DB
- Added SQLite-backed repository integration test for PostgresDocumentRepository
- Table-driven tests for `UploadHandler` and domain validation logic
- Created `DocumentHandler` using Fiber, wired to the application service and user context
- `FromDomain()` presenter added defensive nil handling and tests

### Changed

- Domain model now uses `DomainBase` instead of GORM-coupled base struct
- Extracted `GormBase` (with tags) and `DomainBase` (clean) for separation of concerns
- Renamed repository test to use in-memory SQLite instead of Postgres for speed
- Enforced stricter invariants on filenames (no slashes, max length, lowercase extensions)
- Improved handler robustness by parsing `userID` from request context properly
- Unified mocking and modularized test helpers across layers

---

## [0.2.0] - 2025-07-29

### Added

- Introduced `user` module using Hexagonal architecture:
  - `User` domain model with encapsulated fields (`googleID`, `name`, `email`, `imageURL`)
  - `UserRecord` GORM struct with mapping functions (`toDomain`, `toRecord`)
  - `PostgresUserRepository` with methods: `FindByID`, `FindByGoogleID`, `Save`
  - `UserService.Sync()` handles "create-or-fetch" by Google ID
  - `UserHandler.Sync()` Fiber handler accepting Google ID and user details
- Defined presentation DTO (`UserResponse`) with safe field mapping
- Added `extractBearerToken()` helper for parsing `Authorization` headers

### Changed

- Split base structs:
  - `DomainBase` (clean, untagged) for domain logic
  - `GormBase` (with GORM tags) for persistence concerns
- Wired service-to-handler mapping cleanly through interfaces

---
