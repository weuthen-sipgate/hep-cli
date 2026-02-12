# Story 12: Admin, Scripts, Settings & verbleibende Endpoints

## Beschreibung

Verbleibende Endpoint-Gruppen: Admin, Scripts, User-Settings, Advanced Settings, Sharing, Data Import, ClickHouse und Troubleshooting.

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

### Admin
- `GET /admin/profiles` -- Admin-Profile
- `GET /configdb/tables/list` -- Config-DB Tabellen
- `POST /configdb/tables/resync` -- Config-DB resynchronisieren
- `GET /version/api/info` -- API-Version (bereits in Story 01 teilweise)
- `GET /version/ui/info` -- UI-Version

### Scripts
- `GET /script` -- Scripts auflisten
- `POST /script` -- Script erstellen
- `PUT /script/{uuid}` -- Script aktualisieren
- `DELETE /script/{uuid}` -- Script loeschen

### User Settings
- `GET /user/settings` -- Settings auflisten
- `POST /user/settings` -- Setting erstellen
- `GET /user/settings/{category}` -- Settings nach Kategorie
- `PUT /user/settings/{uuid}` -- Setting aktualisieren
- `DELETE /user/settings/{uuid}` -- Setting loeschen

### Advanced Settings
- `GET /advanced` -- Advanced Settings auflisten
- `POST /advanced` -- Setting erstellen
- `GET /advanced/{uuid}` -- Setting abrufen
- `PUT /advanced/{uuid}` -- Setting aktualisieren
- `DELETE /advanced/{uuid}` -- Setting loeschen

### Sharing
- `POST /share/call/report/dtmf/{uuid}` -- DTMF-Report teilen
- `POST /share/call/report/log/{uuid}` -- Call-Log teilen
- `POST /share/call/transaction/{uuid}` -- Transaktion teilen
- `POST /share/export/call/messages/pcap/{uuid}` -- PCAP teilen
- `POST /share/export/call/messages/text/{uuid}` -- Text teilen
- `GET /share/ipalias/{uuid}` -- IP-Alias teilen
- `GET /share/mapping/protocol/{uuid}` -- Mapping teilen

### Data Import
- `POST /import/data/pcap` -- PCAP importieren
- `POST /import/data/pcap/now` -- PCAP sofort importieren

### ClickHouse
- `POST /clickhouse/query/raw` -- Raw ClickHouse Query

### Troubleshooting
- `GET /troubleshooting/log/:type/:action` -- Troubleshooting-Logs

## Aufgaben

- [ ] `internal/admin/admin.go` -- Admin-Funktionen
- [ ] `internal/script/script.go` -- Script CRUD
- [ ] `cmd/admin.go` -- `hepic admin profiles|configdb|version`
- [ ] `cmd/script.go` -- `hepic script list|create|update|delete`
- [ ] `cmd/settings.go` -- `hepic settings list|create|update|delete`
- [ ] `cmd/advanced.go` -- `hepic advanced list|get|create|update|delete`
- [ ] `cmd/share.go` -- `hepic share report|transaction|pcap|text|ipalias|mapping`
- [ ] `cmd/import.go` -- `hepic import pcap --file capture.pcap`
- [ ] `cmd/clickhouse.go` -- `hepic clickhouse query "SELECT ..."`
- [ ] `cmd/troubleshooting.go` -- `hepic troubleshooting log <type> <action>`

## Akzeptanzkriterien

- Alle CRUD-Operationen fuer Scripts, Settings, Advanced Settings funktionieren
- `hepic admin version` zeigt API- und UI-Version
- `hepic admin configdb resync` synchronisiert Config-DB (mit Bestaetigung)
- `hepic import pcap --file capture.pcap` importiert PCAP-Datei
- `hepic clickhouse query "SELECT ..."` fuehrt Query aus und gibt Ergebnis als JSON
- `hepic share transaction <uuid>` gibt Share-Link zurueck
- `hepic troubleshooting log <type> <action>` zeigt Logs

## Definition of Done

- Alle Commands implementiert
- Unit-Tests fuer Script-CRUD und PCAP-Import
- Destruktive Ops (resync, delete) mit Bestaetigung
- `go vet ./...` ohne Fehler
