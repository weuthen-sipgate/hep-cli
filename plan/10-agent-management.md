# Story 10: Agent Management

## Beschreibung

Verwaltung von Capture-Agents und deren Subscriptions.

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

- `GET /agent/subscribe` -- Agents auflisten
- `GET /agent/subscribe/{uuid}` -- Agent-Details
- `PUT /agent/subscribe/{uuid}` -- Agent aktualisieren
- `DELETE /agent/subscribe/{uuid}` -- Agent loeschen
- `POST /agent/search/{guid}/{type}` -- Agent suchen
- `GET /agent/type/{type}` -- Agents nach Typ
- `POST /agentsub/protocol` -- Agent-Subscription hinzufuegen

## Aufgaben

- [ ] `internal/agent/agent.go` -- Agent-Funktionen
- [ ] `cmd/agent.go` -- `hepic agent` Parent-Command
- [ ] `cmd/agent_list.go` -- `hepic agent list`
- [ ] `cmd/agent_get.go` -- `hepic agent get <uuid>`
- [ ] `cmd/agent_update.go` -- `hepic agent update <uuid> [flags]`
- [ ] `cmd/agent_delete.go` -- `hepic agent delete <uuid>`
- [ ] `cmd/agent_search.go` -- `hepic agent search --guid X --type Y`
- [ ] `cmd/agent_type.go` -- `hepic agent type <type>`

## Akzeptanzkriterien

- `hepic agent list` zeigt alle registrierten Agents
- `hepic agent get <uuid>` zeigt Agent-Details
- `hepic agent search --guid X --type Y` findet Agents
- `hepic agent type home` filtert nach Agent-Typ
- Alle CRUD-Operationen funktionieren

## Definition of Done

- Alle Commands implementiert
- Unit-Tests fuer List und Search
- `go vet ./...` ohne Fehler
