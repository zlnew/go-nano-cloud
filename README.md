# go-nano-cloud

Minimal HTTP service for uploading, listing, downloading, and deleting files on local disk. Two buckets are mounted by default: a public bucket that allows unauthenticated reads and a private bucket that is fully protected by an API key.

## Features

- Public and private buckets rooted under the configured storage path.
- Upload arbitrary files via multipart POST (field `file`).
- List files with size and path metadata.
- Download files with content-disposition headers.
- Delete files by key.
- Pluggable storage interface (`internal/storage`) with a local disk implementation.
- API key middleware protects mutating endpoints and all private bucket access.
- Configurable request body and multipart limits.

## Project Layout

- `cmd/api`: service entrypoint; loads env, wires storage and Chi router.
- `internal/config`: environment parsing for HTTP settings, size limits, and storage path.
- `internal/http`: router, middleware, and handlers.
- `internal/storage`: storage interface and local disk backend.
- `uploads/`: default storage root; buckets are created beneath it (`uploads/public`, `uploads/private`).

## Quick Start

Requirements: Go 1.25+, Bash-compatible shell.

```bash
# optional: define local overrides
cat > .env <<'EOF'
HTTP_ADDRESS=:3000
STORAGE_LOCAL_PATH=uploads
API_KEY=local-dev-key
# MAX_REQUEST_BODY_SIZE=20
# MAX_MULTIPART_MEMORY=8
EOF

# run locally
make run
```

Server starts on `HTTP_ADDRESS` (default `:3000`).
Endpoints under `/private` and all write operations expect header `X-API-Key` matching `API_KEY`.

## Configuration

Set via environment variables or `.env` (loaded by `godotenv`):

- `HTTP_ADDRESS` (default `:3000`)
- `HTTP_READ_TIMEOUT` (seconds, default `5`)
- `HTTP_WRITE_TIMEOUT` (seconds, default `10`)
- `HTTP_READ_HEADER_TIMEOUT` (seconds, default `2`)
- `HTTP_IDLE_TIMEOUT` (seconds, default `60`)
- `STORAGE_LOCAL_PATH` (default `uploads`)
- `API_KEY` (required for protected routes)
- `MAX_REQUEST_BODY_SIZE` (MB, default `20`)
- `MAX_MULTIPART_MEMORY` (MB, default `8`)

Point `STORAGE_LOCAL_PATH` to a dedicated folder (e.g., `uploads/dev`) to keep environments isolated. Bucket content is stored relative to that root.

## API Endpoints (examples assume `:3000`)

Uploads and deletes always require `X-API-Key`. The private bucket requires `X-API-Key` for all operations.

**Public bucket (read without auth, writes require `X-API-Key`):**

- List: `GET /public`
  ```bash
  curl -i http://localhost:3000/public
  ```
- Download: `GET /public/<key>`
  ```bash
  curl -L -o saved.txt http://localhost:3000/public/path/to/file.txt
  ```
- Upload: `POST /public` (multipart form field `file`)
  ```bash
  curl -i -H "X-API-Key: $API_KEY" -F "file=@/path/to/file.txt" http://localhost:3000/public
  ```
- Delete: `DELETE /public/<key>`
  ```bash
  curl -i -H "X-API-Key: $API_KEY" -X DELETE http://localhost:3000/public/path/to/file.txt
  ```

**Private bucket (all operations require `X-API-Key`):**

- List: `GET /private`
  ```bash
  curl -i -H "X-API-Key: $API_KEY" http://localhost:3000/private
  ```
- Download: `GET /private/<key>`
  ```bash
  curl -L -H "X-API-Key: $API_KEY" -o saved.txt http://localhost:3000/private/path/to/file.txt
  ```
- Upload: `POST /private` (multipart form field `file`)
  ```bash
  curl -i -H "X-API-Key: $API_KEY" -F "file=@/path/to/secret.txt" http://localhost:3000/private
  ```
- Delete: `DELETE /private/<key>`
  ```bash
  curl -i -H "X-API-Key: $API_KEY" -X DELETE http://localhost:3000/private/path/to/file.txt
  ```

List responses are JSON arrays that include `path`, `name`, and `size` for each object. Filenames are cleaned server-side; absolute paths and attempts to traverse (`..`) are rejected.

## Development Notes

- Make targets: `make run` (start), `make build` (binary at `bin/nano-cloud`), `make clean`.
- Format and lint: `gofmt -w . && go vet ./...`
- Tests: `go test ./...` (add table-driven tests for handlers and storage).
- Middleware composes via `middleware.Chain`; add lightweight cross-cutting concerns here.
- Errors: prefer wrapped errors when propagating (`fmt.Errorf("save file: %w", err)`).

## Contributing

- Use short, imperative commit messages (e.g., `add upload handler validation`).
- Include curl examples and env var notes in PR descriptions when changing endpoints/config.
- If storage or HTTP behavior changes, add or update tests and mention coverage in the PR.
