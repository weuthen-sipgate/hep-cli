# hepic-cli

## Produktvision

**hepic-cli** ist ein kommandozeilenbasiertes Werkzeug, das als Bruecke zwischen Claude Code und der HEPIC-API dient. Es ermoeglicht einem KI-Agenten (Claude Code), die HEPIC-Plattform vollstaendig ueber strukturierte CLI-Befehle zu bedienen.

Die HEPIC-Plattform ist ein Telekommunikations-Monitoring- und Call-Forensik-System fuer SIP-Daten mit 121 API-Endpoints in 25 Kategorien. Die API-Spezifikation liegt in `swagger.json` (Swagger 2.0, API v1.2.1).

### Kernziele

- **Agent-First Design** -- JSON-Default-Ausgabe, konsistente Befehlsstruktur, klare Fehlermeldungen
- **Vollstaendige API-Abdeckung** -- Alle 121 Endpoints ueber CLI erreichbar
- **Intuitives Ressourcen-Modell** -- `hepic <ressource> <aktion> [optionen]`

### Zielgruppen

- Primaer: Claude Code / KI-Agenten
- Sekundaer: Netzwerk-Ingenieure & SIP-Analysten

## Technische Entscheidungen

| Aspekt | Entscheidung |
|---|---|
| Sprache | Go |
| CLI Framework | Cobra + Viper |
| Codegen-Strategie | Hybrid: Models/Types aus swagger.json generiert, HTTP-Client + CLI manuell |
| Generierter Code | Eingecheckt ins Repository |
| Go-Modul | `hepic-cli` (lokaler Modulpfad) |
| Authentifizierung | Nur API-Key (Auth-Token Header), kein JWT Login-Flow |
| Konfiguration | `~/.hepic/config.yaml` + Umgebungsvariablen (HEPIC_HOST, HEPIC_TOKEN) |
| Output-Default | JSON, mit `--format=table/yaml` fuer menschliche Nutzer |
| Code-Organisation | Nach Ressource gruppiert |
| Tests | Unit-Tests mit Mock-HTTP-Client |
| Scope | Alle 121 Endpoints |

## Projektstruktur

```
hepic-cli/
├── CLAUDE.md
├── swagger.json
├── plan/                    # Implementierungs-Stories
├── cmd/                     # Cobra Command-Definitionen
│   └── root.go              # Root-Command, globale Flags
├── internal/
│   ├── api/                 # HTTP-Client, Auth, Request/Response-Handling
│   ├── config/              # Viper-Config, ~/.hepic/config.yaml
│   ├── models/              # Generierte Types/Structs aus Swagger
│   ├── output/              # Output-Formatter (JSON, Table, YAML)
│   ├── call/                # Call Search, Transaction, Reports
│   ├── recording/           # Call Recordings, Interceptions
│   ├── export/              # PCAP, SIPp, Text Export
│   ├── user/                # User Management, Auth Tokens
│   ├── agent/               # Agent Subscribe, Search
│   ├── config_resources/    # IP Alias, HEPSub, Mapping, Protocol
│   ├── dashboard/           # Dashboard CRUD
│   ├── statistic/           # Statistiken, Prometheus, Grafana
│   ├── admin/               # Admin, ConfigDB, Version
│   └── script/              # Script Management
├── main.go
├── go.mod
└── go.sum
```

## Befehlsstruktur

```
hepic init                              # Interaktive Erstkonfiguration
hepic <ressource> <aktion> [flags]      # Standard-Muster
```

Beispiele:
```
hepic call search --from 2025-01-01 --caller "+49..."
hepic call transaction --call-id <id>
hepic recording list --from 2025-01-01
hepic recording play <uuid>
hepic export pcap --call-id <id> -o capture.pcap
hepic ipalias list
hepic user list
hepic version
```

## Globale Flags

- `--format json|table|yaml` (default: json)
- `--host <url>` -- Ueberschreibt Config/Env
- `--token <api-key>` -- Ueberschreibt Config/Env
- `--verbose` -- Debug-Ausgabe
- `--no-color` -- Keine ANSI-Farben (fuer Pipes/Agent-Nutzung)

## Konventionen

### Code-Stil
- Go Standard: `gofmt`, `go vet`
- Packages benennen nach Ressource, nicht nach technischer Schicht
- Fehler immer als strukturiertes JSON zurueckgeben (stderr), nie panic()
- Context (context.Context) durch alle API-Calls durchreichen

### API-Client
- Zentraler HTTP-Client in `internal/api/client.go`
- Auth-Token wird automatisch aus Config/Env/Flag gelesen
- Base-URL konfigurierbar
- Einheitliches Error-Handling mit HTTP-Statuscode + API-Fehlermeldung

### Testing
- Unit-Tests mit `httptest.NewServer` fuer API-Mocks
- Testdaten als JSON-Fixtures in `testdata/`
- Keine Tests gegen echte API im CI

## API-Referenz

Die vollstaendige API-Spezifikation liegt in `swagger.json`.
- Host: telco-capture01.live.ml01.sipgate.net
- Base Path: /api/v3
- Auth: API-Key via `Auth-Token` Header
- Content-Type: application/json

## Implementierungsplan

Die Stories liegen im Verzeichnis `plan/`. Siehe [plan/README.md](plan/README.md) fuer Uebersicht, Abhaengigkeitsgraph und Reihenfolge.

### Phasen

**Phase 1 - Fundament** (sequenziell):
- Story 01: Projekt-Setup → Story 02: Config & Init → Story 03: HTTP-Client & Output
- Story 04: Model-Generierung (parallel zu 02/03)

**Phase 2 - Features** (parallelisierbar):
- Stories 05-12: Alle Endpoint-Gruppen, unabhaengig voneinander
- Empfohlen: 05 (Call Search) → 06 (Export) zuerst

**Phase 3 - Abschluss**:
- Story 13: Integration, Shell-Completion, Dokumentation

### Stories

| # | Story | Status |
|---|---|---|
| 01 | [Projekt-Setup](plan/01-project-setup.md) | done |
| 02 | [Config & Init](plan/02-config-and-init.md) | done |
| 03 | [HTTP-Client & Output](plan/03-api-client-and-output.md) | done |
| 04 | [Model-Generierung](plan/04-models-generation.md) | done |
| 05 | [Call Search & Transaction](plan/05-call-search-and-transaction.md) | done |
| 06 | [Call Export](plan/06-export.md) | done |
| 07 | [Recording & Interception](plan/07-recording-and-interception.md) | done |
| 08 | [User Management](plan/08-user-management.md) | done |
| 09 | [Netzwerk-Konfiguration](plan/09-network-config.md) | done |
| 10 | [Agent Management](plan/10-agent-management.md) | done |
| 11 | [Dashboard & Statistiken](plan/11-dashboard-and-statistics.md) | done |
| 12 | [Admin, Scripts & Rest](plan/12-admin-and-remaining.md) | done |
| 13 | [Integration & Polish](plan/13-integration-and-polish.md) | done |
