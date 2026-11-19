# Saragosa

**Saragosa** is a Linux-focused network intrusion detection and prevention management platform built around the Suricata engine.

Saragosa provides the application layer for sensor visibility, EVE event collection, alert investigation, ruleset administration, and audited operations. Suricata remains the packet inspection engine.

> **Project status:** Early foundation. The web shell, API health endpoint, and stdin-based EVE JSON parser are present. Database persistence, sensor connectivity, authentication, and active IPS are not implemented.

## Current foundation

- Next.js dashboard shell with honest, non-live status labels.
- Go API with `GET /api/v1/health`, bound to loopback by default.
- Go collector foundation that validates EVE JSON Lines from standard input and reports aggregate counts.
- Go unit tests for the health response contract and EVE event parser.
- Architecture and development documentation.
- GitHub Actions workflow for Go tests and web typecheck/build.

The health endpoint reports API process health only. The collector does not yet tail Suricata files or persist events. No dashboard metric represents live network telemetry.

## Technology

| Area | Technology |
|---|---|
| Detection engine | Suricata (integration planned) |
| Backend and collector | Go 1.27 |
| Web dashboard | Next.js 16, React 19, TypeScript |
| UI | Tailwind CSS 4 |
| Database | PostgreSQL (planned) |
| Event format | Suricata EVE JSON |
| Initial target | Ubuntu Server Linux |
| Runtime | Node.js 26 LTS |

## Repository layout

- `app/`: Next.js dashboard
- `services/api/`: Go HTTP API
- `services/collector/`: EVE JSON parser/collector foundation
- `docs/PRD.md`: product requirements and phased acceptance criteria
- `docs/architecture.md`: deployment model and trust boundaries
- `docs/DEVELOPMENT.md`: local development instructions

## Development

See [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md) for commands and prerequisites.

## Roadmap

1. Foundation
2. Suricata passive IDS lab and EVE verification
3. Persistent collector and event pipeline
4. Alert API and live dashboard
5. Authentication, RBAC, and audit
6. Ruleset management
7. Experimental inline IPS in an isolated lab
8. Hardening and release candidate

## Safety

- Passive IDS is the default.
- Inline IPS is not implemented and must only be tested in an isolated lab.
- The API binds to `127.0.0.1:8080` by default.
- Do not run the application as root.
- Never commit credentials, private keys, sensitive PCAP files, or production secrets.
- Saragosa is not a replacement for an enterprise NGFW, SIEM, EDR, or NDR platform.

## License

See [LICENSE](LICENSE) for the MIT License.
