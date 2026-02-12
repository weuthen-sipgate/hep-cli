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
