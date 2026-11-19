# Saragosa Architecture

## Deployment model

Saragosa uses a hybrid deployment model:

- **Single-server mode:** Suricata, collector, API, database, and web dashboard may run on one Linux host for a small lab.
- **Distributed mode:** one or more Suricata sensors can send EVE JSON to a separately hosted collector/API platform.
- The first implementation remains IDS/passive. Inline IPS is out of scope for the foundation stage.

## Logical components

```text
Network traffic
    |
Suricata sensor (passive IDS)
    |
EVE JSON file
    |
Collector
    |
Event validation and normalization
    |
PostgreSQL (future phase)
    |
Go API
    |
Next.js dashboard
```

## Trust boundaries

1. Sensor-to-collector transport must be authenticated and encrypted before remote ingestion is enabled.
2. The dashboard must not read EVE files directly.
3. The API and collector must not run as root.
4. Any privileged host integration must use a separately reviewed, narrowly scoped helper.
5. The API binds to `127.0.0.1:8080` by default. A trusted reverse proxy is required before remote access.
6. Secrets must be supplied through environment variables or a secret manager, never committed.

## Current foundation

- Go API exposes `GET /api/v1/health`. Its response explicitly reports `api-process-only`; it does not claim Suricata or database health.
- Go collector currently reads EVE JSON Lines from standard input, validates the event type and optional timestamp, and prints aggregate counts. It does not persist events or tail a file yet.
- Next.js dashboard is a truthful static foundation view. Its values are labels, not live sensor metrics.
- Database, authentication, sensor registration, remote transport, and rule management are not implemented yet.

## Next architecture gates

1. Confirm supported Suricata version and EVE schema from official documentation.
2. Decide authenticated local/remote collector transport.
3. Define event identity, deduplication, retention, and PostgreSQL indexes.
4. Define migration and backup/restore strategy.
5. Add authentication and role checks before exposing operational endpoints.
6. Keep inline IPS disabled until a separate lab-tested milestone.
