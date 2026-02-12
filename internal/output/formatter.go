package output

import "io"

// Formatter formats data for output to a writer.
type Formatter interface {
	Format(w io.Writer, data interface{}) error
}

// formatters maps format names to their Formatter implementation.
var formatters = map[string]Formatter{
	"json":  &JSONFormatter{},
	"table": &TableFormatter{},
	"yaml":  &YAMLFormatter{},
}

// GetFormatter returns the Formatter for the given format name.
// Falls back to JSON if the format is unknown.
func GetFormatter(format string) Formatter {
	if f, ok := formatters[format]; ok {
		return f
	}
	return formatters["json"]
}
