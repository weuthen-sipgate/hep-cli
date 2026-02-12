# Story 03: HTTP-Client & Output-Formatter

## Beschreibung

Zentraler API-Client fuer alle HTTP-Requests und ein Output-System, das JSON, Table und YAML unterstuetzt. Diese beiden Bausteine sind die Grundlage fuer alle weiteren Commands.

## Abhaengigkeiten

- Story 02 (Konfiguration)

## Aufgaben

### API-Client (`internal/api/`)
- [x] `client.go` -- Client-Struct mit BaseURL, Token, http.Client
- [x] `client.go` -- Methoden: Get, Post, Put, Delete mit JSON Body
- [x] `client.go` -- Auth-Token automatisch als `Auth-Token` Header setzen
- [x] `client.go` -- Einheitliches Error-Handling: HTTP-Status + API-Fehlermeldung parsen
- [x] `client.go` -- Context-Support (context.Context) fuer alle Requests
- [x] `client.go` -- Verbose-Modus: Request/Response loggen bei --verbose
- [x] `errors.go` -- APIError-Typ mit StatusCode, Message, Detail

### Output-Formatter (`internal/output/`)
- [x] `formatter.go` -- Formatter-Interface + Registry (GetFormatter)
- [x] `json.go` -- JSONFormatter (pretty-printed)
- [x] `table.go` -- TableFormatter (fuer Listen und einzelne Objekte)
- [x] `yaml.go` -- YAMLFormatter
- [x] `output.go` -- Print-Funktion die --format Flag auswertet

## Akzeptanzkriterien

- API-Client sendet Requests mit korrektem Auth-Token Header
- API-Fehler (4xx, 5xx) werden als strukturierter Fehler zurueckgegeben
- `--format json` gibt pretty-printed JSON auf stdout
- `--format table` gibt tabellarische Darstellung auf stdout
- `--format yaml` gibt YAML auf stdout
- Fehler gehen immer als JSON auf stderr (auch bei --format table)
- `--verbose` loggt Request-Method, URL, Status auf stderr

## Definition of Done

- API-Client mit Unit-Tests (httptest-basiert)
- Alle drei Formatter mit Unit-Tests
- Error-Handling getestet (Netzwerkfehler, 401, 404, 500)
- `go vet ./...` ohne Fehler
