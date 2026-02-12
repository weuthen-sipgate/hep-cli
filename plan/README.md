# Implementierungsplan hepic-cli

## Story-Uebersicht

| # | Story | Abhaengigkeiten | Endpoints |
|---|---|---|---|
| 01 | [Projekt-Setup & Grundgeruest](01-project-setup.md) | - | 1 |
| 02 | [Konfiguration & `hepic init`](02-config-and-init.md) | 01 | 1 |
| 03 | [HTTP-Client & Output-Formatter](03-api-client-and-output.md) | 02 | 0 (Infrastruktur) |
| 04 | [Model-Generierung aus Swagger](04-models-generation.md) | 01 | 0 (Infrastruktur) |
| 05 | [Call Search & Transaction](05-call-search-and-transaction.md) | 03, 04 | ~9 |
| 06 | [Call Export](06-export.md) | 05 | ~9 |
| 07 | [Recording & Interception](07-recording-and-interception.md) | 03, 04 | ~9 |
| 08 | [User Management & Auth Tokens](08-user-management.md) | 03, 04 | ~16 |
| 09 | [Netzwerk-Konfiguration](09-network-config.md) | 03, 04 | ~22 |
| 10 | [Agent Management](10-agent-management.md) | 03, 04 | ~7 |
| 11 | [Dashboard & Statistiken](11-dashboard-and-statistics.md) | 03, 04 | ~20 |
| 12 | [Admin, Scripts & Rest](12-admin-and-remaining.md) | 03, 04 | ~25 |
| 13 | [Integration & Polish](13-integration-and-polish.md) | alle | 0 (Qualitaet) |

## Abhaengigkeitsgraph

```
01 Projekt-Setup
├── 02 Config & Init
│   └── 03 HTTP-Client & Output ─┬── 05 Call Search ── 06 Export
│                                ├── 07 Recording
├── 04 Models ───────────────────├── 08 User Management
                                 ├── 09 Netzwerk-Config
                                 ├── 10 Agent Management
                                 ├── 11 Dashboard & Stats
                                 └── 12 Admin & Rest
                                          │
                                 13 Integration & Polish (nach allem)
```

## Reihenfolge

**Phase 1 - Fundament** (sequenziell):
1. Story 01 → 02 → 03 (+ Story 04 parallel zu 02/03)

**Phase 2 - Features** (parallelisierbar nach Phase 1):
- Stories 05-12 koennen unabhaengig voneinander bearbeitet werden
- Empfohlen: 05 → 06 zuerst (Kern-Workflow), dann Rest

**Phase 3 - Abschluss**:
- Story 13 nach Abschluss aller Feature-Stories
