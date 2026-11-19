# Development Guide

## Prerequisites

- Go 1.27.x
- Node.js 26 LTS
- npm compatible with the selected Node.js release

Use the exact toolchain versions documented in this repository. Re-check official security updates before upgrading dependencies.

## Run the API

```sh
cd services/api
go test ./...
go run ./cmd/api
```

The API listens on `127.0.0.1:8080` by default.

Health check:

```sh
curl http://127.0.0.1:8080/api/v1/health
```

This endpoint reports API-process health only. It does not verify Suricata, the collector, or a database.

## Run the collector foundation

The current collector reads JSON Lines from standard input. It does not open the Suricata log file or write to a database.

```sh
cd services/collector
go test ./...
printf '%s\\n' '{"timestamp":"2026-11-19T10:00:00Z","event_type":"alert"}' | go run ./cmd/collector
```

Malformed lines are counted as rejected and do not stop processing of subsequent lines. The collector does not print raw event payloads.

## Run the dashboard

From the repository root:

```sh
npm install
npm run dev
```

Open `http://localhost:3000`.

Production build:

```sh
npm run typecheck
npm run build
```

## Safety

- Do not enable inline IPS on a production interface during development.
- Do not run the application as root.
- Do not commit `.env`, credentials, private keys, or sensitive PCAP files.
- Do not treat placeholder dashboard labels as sensor telemetry.
