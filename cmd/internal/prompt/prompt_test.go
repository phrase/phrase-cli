package prompt

import (
	"bufio"
	"strings"
	"testing"
)

func TestLine(t *testing.T) {
	input := "  My Project Name  \n"
	reader := bufio.NewReader(strings.NewReader(input))
	oldStdin := stdin
	defer func() { stdin = oldStdin }()
	stdin = reader

	var result string
	err := Line("Prompt:", &result)
	if err != nil {
		t.Errorf("Line returned error: %v", err)
	}

	expected := "My Project Name"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestLine_Empty(t *testing.T) {
	input := "\n"
	reader := bufio.NewReader(strings.NewReader(input))
	oldStdin := stdin
	defer func() { stdin = oldStdin }()
	stdin = reader

	var result string
	err := Line("Prompt:", &result)
	if err == nil {
		t.Error("Line expected error for empty input, got nil")
	}
}

func TestWithDefault_Default(t *testing.T) {
	input := "\n"
	reader := bufio.NewReader(strings.NewReader(input))
	oldStdin := stdin
	defer func() { stdin = oldStdin }()
	stdin = reader

	var result string
	err := WithDefault("Prompt:", &result, "default")
	if err != nil {
		t.Errorf("WithDefault returned error: %v", err)
	}

	expected := "default"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestWithDefault_Input(t *testing.T) {
	input := "user input\n"
	reader := bufio.NewReader(strings.NewReader(input))
	oldStdin := stdin
	defer func() { stdin = oldStdin }()
	stdin = reader

	var result string
	err := WithDefault("Prompt:", &result, "default")
	if err != nil {
		t.Errorf("WithDefault returned error: %v", err)
	}

	expected := "user input"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
