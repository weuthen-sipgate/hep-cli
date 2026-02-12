package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
)

// Print writes data to stdout in the format specified by the --format flag.
func Print(data interface{}) error {
	return Fprint(os.Stdout, viper.GetString("format"), data)
}

// Fprint writes data to the given writer in the specified format.
func Fprint(w io.Writer, format string, data interface{}) error {
	return GetFormatter(format).Format(w, data)
}

// PrintError writes a structured error to stderr as JSON.
func PrintError(err error) {
	msg := "unknown error"
	if err != nil {
		msg = err.Error()
	}
	errObj := map[string]string{"error": msg}
	data, _ := json.Marshal(errObj)
	fmt.Fprintln(os.Stderr, string(data))
}
