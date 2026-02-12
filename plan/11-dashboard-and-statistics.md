# Story 11: Dashboard, Statistiken & Monitoring

## Beschreibung

Dashboard-Verwaltung, Statistik-Abfragen, Prometheus-Metriken und Grafana-Proxy.

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

### Dashboard
- `GET /dashboard/info` -- Dashboard-Liste
- `PUT /dashboard/store/{dashboardId}` -- Dashboard erstellen/aktualisieren
- `DELETE /dashboard/store/{dashboardId}` -- Dashboard loeschen

### Statistiken
- `GET /statistic/_db` -- DB-Statistiken
- `POST /statistic/_measurements/{dbid}` -- Measurements abfragen
- `POST /statistic/_metrics` -- Metriken abfragen
- `POST /statistic/_retentions` -- Retention Policies
- `POST /statistic/data` -- Statistische Daten abfragen

### Prometheus
- `POST /prometheus/data` -- Prometheus-Daten abfragen
- `POST /prometheus/value` -- Metrik-Werte
- `GET /prometheus/labels` -- Labels auflisten
- `GET /prometheus/label/{userlabel}` -- Label-Details

### Grafana Proxy
- `GET /proxy/grafana/dashboards/uid/{uid}` -- Dashboard abrufen
- `GET /proxy/grafana/folders` -- Ordner auflisten
- `GET /proxy/grafana/org` -- Org-Info
- `GET /proxy/grafana/search/{uid}` -- Dashboard suchen
- `GET /proxy/grafana/status` -- Grafana-Status

### Database
- `GET /database/node/list` -- DB-Nodes auflisten

## Aufgaben

- [x] `internal/dashboard/dashboard.go` -- Dashboard CRUD
- [x] `internal/statistic/statistic.go` -- Statistik-Abfragen
- [x] `internal/statistic/prometheus.go` -- Prometheus-Integration
- [x] `internal/statistic/grafana.go` -- Grafana-Proxy
- [x] `cmd/dashboard.go` -- `hepic dashboard list|get|update|delete`
- [x] `cmd/statistic.go` -- `hepic statistic db|data|metrics|measurements|retentions`
- [x] `cmd/prometheus.go` -- `hepic prometheus query|value|labels`
- [x] `cmd/grafana.go` -- `hepic grafana dashboard|folders|org|status`
- [x] `cmd/database.go` -- `hepic database nodes`

## Akzeptanzkriterien

- `hepic dashboard list` zeigt alle Dashboards
- `hepic statistic db` zeigt DB-Statistiken
- `hepic statistic data --from --to` gibt Zeitreihen-Daten zurueck
- `hepic prometheus labels` zeigt verfuegbare Metriken
- `hepic grafana status` zeigt Grafana-Verbindungsstatus
- `hepic database nodes` zeigt DB-Node-Informationen

## Definition of Done

- Alle Commands implementiert
- Unit-Tests fuer Dashboard-CRUD und Statistik-Abfragen
- `go vet ./...` ohne Fehler
