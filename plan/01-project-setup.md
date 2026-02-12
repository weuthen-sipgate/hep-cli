# Story 01: Projekt-Setup & Grundgeruest

## Beschreibung

Initialisierung des Go-Projekts mit Cobra/Viper, grundlegender Projektstruktur und einem funktionierenden `hepic version` Befehl als Smoke-Test.

## Aufgaben

- [x] `go mod init hepic-cli`
- [x] Abhaengigkeiten: cobra, viper einbinden
- [x] `main.go` mit Cobra Root-Command
- [x] `cmd/root.go` mit globalen Flags (--format, --host, --token, --verbose, --no-color)
- [x] `cmd/version.go` -- `hepic version` gibt Versionsnummer aus
- [x] Grundlegende Verzeichnisstruktur anlegen (internal/api, internal/config, internal/models, internal/output)
- [x] `.gitignore` fuer Go-Projekte

## Akzeptanzkriterien

- `go build` erzeugt ein lauffaehiges Binary
- `hepic --help` zeigt Hilfetext mit globalen Flags
- `hepic version` gibt Versionsinformation als JSON aus
- `hepic version --format table` gibt menschenlesbare Version aus
- Globale Flags (--host, --token, --format, --verbose) sind definiert und parsebar

## Definition of Done

- Code kompiliert ohne Fehler und Warnungen (`go vet ./...`)
- `go build` erzeugt Binary
- Unit-Test fuer Version-Command vorhanden und gruen
- CLAUDE.md ist aktuell
