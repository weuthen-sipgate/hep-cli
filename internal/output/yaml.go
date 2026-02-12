package output

import (
	"io"

	"go.yaml.in/yaml/v3"
)

// YAMLFormatter outputs data as YAML.
type YAMLFormatter struct{}

func (f *YAMLFormatter) Format(w io.Writer, data interface{}) error {
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	defer enc.Close()
	return enc.Encode(data)
}
