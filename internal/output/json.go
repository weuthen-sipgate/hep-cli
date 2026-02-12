package output

import (
	"encoding/json"
	"io"
)

// JSONFormatter outputs data as pretty-printed JSON.
type JSONFormatter struct{}

func (f *JSONFormatter) Format(w io.Writer, data interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}
