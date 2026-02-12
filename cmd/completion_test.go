package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestCompletionBash_GeneratesOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	err := rootCmd.GenBashCompletion(buf)
	if err != nil {
		t.Fatalf("unexpected error generating bash completion: %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Fatal("expected bash completion output, got empty string")
	}
	if !strings.Contains(output, "hepic") {
		t.Errorf("expected bash completion to reference 'hepic'")
	}
	if !strings.Contains(output, "complete") {
		t.Errorf("expected bash completion to contain 'complete' directive")
	}
}

func TestCompletionZsh_GeneratesOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	err := rootCmd.GenZshCompletion(buf)
	if err != nil {
		t.Fatalf("unexpected error generating zsh completion: %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Fatal("expected zsh completion output, got empty string")
	}
	if !strings.Contains(output, "hepic") {
		t.Errorf("expected zsh completion to reference 'hepic'")
	}
}

func TestCompletionFish_GeneratesOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	err := rootCmd.GenFishCompletion(buf, true)
	if err != nil {
		t.Fatalf("unexpected error generating fish completion: %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Fatal("expected fish completion output, got empty string")
	}
	if !strings.Contains(output, "complete") {
		t.Errorf("expected fish completion to contain 'complete' command")
	}
	if !strings.Contains(output, "hepic") {
		t.Errorf("expected fish completion to reference 'hepic'")
	}
}

func TestCompletionCommand_InvalidShell(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"completion", "powershell"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error for invalid shell argument")
	}
}

func TestCompletionCommand_NoArgs(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"completion"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error when no shell argument is provided")
	}
}

func TestCompletionCommand_TooManyArgs(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"completion", "bash", "extra"})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error when too many arguments are provided")
	}
}

func TestCompletionCommand_ValidArgs(t *testing.T) {
	if len(completionCmd.ValidArgs) != 3 {
		t.Errorf("expected 3 valid args, got %d", len(completionCmd.ValidArgs))
	}

	expectedShells := map[string]bool{"bash": false, "zsh": false, "fish": false}
	for _, arg := range completionCmd.ValidArgs {
		if _, ok := expectedShells[arg]; !ok {
			t.Errorf("unexpected valid arg: %s", arg)
		}
		expectedShells[arg] = true
	}
	for shell, found := range expectedShells {
		if !found {
			t.Errorf("missing valid arg: %s", shell)
		}
	}
}

func TestCompletionCommand_Properties(t *testing.T) {
	if completionCmd.Use != "completion [bash|zsh|fish]" {
		t.Errorf("unexpected Use: %s", completionCmd.Use)
	}
	if completionCmd.Short == "" {
		t.Error("completion command should have a short description")
	}
	if completionCmd.Long == "" {
		t.Error("completion command should have a long description")
	}
	if !completionCmd.DisableFlagsInUseLine {
		t.Error("completion command should disable flags in usage line")
	}
}

func TestCompletionCommand_ArgsValidation(t *testing.T) {
	// Verify that cobra.ExactArgs(1) and cobra.OnlyValidArgs are enforced
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"bash is valid", []string{"completion", "bash"}, false},
		{"zsh is valid", []string{"completion", "zsh"}, false},
		{"fish is valid", []string{"completion", "fish"}, false},
		{"invalid shell", []string{"completion", "powershell"}, true},
		{"no args", []string{"completion"}, true},
		{"too many args", []string{"completion", "bash", "zsh"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)
			rootCmd.SetArgs(tt.args)

			err := rootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("args=%v: wantErr=%v, gotErr=%v (%v)", tt.args, tt.wantErr, err != nil, err)
			}
		})
	}
}
