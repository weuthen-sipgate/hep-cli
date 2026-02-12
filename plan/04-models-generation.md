# Story 04: Model-Generierung aus Swagger

## Beschreibung

Go-Structs fuer alle 97 Schema-Definitionen aus `swagger.json` generieren und in `internal/models/` einchecken. Dazu ein Generator-Script, das bei Bedarf erneut ausgefuehrt werden kann.

## Abhaengigkeiten

- Story 01 (Projekt-Setup)

## Aufgaben

- [x] Generator-Ansatz waehlen: openapi-generator, go-swagger, oder eigenes Script
- [x] Generator ausfuehren gegen `swagger.json`
- [x] Generierte Structs in `internal/models/` ablegen
- [x] Structs pruefen und ggf. manuell nachbessern (JSON-Tags, Typen)
- [x] `generate.go` oder `Makefile`-Target fuer Regenerierung dokumentieren
- [x] Sicherstellen, dass alle Structs korrekte `json:"..."` Tags haben

## Akzeptanzkriterien

- Fuer jede Schema-Definition in swagger.json existiert ein Go-Struct
- Alle Structs haben korrekte JSON-Tags
- `go build` kompiliert fehlerfrei mit den generierten Models
- Ein dokumentierter Weg existiert, die Models bei API-Aenderung neu zu generieren

## Definition of Done

- Generierte Models in `internal/models/` eingecheckt
- Kompilierung erfolgreich
- Generator-Aufruf in README oder Makefile dokumentiert
- `go vet ./...` ohne Fehler
