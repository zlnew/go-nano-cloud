# Mini S3 (Local Object Store)

Minimal HTTP service for uploading, listing, downloading, and deleting files on local disk. Handy for prototyping or running lightweight storage in development environments.

## Features
- Upload arbitrary files via multipart POST.
- List stored files in the configured root.
- Download files with proper content headers.
- Delete files by path.
- Pluggable storage interface (`internal/storage`) with a local disk implementation.

## Project Layout
- `cmd/api`: service entrypoint; loads env, wires storage and router.
- `internal/config`: environment parsing for HTTP settings and storage path.
- `internal/http`: router, middleware, and handlers.
- `internal/storage`: storage interface and local disk backend.
- `uploads/`: default storage root for local runs.

## Quick Start
Requirements: Go 1.25+, Bash-compatible shell.

```bash
# optional: define local overrides
cat > .env <<'EOF'
HTTP_ADDRESS=:3000
STORAGE_LOCAL_PATH=uploads
EOF

# run locally
make run
```

Server starts on `HTTP_ADDRESS` (default `:3000`).

## Configuration
Set via environment variables or `.env` (loaded by `godotenv`):
- `HTTP_ADDRESS` (default `:3000`)
- `HTTP_READ_TIMEOUT` (seconds, default `5`)
- `HTTP_WRITE_TIMEOUT` (seconds, default `10`)
- `HTTP_READ_HEADER_TIMEOUT` (seconds, default `2`)
- `HTTP_IDLE_TIMEOUT` (seconds, default `60`)
- `STORAGE_LOCAL_PATH` (default `uploads`)

Point `STORAGE_LOCAL_PATH` to a dedicated folder (e.g., `uploads/dev`) to keep environments isolated.

## API Endpoints (examples assume `:3000`)
- Health: `GET /ping`
  ```bash
  curl -i http://localhost:3000/ping
  ```
- Upload: `POST /upload` (multipart form field `file`)
  ```bash
  curl -i -F "file=@/path/to/file.txt" http://localhost:3000/upload
  ```
- List files: `GET /files`
  ```bash
  curl -i http://localhost:3000/files
  ```
- Download: `GET /files/<filename>`
  ```bash
  curl -L -o saved.txt http://localhost:3000/files/file.txt
  ```
- Delete: `DELETE /delete/<filename>`
  ```bash
  curl -i -X DELETE http://localhost:3000/delete/file.txt
  ```

## Development Notes
- Make targets: `make run` (start), `make build` (binary at `bin/mini-s3`), `make clean`.
- Format and lint: `gofmt -w . && go vet ./...`
- Tests: `go test ./...` (add table-driven tests for handlers and storage).
- Middleware composes via `middleware.Chain`; add lightweight cross-cutting concerns here.
- Errors: prefer wrapped errors when propagating (`fmt.Errorf("save file: %w", err)`).

## Contributing
- Use short, imperative commit messages (e.g., `add upload handler validation`).
- Include curl examples and env var notes in PR descriptions when changing endpoints/config.
- If storage or HTTP behavior changes, add or update tests and mention coverage in the PR.
