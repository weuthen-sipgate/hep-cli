# Story 13: Integration, Polish & Dokumentation

## Beschreibung

Abschluss-Story: End-to-End-Tests gegen echte API, Shell-Completion, Hilfe-Texte, und finale Qualitaetssicherung.

## Abhaengigkeiten

- Alle vorherigen Stories

## Aufgaben

- [x] Shell-Completion generieren (bash, zsh, fish) via `hepic completion bash|zsh|fish`
- [x] Hilfe-Texte fuer alle Commands pruefen und vervollstaendigen
- [x] `hepic --help` zeigt uebersichtliche Kommando-Gruppen (5 Gruppen: Call Analysis, Data Management, Configuration, Monitoring & Statistics, Administration)
- [ ] Manueller End-to-End-Test aller Kern-Workflows gegen echte API (erfordert echte API-Instanz)
- [x] Error-Messages pruefen: sind sie hilfreich und konsistent?
  - Doppelte Error-Ausgabe (output.PrintError + root Execute) behoben
  - JSON-Escaping im Root-Error-Handler gefixt
  - context.Background() durch cmd.Context() ersetzt (Signal-Cancellation)
  - URL Path Encoding fuer alle User-Inputs hinzugefuegt
- [x] `go vet ./...` und `golangci-lint` ohne Fehler
- [x] Makefile mit Targets: build, test, lint, generate
- [x] CLAUDE.md final aktualisieren

## Akzeptanzkriterien

- `hepic completion bash` gibt gueltiges Bash-Completion-Script aus
- Alle Commands haben aussagekraeftige `--help` Texte mit Beispielen
- Kern-Workflow funktioniert End-to-End: init -> call search -> export pcap
- Keine Panic-Situationen bei unerwarteten API-Antworten
- `make build && make test && make lint` laeuft fehlerfrei durch

## Definition of Done

- Shell-Completion fuer bash/zsh/fish implementiert
- Alle Tests gruen
- Lint-frei
- CLAUDE.md aktuell
- Makefile vollstaendig
