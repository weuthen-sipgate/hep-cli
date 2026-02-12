# Story 05: Call Search & Transaction

## Beschreibung

Kernfunktionalitaet: SIP-Calls suchen, Transaktionsdetails abrufen und Call-Reports (DTMF, Log, QoS) anzeigen. Dies ist die meistgenutzte Funktionsgruppe.

## Abhaengigkeiten

- Story 03 (API-Client & Output)
- Story 04 (Models)

## API-Endpoints

- `POST /search/call/data` -- Call-Daten suchen
- `POST /search/call/message` -- Call-Nachrichten suchen
- `POST /search/call/decode/message` -- SIP-Nachrichten dekodieren
- `POST /call/transaction` -- Transaktionsdetails
- `POST /search/transaction/type` -- Nach Transaktionstyp suchen
- `POST /call/report/dtmf` -- DTMF-Report
- `POST /call/report/log` -- Call-Log-Report
- `POST /call/report/qos` -- QoS-Report
- `POST /search/remote/data` -- Remote-Suche

## Aufgaben

- [ ] `internal/call/search.go` -- Search-Funktionen
- [ ] `internal/call/transaction.go` -- Transaction-Funktionen
- [ ] `internal/call/report.go` -- Report-Funktionen (DTMF, Log, QoS)
- [ ] `cmd/call.go` -- `hepic call` Parent-Command
- [ ] `cmd/call_search.go` -- `hepic call search` mit Filtern (--from, --to, --caller, --callee, --call-id)
- [ ] `cmd/call_message.go` -- `hepic call message`
- [ ] `cmd/call_decode.go` -- `hepic call decode`
- [ ] `cmd/call_transaction.go` -- `hepic call transaction`
- [ ] `cmd/call_report.go` -- `hepic call report dtmf|log|qos`

## Akzeptanzkriterien

- `hepic call search --from 2025-01-01 --to 2025-01-31` gibt Suchergebnisse als JSON zurueck
- `hepic call search --caller "+49123"` filtert nach Anrufer
- `hepic call transaction --call-id <id>` zeigt Transaktionsdetails
- `hepic call report dtmf --call-id <id>` zeigt DTMF-Report
- `hepic call report qos --call-id <id>` zeigt QoS-Report
- Alle Befehle unterstuetzen --format json|table|yaml
- Fehlende Pflichtparameter erzeugen hilfreiche Fehlermeldung

## Definition of Done

- Alle Commands implementiert und manuell gegen API getestet
- Unit-Tests mit Mock-HTTP fuer Search und Transaction
- Table-Output zeigt sinnvolle Spalten (Caller, Callee, Duration, Status, Timestamp)
- `go vet ./...` ohne Fehler
