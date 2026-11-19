# Saragosa

**Saragosa** is a Linux-focused network intrusion detection and prevention management platform built around the Suricata engine.

The project aims to provide a practical web interface for sensor health, security events, alert investigation, ruleset administration, and carefully controlled response workflows.

> **Project status:** Planning / PRD stage. Do not treat planned capabilities as implemented features.

## Project direction

Saragosa does not intend to reimplement packet inspection. It uses **Suricata** for network inspection and builds an application layer for event collection, visibility, configuration management, and audited operations.

### Planned capabilities

- Suricata sensor and service health
- EVE JSON event ingestion
- Alert search, filtering, and investigation
- Ruleset import and validation
- Role-based access control
- Administrative audit logs
- IDS passive mode as the default
- Experimental, opt-in inline IPS workflow

## Proposed technology

| Area | Technology |
|---|---|
| Detection engine | Suricata |
| Backend and collector | Go |
| Web dashboard | Next.js, TypeScript |
| UI | Tailwind CSS, Lucide React |
| Database | PostgreSQL |
| Event format | Suricata EVE JSON |
| Initial target | Ubuntu Server Linux |
| Testing | Go tests, integration tests, PCAP fixtures |

Dependency versions will be pinned during implementation. No performance or detection-rate claims are made before repeatable benchmarks exist.

## Development roadmap

Implementation is divided into gated phases:

1. Repository reconnaissance and planning
2. Project foundation
3. Suricata lab and EVE JSON verification
4. Event collector
5. Alert API and dashboard
6. Authentication, RBAC, and audit
7. Ruleset management
8. Experimental inline IPS in an isolated lab
9. Hardening and release candidate

See [docs/PRD.md](docs/PRD.md) for scope, requirements, acceptance criteria, test strategy, and mandatory AI-agent anti-slop rules.

## Safety notes

- Passive IDS is the default.
- Inline IPS can interrupt network connectivity and must remain opt-in.
- Test blocking behavior only in an isolated lab with a recovery path.
- Saragosa is not a replacement for an enterprise NGFW, SIEM, EDR, or NDR platform.
- Do not expose management interfaces directly to untrusted networks.
- Never commit credentials, private keys, sensitive PCAP files, or production secrets.

## Current limitations

At the PRD stage, no working sensor integration, dashboard, blocking action, or production deployment should be assumed. Feature status must be updated only after implementation and tests provide evidence.

## Documentation

- Indonesian PRD: [docs/PRD.md](docs/PRD.md)
- Indonesian README: [README.id.md](README.id.md)

## License

No license has been selected yet. Until a license is added, do not assume that the code may be reused, modified, or redistributed.
