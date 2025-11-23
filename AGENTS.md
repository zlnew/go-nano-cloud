# Repository Guidelines

## Project Structure & Module Organization

- `cmd/api`: entrypoint; wires config, storage, and HTTP router.
- `internal/config`: env parsing for HTTP timeouts and storage path defaults.
- `internal/http`: router, middleware, and request handlers for ping, upload, list, download, and delete.
- `internal/storage`: storage interface plus local disk implementation writing to `uploads/`.
- `uploads/`: default local bucket; useful for manual testing and fixtures.

## Build, Test, and Development Commands

- `make run`: start the server (defaults to `HTTP_ADDRESS=:3000` and `STORAGE_LOCAL_PATH=uploads`).
- `make build`: compile binary to `bin/mini-s3`.
- `make clean`: remove the built binary.
- `go test ./...`: run all tests once they exist; add table-driven tests per handler and storage adapter.
- `gofmt -w . && go vet ./...`: format and lint before sending changes.
- `HTTP_ADDRESS=:8080 make run`: example of overriding the listen address for local dev.

## Configuration & Environment

- Uses `godotenv` in `cmd/api/main.go`; place local overrides in `.env`.
- Key vars: `HTTP_ADDRESS`, `HTTP_READ_TIMEOUT`, `HTTP_WRITE_TIMEOUT`, `HTTP_READ_HEADER_TIMEOUT`, `HTTP_IDLE_TIMEOUT`, `STORAGE_LOCAL_PATH`.
- Protected endpoints (`POST /upload`, `DELETE /delete/<filename>`) require `API_KEY` in env and `X-API-Key` on requests; missing or mismatched keys return 401.
- Request size tuning: `MAX_REQUEST_BODY_SIZE` (MB, default 20) caps the upload body; `MAX_MULTIPART_MEMORY` (MB, default 8) controls multipart parsing memory.
- Keep `STORAGE_LOCAL_PATH` scoped to a sandbox (e.g., `uploads/dev`) to avoid clobbering shared files.

## Coding Style & Naming Conventions

- Go 1.21+ defaults; always run `gofmt`.
- Handlers live under `internal/http/handlers`; name new routes with explicit prefixes (`/files/`, `/delete/`) to match router wiring.
- Middleware composes via `middleware.Chain`; prefer small, focused middleware functions.
- Use clear error strings; prefer wrapping (`fmt.Errorf("copy file: %w", err)`) when plumbing errors up.

## Testing Guidelines

- Use Go’s `testing` and `net/http/httptest` for handlers; mock `storage.Storage` to isolate HTTP behavior.
- Name tests `Test<Thing>` and table-drive edge cases (empty filename, missing file, method mismatch).
- Verify responses include expected status codes and bodies; assert disk effects in a temp `STORAGE_LOCAL_PATH`.

## Commit & Pull Request Guidelines

- Commit messages: short, imperative summary (`add delete handler validation`); include a brief body when behavior changes.
- PRs should describe the endpoint/config changes, include curl examples for new routes, and note any new env vars.
- If touching storage or HTTP behavior, add or update tests and mention coverage areas in the PR description.
