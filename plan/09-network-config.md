# Story 09: Netzwerk-Konfiguration (IP-Alias, HEPSub, Mapping, Protocol)

## Beschreibung

Verwaltung der Netzwerk-Konfigurationsressourcen: IP-Aliase, HEP-Subscriptions, Protocol-Mappings und Protocol-Definitionen.

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

### IP-Alias
- `GET /ipalias` -- Aliase auflisten
- `POST /ipalias` -- Alias erstellen
- `PUT /ipalias/{uuid}` -- Alias aktualisieren
- `DELETE /ipalias/{uuid}` -- Alias loeschen
- `DELETE /ipalias/all` -- Alle Aliase loeschen
- `GET /ipalias/export` -- Export als CSV
- `POST /ipalias/import` -- Import aus CSV

### HEP-Subscription
- `GET /hepsub/protocol` -- Subscriptions auflisten
- `POST /hepsub/protocol` -- Subscription erstellen
- `PUT /hepsub/protocol/{uuid}` -- Subscription aktualisieren
- `DELETE /hepsub/protocol/{uuid}` -- Subscription loeschen
- `POST /hepsub/search` -- HEP-Daten suchen

### Mapping
- `GET /mapping/protocol` -- Mappings auflisten
- `POST /mapping/protocol` -- Mapping erstellen
- `PUT /mapping/protocol/{uuid}` -- Mapping aktualisieren
- `DELETE /mapping/protocol/{uuid}` -- Mapping loeschen
- `GET /mapping/protocols` -- Alle Protocol-Mappings
- `GET /mapping/protocol/reset` -- Alle Mappings zuruecksetzen
- `GET /mapping/protocol/reset/{uuid}` -- Einzelnes Mapping zuruecksetzen

### Protocol
- `GET /protocol/search/{id}` -- Protocol suchen
- `POST /protocol/{id}` -- Protocol erstellen
- `PUT /protocol/{uuid}` -- Protocol aktualisieren
- `DELETE /protocol/{uuid}` -- Protocol loeschen

### Lookup IP
- `PUT /lookupip/{uuid}` -- IP-Lookup aktualisieren

## Aufgaben

- [x] `internal/config_resources/ipalias.go` -- IP-Alias CRUD + Import/Export
- [x] `internal/config_resources/hepsub.go` -- HEPSub CRUD + Search
- [x] `internal/config_resources/mapping.go` -- Mapping CRUD + Reset
- [x] `internal/config_resources/protocol.go` -- Protocol CRUD
- [x] `cmd/ipalias.go` -- `hepic ipalias list|create|update|delete|import|export`
- [x] `cmd/hepsub.go` -- `hepic hepsub list|create|update|delete|search`
- [x] `cmd/mapping.go` -- `hepic mapping list|create|update|delete|reset`
- [x] `cmd/protocol.go` -- `hepic protocol search|create|update|delete`

## Akzeptanzkriterien

- Alle CRUD-Operationen fuer jede Ressource funktionieren
- `hepic ipalias import --file aliases.csv` importiert aus CSV
- `hepic ipalias export -o aliases.csv` exportiert als CSV
- `hepic mapping reset` setzt alle Mappings zurueck (mit Bestaetigung)
- `hepic ipalias delete --all` loescht alle (mit Bestaetigung)
- Destruktive Operationen (delete all, reset) erfordern --force oder Bestaetigung

## Definition of Done

- Alle CRUD-Commands fuer alle vier Ressourcen implementiert
- Import/Export fuer IP-Alias getestet
- Bestaetigungsdialog fuer destruktive Ops getestet
- `go vet ./...` ohne Fehler
