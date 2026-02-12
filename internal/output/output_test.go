package output

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

type sampleItem struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Count  int    `json:"count"`
}

func TestFprintJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	item := sampleItem{Name: "test", Status: "active", Count: 42}

	err := Fprint(buf, "json", item)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed sampleItem
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	if parsed.Name != "test" {
		t.Errorf("expected name=test, got %s", parsed.Name)
	}
	if parsed.Count != 42 {
		t.Errorf("expected count=42, got %d", parsed.Count)
	}

	// Check pretty-printing
	if !strings.Contains(buf.String(), "\n") {
		t.Error("expected pretty-printed JSON with newlines")
	}
}

func TestFprintJSON_Slice(t *testing.T) {
	buf := new(bytes.Buffer)
	items := []sampleItem{
		{Name: "a", Status: "ok", Count: 1},
		{Name: "b", Status: "err", Count: 2},
	}

	err := Fprint(buf, "json", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed []sampleItem
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("output is not valid JSON array: %v", err)
	}
	if len(parsed) != 2 {
		t.Errorf("expected 2 items, got %d", len(parsed))
	}
}

func TestFprintYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	item := sampleItem{Name: "test", Status: "active", Count: 42}

	err := Fprint(buf, "yaml", item)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "name: test") {
		t.Errorf("expected YAML with name field, got:\n%s", output)
	}
	if !strings.Contains(output, "count: 42") {
		t.Errorf("expected YAML with count field, got:\n%s", output)
	}
}

func TestFprintTable_Slice(t *testing.T) {
	buf := new(bytes.Buffer)
	items := []sampleItem{
		{Name: "alpha", Status: "ok", Count: 1},
		{Name: "beta", Status: "error", Count: 99},
	}

	err := Fprint(buf, "table", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "name") {
		t.Errorf("expected table header 'name', got:\n%s", output)
	}
	if !strings.Contains(output, "alpha") {
		t.Errorf("expected row with 'alpha', got:\n%s", output)
	}
	if !strings.Contains(output, "beta") {
		t.Errorf("expected row with 'beta', got:\n%s", output)
	}
	// Check separator line
	if !strings.Contains(output, "---") {
		t.Errorf("expected separator line, got:\n%s", output)
	}
}

func TestFprintTable_SingleObject(t *testing.T) {
	buf := new(bytes.Buffer)
	item := sampleItem{Name: "test", Status: "active", Count: 42}

	err := Fprint(buf, "table", item)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "name") {
		t.Errorf("expected key 'name', got:\n%s", output)
	}
	if !strings.Contains(output, "test") {
		t.Errorf("expected value 'test', got:\n%s", output)
	}
}

func TestFprintTable_Empty(t *testing.T) {
	buf := new(bytes.Buffer)
	err := Fprint(buf, "table", []sampleItem{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPrintError(t *testing.T) {
	// Just ensure it doesn't panic
	PrintError(nil)
}

func TestOrderedKeys_SortsAlphabetically(t *testing.T) {
	m := map[string]interface{}{
		"zebra":    1,
		"apple":    2,
		"mango":    3,
		"banana":   4,
		"cherry":   5,
	}

	keys := orderedKeys(m)

	expected := []string{"apple", "banana", "cherry", "mango", "zebra"}
	if len(keys) != len(expected) {
		t.Fatalf("expected %d keys, got %d", len(expected), len(keys))
	}
	for i, key := range keys {
		if key != expected[i] {
			t.Errorf("key[%d]: expected %q, got %q", i, expected[i], key)
		}
	}
}

func TestOrderedKeys_EmptyMap(t *testing.T) {
	m := map[string]interface{}{}
	keys := orderedKeys(m)
	if len(keys) != 0 {
		t.Errorf("expected 0 keys for empty map, got %d", len(keys))
	}
}

func TestOrderedKeys_SingleKey(t *testing.T) {
	m := map[string]interface{}{"only": 1}
	keys := orderedKeys(m)
	if len(keys) != 1 {
		t.Fatalf("expected 1 key, got %d", len(keys))
	}
	if keys[0] != "only" {
		t.Errorf("expected key 'only', got %q", keys[0])
	}
}

func TestOrderedKeys_NumericPrefixes(t *testing.T) {
	// sort.Strings sorts lexicographically, so "10" comes before "2"
	m := map[string]interface{}{
		"10_field": 1,
		"2_field":  2,
		"1_field":  3,
	}

	keys := orderedKeys(m)
	expected := []string{"10_field", "1_field", "2_field"}
	if len(keys) != len(expected) {
		t.Fatalf("expected %d keys, got %d", len(expected), len(keys))
	}
	for i, key := range keys {
		if key != expected[i] {
			t.Errorf("key[%d]: expected %q, got %q (lexicographic sort)", i, expected[i], key)
		}
	}
}

func TestOrderedKeys_CaseSensitive(t *testing.T) {
	// sort.Strings is case-sensitive: uppercase letters come before lowercase
	m := map[string]interface{}{
		"Banana": 1,
		"apple":  2,
		"Cherry": 3,
	}

	keys := orderedKeys(m)
	expected := []string{"Banana", "Cherry", "apple"}
	if len(keys) != len(expected) {
		t.Fatalf("expected %d keys, got %d", len(expected), len(keys))
	}
	for i, key := range keys {
		if key != expected[i] {
			t.Errorf("key[%d]: expected %q, got %q (case-sensitive sort)", i, expected[i], key)
		}
	}
}

func TestFormatValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"nil value", nil, ""},
		{"string value", "hello", "hello"},
		{"int value", 42, "42"},
		{"float value", 3.14, "3.14"},
		{"bool value", true, "true"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatValue(tt.input)
			if result != tt.expected {
				t.Errorf("formatValue(%v) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTableFormatter_KeyValueSorted(t *testing.T) {
	// Verify that when rendering a single object as key-value pairs,
	// the keys appear in sorted order
	buf := new(bytes.Buffer)
	data := map[string]interface{}{
		"zebra": "last",
		"alpha": "first",
		"middle": "mid",
	}

	err := Fprint(buf, "table", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimSpace(output), "\n")

	if len(lines) != 3 {
		t.Fatalf("expected 3 lines for 3 keys, got %d: %q", len(lines), output)
	}

	// First line should start with "alpha" (sorted first)
	if !strings.HasPrefix(strings.TrimSpace(lines[0]), "alpha") {
		t.Errorf("expected first line to start with 'alpha', got: %q", lines[0])
	}
	// Second line should start with "middle"
	if !strings.HasPrefix(strings.TrimSpace(lines[1]), "middle") {
		t.Errorf("expected second line to start with 'middle', got: %q", lines[1])
	}
	// Third line should start with "zebra"
	if !strings.HasPrefix(strings.TrimSpace(lines[2]), "zebra") {
		t.Errorf("expected third line to start with 'zebra', got: %q", lines[2])
	}
}

func TestTableFormatter_ColumnsTruncated(t *testing.T) {
	buf := new(bytes.Buffer)
	longValue := strings.Repeat("x", 100)
	items := []map[string]interface{}{
		{"name": longValue},
	}

	err := Fprint(buf, "table", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	// The value should be truncated to 57 chars + "..."
	if !strings.Contains(output, "...") {
		t.Errorf("expected truncated value with '...' for long strings, got:\n%s", output)
	}
}

func TestDefaultFormatIsJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	err := Fprint(buf, "", map[string]string{"key": "val"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var parsed map[string]string
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("default format should be JSON: %v", err)
	}
}
