# Story 07: Call Recording & Interception

## Beschreibung

Verwaltung von Call-Aufnahmen (Suche, Abspielen, Download) und aktiven Interceptions (CRUD + RTP-Recording).

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

### Recordings
- `POST /call/recording/data` -- Recording-Daten abfragen
- `GET /call/recording/play/{uuid}` -- Recording abspielen
- `GET /call/recording/download/{type}/{uuid}` -- Recording herunterladen
- `GET /call/recording/info/{uuid}` -- Recording-Metadaten

### Interceptions
- `GET /interceptions` -- Aktive Interceptions auflisten
- `POST /interceptions` -- Interception erstellen
- `PUT /interceptions/{uuid}` -- Interception aktualisieren
- `DELETE /interceptions/{uuid}` -- Interception loeschen
- `POST /interception/add/rtprecord` -- RTP-Recording hinzufuegen

## Aufgaben

- [ ] `internal/recording/recording.go` -- Recording-Funktionen
- [ ] `internal/recording/interception.go` -- Interception-CRUD
- [ ] `cmd/recording.go` -- `hepic recording` Parent-Command
- [ ] `cmd/recording_search.go` -- `hepic recording search --from --to`
- [ ] `cmd/recording_info.go` -- `hepic recording info <uuid>`
- [ ] `cmd/recording_download.go` -- `hepic recording download <uuid> -o file.wav`
- [ ] `cmd/interception.go` -- `hepic interception` Parent-Command
- [ ] `cmd/interception_list.go` -- `hepic interception list`
- [ ] `cmd/interception_create.go` -- `hepic interception create --caller X --callee Y`
- [ ] `cmd/interception_update.go` -- `hepic interception update <uuid> [flags]`
- [ ] `cmd/interception_delete.go` -- `hepic interception delete <uuid>`

## Akzeptanzkriterien

- `hepic recording search --from 2025-01-01` listet Aufnahmen
- `hepic recording download <uuid> -o call.wav` speichert Audio-Datei
- `hepic interception list` zeigt aktive Interceptions
- `hepic interception create --caller "+49..." --callee "+49..."` erstellt Interception
- `hepic interception delete <uuid>` loescht mit Bestaetigung (--force fuer Skip)
- Alle CRUD-Operationen geben Ergebnis als JSON zurueck

## Definition of Done

- Alle Commands implementiert
- Unit-Tests fuer Recording-Search und Interception-CRUD
- Delete-Bestaetigung implementiert und getestet
- `go vet ./...` ohne Fehler
