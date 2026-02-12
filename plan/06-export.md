# Story 06: Call Export

## Beschreibung

Export von Call-Daten in verschiedenen Formaten: PCAP, SIPp, Text, sowie Transaktions-Reports und Archive.

## Abhaengigkeiten

- Story 05 (Call Search)

## API-Endpoints

- `POST /export/call/data/pcap` -- Export als PCAP
- `POST /export/call/messages/pcap` -- Nachrichten als PCAP
- `POST /export/call/messages/sipp` -- Export als SIPp
- `POST /export/call/messages/text` -- Export als Text
- `POST /export/call/stenographer` -- Stenographer-Export
- `POST /export/call/transaction/report` -- Transaktions-Report
- `POST /export/call/transaction/link` -- Transaktions-Link
- `POST /export/call/transaction/archive` -- Transaktions-Archiv
- `GET /export/action/{type}` -- Action-Logs (active, hepicapp, logs, picserver, rtpagent)

## Aufgaben

- [ ] `internal/export/export.go` -- Export-Funktionen
- [ ] `cmd/export.go` -- `hepic export` Parent-Command
- [ ] `cmd/export_pcap.go` -- `hepic export pcap --call-id <id> -o file.pcap`
- [ ] `cmd/export_sipp.go` -- `hepic export sipp --call-id <id> -o file.xml`
- [ ] `cmd/export_text.go` -- `hepic export text --call-id <id>`
- [ ] `cmd/export_report.go` -- `hepic export report --call-id <id>`
- [ ] `cmd/export_archive.go` -- `hepic export archive --call-id <id> -o file.tar.gz`
- [ ] `cmd/export_action.go` -- `hepic export action <type>`
- [ ] Binaere Responses (PCAP) korrekt in Datei schreiben statt auf stdout

## Akzeptanzkriterien

- `hepic export pcap --call-id <id> -o capture.pcap` schreibt gueltige PCAP-Datei
- `hepic export text --call-id <id>` gibt SIP-Nachrichten als Text auf stdout aus
- `-o`/`--output` Flag fuer alle Datei-Exports vorhanden
- Ohne `-o` wird an stdout geschrieben (ausser bei Binaer-Formaten, da Fehler)
- `hepic export action logs` zeigt Action-Logs
- Fortschrittsanzeige bei grossen Exports (auf stderr)

## Definition of Done

- Alle Export-Commands implementiert
- Unit-Tests mit Mock-HTTP fuer PCAP und Text-Export
- Binaer-Handling getestet (PCAP-Datei wird korrekt geschrieben)
- `go vet ./...` ohne Fehler
