// Command generate reads swagger.json and produces Go model structs.
//
// Usage:
//
//	go run ./tools/generate
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

type Swagger struct {
	Definitions map[string]Definition `json:"definitions"`
}

type Definition struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
	Items      *Property           `json:"items"`
}

type Property struct {
	Type        string    `json:"type"`
	Format      string    `json:"format"`
	Description string    `json:"description"`
	Ref         string    `json:"$ref"`
	Items       *Property `json:"items"`
	Properties  map[string]Property `json:"properties"`
}

func main() {
	data, err := os.ReadFile("swagger.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading swagger.json: %v\n", err)
		os.Exit(1)
	}

	var swagger Swagger
	if err := json.Unmarshal(data, &swagger); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing swagger.json: %v\n", err)
		os.Exit(1)
	}

	outDir := filepath.Join("internal", "models")

	var buf strings.Builder
	buf.WriteString("// Code generated from swagger.json — DO NOT EDIT.\n")
	buf.WriteString("// Regenerate with: go run ./tools/generate\n\n")
	buf.WriteString("package models\n\n")
	buf.WriteString("import \"encoding/json\"\n\n")
	buf.WriteString("// Ensure json import is used.\n")
	buf.WriteString("var _ json.RawMessage\n\n")

	// Sorted definition names for deterministic output
	names := make([]string, 0, len(swagger.Definitions))
	for name := range swagger.Definitions {
		names = append(names, name)
	}
	sort.Strings(names)

	// Skip definitions that map to primitive Go types or external types
	skip := map[string]bool{
		"Duration":     true,
		"File":         true,
		"file":         true,
		"IP":           true,
		"Int":          true,
		"JSONText":     true,
		"RemoteLabels": true,
		"RemoteValues": true,
	}

	for _, name := range names {
		def := swagger.Definitions[name]
		if skip[name] {
			continue
		}
		if def.Type != "object" && def.Type != "" {
			continue
		}
		if len(def.Properties) == 0 {
			continue
		}

		requiredSet := make(map[string]bool)
		for _, r := range def.Required {
			requiredSet[r] = true
		}

		goName := toGoName(name)
		buf.WriteString(fmt.Sprintf("// %s represents the %s schema.\n", goName, name))
		buf.WriteString(fmt.Sprintf("type %s struct {\n", goName))

		// Sort properties
		propNames := make([]string, 0, len(def.Properties))
		for pname := range def.Properties {
			propNames = append(propNames, pname)
		}
		sort.Strings(propNames)

		for _, pname := range propNames {
			prop := def.Properties[pname]
			goType := resolveType(prop)
			goFieldName := toGoName(pname)
			omit := ""
			if !requiredSet[pname] {
				omit = ",omitempty"
			}
			tag := fmt.Sprintf("`json:\"%s%s\"`", pname, omit)
			buf.WriteString(fmt.Sprintf("\t%s %s %s\n", goFieldName, goType, tag))
		}

		buf.WriteString("}\n\n")
	}

	outPath := filepath.Join(outDir, "models_generated.go")
	if err := os.WriteFile(outPath, []byte(buf.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", outPath, err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", outPath)
}

func resolveType(p Property) string {
	if p.Ref != "" {
		refName := p.Ref[strings.LastIndex(p.Ref, "/")+1:]
		switch refName {
		case "JSONText":
			return "json.RawMessage"
		case "IP":
			return "string"
		case "Int":
			return "int64"
		case "File", "file":
			return "string"
		case "Duration":
			return "int64"
		default:
			return toGoName(refName)
		}
	}

	switch p.Type {
	case "string":
		if p.Format == "date-time" {
			return "string"
		}
		return "string"
	case "integer":
		return mapIntFormat(p.Format)
	case "number":
		if p.Format == "double" {
			return "float64"
		}
		return "float64"
	case "boolean":
		return "bool"
	case "array":
		if p.Items != nil {
			itemType := resolveType(*p.Items)
			return "[]" + itemType
		}
		return "[]interface{}"
	case "object":
		if len(p.Properties) > 0 {
			return "map[string]interface{}"
		}
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

func mapIntFormat(format string) string {
	switch format {
	case "uint8":
		return "uint8"
	case "uint16":
		return "uint16"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "int8":
		return "int8"
	case "int16":
		return "int16"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	default:
		return "int"
	}
}

func toGoName(s string) string {
	// Handle special prefixes
	s = strings.TrimPrefix(s, "_")

	parts := splitOnSeparators(s)
	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Handle common acronyms
		upper := strings.ToUpper(part)
		switch upper {
		case "ID", "UUID", "IP", "URL", "HTTP", "API", "JSON", "SIP", "RTP", "RTCP", "QOS", "DB", "UI", "PCAP", "CSV", "DTMF", "HEP", "TTL", "GID":
			result.WriteString(upper)
		default:
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			result.WriteString(string(runes))
		}
	}
	return result.String()
}

func splitOnSeparators(s string) []string {
	// Split on underscores, dashes, and camelCase boundaries
	var parts []string
	var current strings.Builder

	runes := []rune(s)
	for i, r := range runes {
		if r == '_' || r == '-' {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			continue
		}
		if i > 0 && unicode.IsUpper(r) && !unicode.IsUpper(runes[i-1]) {
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
		}
		current.WriteRune(unicode.ToLower(r))
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}
	return parts
}
