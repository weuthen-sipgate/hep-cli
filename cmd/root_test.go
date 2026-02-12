package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCommand_Properties(t *testing.T) {
	if rootCmd.Use != "hepic" {
		t.Errorf("expected Use='hepic', got %q", rootCmd.Use)
	}
	if rootCmd.Short == "" {
		t.Error("root command should have a short description")
	}
	if rootCmd.Long == "" {
		t.Error("root command should have a long description")
	}
	if !rootCmd.SilenceUsage {
		t.Error("root command should silence usage on error")
	}
	if !rootCmd.SilenceErrors {
		t.Error("root command should silence errors (PrintError handles them)")
	}
}

func TestRootCommand_Groups(t *testing.T) {
	groups := rootCmd.Groups()

	expectedGroups := map[string]string{
		"call":       "Call Analysis:",
		"data":       "Data Management:",
		"config":     "Configuration:",
		"monitoring": "Monitoring & Statistics:",
		"admin":      "Administration:",
	}

	if len(groups) != len(expectedGroups) {
		t.Errorf("expected %d groups, got %d", len(expectedGroups), len(groups))
	}

	foundGroups := make(map[string]string)
	for _, g := range groups {
		foundGroups[g.ID] = g.Title
	}

	for id, expectedTitle := range expectedGroups {
		title, ok := foundGroups[id]
		if !ok {
			t.Errorf("missing group with ID %q", id)
			continue
		}
		if title != expectedTitle {
			t.Errorf("group %q: expected title %q, got %q", id, expectedTitle, title)
		}
	}
}

func TestRootCommand_PersistentFlags(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
	}{
		{"host flag", "host"},
		{"token flag", "token"},
		{"format flag", "format"},
		{"verbose flag", "verbose"},
		{"no-color flag", "no-color"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag := rootCmd.PersistentFlags().Lookup(tt.flagName)
			if flag == nil {
				t.Errorf("expected persistent flag %q to be registered", tt.flagName)
			}
		})
	}
}

func TestRootCommand_FormatFlagDefault(t *testing.T) {
	flag := rootCmd.PersistentFlags().Lookup("format")
	if flag == nil {
		t.Fatal("format flag not found")
	}
	if flag.DefValue != "json" {
		t.Errorf("expected format default to be 'json', got %q", flag.DefValue)
	}
}

func TestRootCommand_HelpOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()

	// Verify group titles appear in help output
	expectedSections := []string{
		"Call Analysis:",
		"Data Management:",
		"Configuration:",
		"Monitoring & Statistics:",
		"Administration:",
	}

	for _, section := range expectedSections {
		if !strings.Contains(output, section) {
			t.Errorf("expected help output to contain group %q, got:\n%s", section, output)
		}
	}
}

func TestRootCommand_SubcommandsExist(t *testing.T) {
	// Check that key subcommands are registered
	expectedCommands := []string{
		"completion",
		"version",
	}

	commands := rootCmd.Commands()
	commandNames := make(map[string]bool)
	for _, cmd := range commands {
		commandNames[cmd.Name()] = true
	}

	for _, name := range expectedCommands {
		if !commandNames[name] {
			t.Errorf("expected subcommand %q to be registered", name)
		}
	}
}

func TestRootCommand_SubcommandsHaveGroups(t *testing.T) {
	// Verify that subcommands are assigned to groups
	validGroupIDs := map[string]bool{
		"call":       true,
		"data":       true,
		"config":     true,
		"monitoring": true,
		"admin":      true,
		"":           true, // ungrouped commands like completion, version, init
	}

	for _, cmd := range rootCmd.Commands() {
		if !validGroupIDs[cmd.GroupID] {
			t.Errorf("command %q has invalid GroupID %q", cmd.Name(), cmd.GroupID)
		}
	}
}

func TestExecute_ReturnsNilOnSuccess(t *testing.T) {
	rootCmd.SetArgs([]string{"version"})
	if err := Execute(); err != nil {
		t.Errorf("expected nil error for successful command, got: %v", err)
	}
}

func TestExecute_ReturnsErrorOnFailure(t *testing.T) {
	rootCmd.SetArgs([]string{"nonexistent-command-xyz"})
	err := Execute()
	if err == nil {
		t.Error("expected error for unknown command")
	}
}

func TestRootCommand_CompletionIsSubcommand(t *testing.T) {
	var found *cobra.Command
	for _, cmd := range rootCmd.Commands() {
		if cmd.Name() == "completion" {
			found = cmd
			break
		}
	}
	if found == nil {
		t.Fatal("completion command not found as subcommand of root")
	}
	// Completion should not be in any group (it's a utility command)
	if found.GroupID != "" {
		t.Errorf("expected completion to have no group, got %q", found.GroupID)
	}
}
