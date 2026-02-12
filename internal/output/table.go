package output

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

// TableFormatter outputs data as aligned text columns (lists) or key-value pairs (single objects).
type TableFormatter struct{}

func (f *TableFormatter) Format(w io.Writer, data interface{}) error {
	// Convert to JSON first to get a uniform map/slice representation.
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Try as array first
	var items []map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &items); err == nil && len(items) > 0 {
		return renderTable(w, items)
	}

	// Try as single object
	var item map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &item); err == nil {
		return renderKeyValue(w, item)
	}

	// Fallback to JSON
	return (&JSONFormatter{}).Format(w, data)
}

func renderTable(w io.Writer, items []map[string]interface{}) error {
	if len(items) == 0 {
		return nil
	}

	// Collect keys from first item to determine columns
	keys := orderedKeys(items[0])
	if len(keys) == 0 {
		return nil
	}

	// Calculate column widths
	widths := make([]int, len(keys))
	for i, k := range keys {
		widths[i] = len(k)
	}
	for _, item := range items {
		for i, k := range keys {
			val := formatValue(item[k])
			if len(val) > widths[i] {
				widths[i] = len(val)
			}
		}
	}

	// Cap column widths
	for i := range widths {
		if widths[i] > 60 {
			widths[i] = 60
		}
	}

	// Print header
	for i, k := range keys {
		fmt.Fprintf(w, "%-*s", widths[i]+2, k)
	}
	fmt.Fprintln(w)

	// Print separator
	for i := range keys {
		for j := 0; j < widths[i]; j++ {
			fmt.Fprint(w, "-")
		}
		fmt.Fprint(w, "  ")
	}
	fmt.Fprintln(w)

	// Print rows
	for _, item := range items {
		for i, k := range keys {
			val := formatValue(item[k])
			if len(val) > 60 {
				val = val[:57] + "..."
			}
			fmt.Fprintf(w, "%-*s", widths[i]+2, val)
		}
		fmt.Fprintln(w)
	}
	return nil
}

func renderKeyValue(w io.Writer, item map[string]interface{}) error {
	keys := orderedKeys(item)
	maxKeyLen := 0
	for _, k := range keys {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}
	for _, k := range keys {
		fmt.Fprintf(w, "%-*s  %s\n", maxKeyLen, k, formatValue(item[k]))
	}
	return nil
}

func orderedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func formatValue(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}
