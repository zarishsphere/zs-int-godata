# PRD-MVP — `zs-int-godata`

> **Document:** Product Requirements (MVP) | **Version:** 1.0.0-mvp
> **Repository:** [https://github.com/zarishsphere/zs-int-godata](https://github.com/zarishsphere/zs-int-godata)
> **Layer:** Layer 10 — Integrations | **Catalog #:** 204
> **Language:** Go 1.26.1 | **License:** Apache 2.0

---

## Executive Summary

**Go.Data contact tracing export.**

This document defines the **Minimum Viable Product (MVP)** scope for `zs-int-godata` within the ZarishSphere sovereign digital health platform. It covers what must be built first, acceptance criteria, user stories, and the complete repository file structure.


### Platform Non-Negotiables (apply to every repository)

| Constraint | Rule |
|-----------|------|
| **Zero Cost** | All tooling, hosting, and services must use genuinely free tiers |
| **Open Source** | Apache 2.0 license; all code public |
| **FHIR R5 Native** | All clinical data modelled as FHIR R5 resources |
| **Offline-First** | Must function without network connectivity |
| **No-Coder Friendly** | GUI-first, template-driven, automatable |
| **Documentation as Code** | All decisions in GitHub via RFC/ADR |
| **Multi-tenant** | tenant_id scoping on all data operations |
| **HIPAA/GDPR** | AuditEvent on all PHI access; field-level encryption |

---

## Problem Statement

Outbreak response requires case and contact data to flow from ZarishSphere into Go.Data for contact tracing investigation.

## MVP Goals

1. Implement the primary data flow direction (Outbound (ZS → Go.Data))
2. Field-level mapping between GODATA and FHIR R5 documented and implemented
3. Error handling with dead-letter queue and retry
4. Integration test with real (mock) GODATA endpoint

## MVP Functional Requirements

| ID | Requirement | Acceptance Criteria | Priority |
|----|------------|---------------------|---------|
| M-01 | Primary direction (Outbound (ZS → Go.Data)) working | Data flows; FHIR resources created/updated | P0 |
| M-02 | Field mapping documented in MAPPING.md | All mapped fields listed with types | P0 |
| M-03 | Retry with exponential backoff | Failed requests retry up to 5 times | P0 |
| M-04 | Dead-letter queue via NATS | Failed messages visible in NATS DLQ | P1 |
| M-05 | FHIR AuditEvent for each exchange | AuditEvent appears in audit log | P1 |

## Protocol Support

- Go.Data REST API
- FHIR R5 Case → Go.Data outbreak case

## Mapping Note

FHIR Patient + Condition (confirmed case) → Go.Data case record

## MVP Complete Repository Tree

```
zs-int-godata/
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
├── .env.example
├── .gitignore
├── CHANGELOG.md
├── .github/
│   ├── CODEOWNERS
│   └── workflows/
│       └── ci.yml
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── adapter/
│   │   ├── godata_client.go           # GODATA API client
│   │   └── godata_client_test.go
│   ├── mapper/
│   │   ├── to_fhir.go                     # GODATA → FHIR R5 mapper
│   │   ├── from_fhir.go                   # FHIR R5 → GODATA mapper
│   │   └── mapper_test.go
│   ├── sync/
│   │   ├── syncer.go                      # Sync orchestration
│   │   └── retry.go                       # Exponential backoff retry
│   └── config/
│       └── config.go
├── config/
│   └── config.yaml
├── docs/
│   ├── openapi.yaml
│   └── MAPPING.md                         # Full field mapping table
└── tests/
    └── integration/
        └── suite_test.go
```

---


## Owners & Governance

| Role | GitHub Handle | Responsibility |
|------|--------------|----------------|
| Platform Lead | `@arwa-zarish` | Final approval, RFC votes |
| Technical Lead | `@code-and-brain` | Architecture, Go/TS review |
| DevOps Lead | `@DevOps-Ariful-Islam` | CI/CD, infra, deployment |
| Health Programs | `@BGD-Health-Program` | Clinical content, country programs |

**PR Policy:** All changes via Pull Request. Minimum 1 owner review. CI must pass. No direct commits to `main`.


---

## MVP Acceptance Checklist

- [ ] All MVP files exist in repository with real content (not placeholders)
- [ ] CI pipeline passes on `main` branch
- [ ] No secrets, credentials, or PHI committed
- [ ] README.md reflects current state with setup instructions
- [ ] CODEOWNERS file present
- [ ] All MVP functional requirements verified manually or via automated tests
- [ ] Linked to `CATALOGS.md` and `TODO.md` in `zs-docs-platform`

---

*This document is the authoritative MVP specification for `zs-int-godata`.*
*Changes require a Pull Request with at least 1 owner approval.*
