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

- ✅ Set up commit message linting with `commitlint` and `husky`
  - Enforced [Conventional Commit](https://www.conventionalcommits.org/) standard on commits via `commit-msg` hook
- ✅ Added `commitlint.config.js` with `@commitlint/config-conventional`
- ✅ Installed `husky` with `prepare` script and `.husky/commit-msg` hook

### Changed

- 🧹 Updated `.gitignore` to correctly ignore:
  - Root `node_modules/`
  - Editor & OS junk files
  - Frontend and backend build artifacts

---
