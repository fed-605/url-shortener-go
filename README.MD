## Url Shortener

A production-ready URL shortener with fast redirects and basic authentication for write operations.

### Features
- URL shortening with optional custom aliases
- Redirects via short alias
- Basic Auth for write endpoints
- RESTful API
- PostgreSQL storage
- Redis cache
- Dockerized setup with migrations

### Tech Stack
- Go 1.25+
- PostgreSQL
- Redis
- Chi Router

### Requirements
- Docker + Docker Compose
- Go 1.25+ (local run)
- PostgreSQL 16+ and Redis 7+ (local run)
- `migrate` CLI (local run)
- `make`

### Quick start (Docker)
```bash
git clone https://github.com/fed-605/url-shortener-go.git
cd url-shortener-go
make docker-up
```

### Local run
```bash
git clone https://github.com/fed-605/url-shortener-go.git
cd url-shortener-go
make run
```

### Service URL
- `http://localhost:8080`

### Configuration
- Docker uses `config/config.docker.yaml` and `config/docker.env`.
- Local run uses `config/local.yaml`.

### Make targets
- `make run` — run locally
- `make docker-up` — run with Docker
- `make docker-down` — stop Docker
- `make migrate-up` — apply migrations
- `make migrate-down` — rollback migrations
- `make migrate-version` — show migrations version

### Auth for write endpoints
Write endpoints (`POST /url`, `DELETE /url/{alias}`) use basic auth.

Basic auth from `config/config.docker.yaml`:
- user: `myuser`
- password: `pass`

### API
- `POST /url` — create a short URL (basic auth)
- `GET /{alias}` — redirect to original URL
- `DELETE /url/{alias}` — delete alias (basic auth)
- `GET /url/` — list all records (basic auth)
