# Story 08: User Management & Auth Tokens

## Beschreibung

Benutzerverwaltung (CRUD, Import/Export, Gruppen) und API-Token-Verwaltung.

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

### Users
- `GET /users` -- Alle User auflisten
- `POST /users` -- User erstellen
- `PUT /users/{uuid}` -- User aktualisieren
- `DELETE /users/{uuid}` -- User loeschen
- `PUT /users/update/password/{uuid}` -- Passwort aendern
- `POST /users/import` -- User aus CSV importieren
- `POST /users/import/h2` -- User aus H2-DB importieren
- `GET /users/export` -- User als CSV exportieren
- `GET /users/groups` -- Gruppen auflisten

### Auth/Tokens
- `GET /auth/type/list` -- Auth-Typen auflisten
- `POST /token/auth` -- Auth-Token erstellen
- `GET /token/auth/{uuid}` -- Token abrufen
- `DELETE /token/auth/{uuid}` -- Token loeschen

## Aufgaben

- [ ] `internal/user/user.go` -- User-CRUD-Funktionen
- [ ] `internal/user/token.go` -- Token-Funktionen
- [ ] `cmd/user.go` -- `hepic user` Parent-Command
- [ ] `cmd/user_list.go` -- `hepic user list`
- [ ] `cmd/user_create.go` -- `hepic user create --name X --email Y --password Z`
- [ ] `cmd/user_update.go` -- `hepic user update <uuid> [flags]`
- [ ] `cmd/user_delete.go` -- `hepic user delete <uuid>`
- [ ] `cmd/user_password.go` -- `hepic user password <uuid>`
- [ ] `cmd/user_import.go` -- `hepic user import --file users.csv`
- [ ] `cmd/user_export.go` -- `hepic user export -o users.csv`
- [ ] `cmd/user_groups.go` -- `hepic user groups`
- [ ] `cmd/token.go` -- `hepic token` Parent-Command
- [ ] `cmd/token_create.go` -- `hepic token create --name X`
- [ ] `cmd/token_list.go` -- `hepic token list`
- [ ] `cmd/token_delete.go` -- `hepic token delete <uuid>`

## Akzeptanzkriterien

- `hepic user list` zeigt alle Benutzer
- `hepic user create --name admin --email a@b.c --password X` erstellt User
- `hepic user delete <uuid>` loescht mit Bestaetigung
- `hepic user import --file users.csv` importiert aus CSV
- `hepic user export -o users.csv` exportiert als CSV
- `hepic token create --name "ci-token"` erstellt API-Token und zeigt es an
- Passwort-Felder werden nie in der Ausgabe angezeigt

## Definition of Done

- Alle Commands implementiert
- Unit-Tests fuer User-CRUD und Token-Management
- Sensitive Daten (Passwoerter) werden gefiltert
- `go vet ./...` ohne Fehler
