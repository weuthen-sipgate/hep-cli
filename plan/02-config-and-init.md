# Story 02: Konfiguration & `hepic init`

## Beschreibung

Implementierung der Konfigurationsverwaltung (Lesen/Schreiben von `~/.hepic/config.yaml`) und des `hepic init` Commands fuer die Erstkonfiguration.

## Abhaengigkeiten

- Story 01 (Projekt-Setup)

## Aufgaben

- [x] `internal/config/config.go` -- Config-Struct, Load/Save mit Viper
- [x] Config-Hierarchie: Config-Datei < Env-Variablen < CLI-Flags
- [x] Unterstuetzte Config-Werte: `host`, `token`, `format`
- [x] Env-Variablen: `HEPIC_HOST`, `HEPIC_TOKEN`, `HEPIC_FORMAT`
- [x] `cmd/init.go` -- Interaktiver Modus: fragt Host und Token ab
- [x] `cmd/init.go` -- Non-interaktiver Modus: `hepic init --host X --token Y`
- [x] Validierung: Pruefen ob Host erreichbar und Token gueltig (GET /version/api/info)
- [x] Config-Datei wird in `~/.hepic/config.yaml` geschrieben

## Akzeptanzkriterien

- `hepic init` fragt interaktiv nach Host und Token
- `hepic init --host https://example.com --token abc123` schreibt Config non-interaktiv
- Nach `init` existiert `~/.hepic/config.yaml` mit korrekten Werten
- Env-Variablen ueberschreiben Config-Datei-Werte
- CLI-Flags ueberschreiben Env-Variablen
- Bei ungueltigem Host/Token gibt `init` einen Fehler aus
- Ohne Konfiguration gibt jeder Befehl einen hilfreichen Fehler aus ("Run hepic init first")

## Definition of Done

- Config Load/Save funktioniert mit Unit-Tests
- Hierarchie (File < Env < Flag) ist getestet
- `hepic init` funktioniert interaktiv und non-interaktiv
- `go vet ./...` ohne Fehler
