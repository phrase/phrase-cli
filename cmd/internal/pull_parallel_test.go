package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildDownloadOpts_DefaultFileFormat(t *testing.T) {
	target := &Target{
		File:      "locales/<locale_name>.json",
		ProjectID: "proj1",
	}
	localeFile := &LocaleFile{
		FileFormat: "json",
		Tag:        "",
	}

	opts, err := target.buildDownloadOpts(localeFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.FileFormat.Value() != "json" {
		t.Errorf("expected file format 'json', got %q", opts.FileFormat.Value())
	}
}

func TestBuildDownloadOpts_TagHandling(t *testing.T) {
	target := &Target{
		File:      "locales/<locale_name>/<tag>.json",
		ProjectID: "proj1",
	}
	localeFile := &LocaleFile{
		FileFormat: "json",
		Tag:        "web",
	}

	opts, err := target.buildDownloadOpts(localeFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if opts.Tags.Value() != "web" {
		t.Errorf("expected tags 'web', got %q", opts.Tags.Value())
	}
	if opts.Tag.Value() != "" {
		t.Errorf("expected tag to be empty string, got %q", opts.Tag.Value())
	}
}

// A nil *os.File represents a 200 response with an empty body, which the API
// returns when a filter (e.g. --tags) matches no translations. copyToDestination
// must write an empty file for this case rather than failing, since it's a valid
// (if unusual) export result, not a failed download.
func TestCopyToDestination_NilFileWritesEmptyFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "en.json")

	if err := copyToDestination(nil, path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected file to be created at %s: %v", path, err)
	}
	if len(data) != 0 {
		t.Errorf("expected empty file, got %q", data)
	}
}

func TestCopyToDestination_NilFileTruncatesExistingContent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "en.json")
	if err := os.WriteFile(path, []byte(`{"hello":"world"}`), 0o644); err != nil {
		t.Fatalf("failed to seed existing file: %v", err)
	}

	if err := copyToDestination(nil, path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("expected previously downloaded content to be replaced with empty file, got %q", data)
	}
}

func TestCopyToDestination_CopiesFileContent(t *testing.T) {
	srcPath := filepath.Join(t.TempDir(), "src.json")
	if err := os.WriteFile(srcPath, []byte(`{"hello":"world"}`), 0o644); err != nil {
		t.Fatalf("failed to write source file: %v", err)
	}
	src, err := os.Open(srcPath)
	if err != nil {
		t.Fatalf("failed to open source file: %v", err)
	}
	defer src.Close()

	destPath := filepath.Join(t.TempDir(), "dest.json")
	if err := copyToDestination(src, destPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("failed to read dest file: %v", err)
	}
	if string(data) != `{"hello":"world"}` {
		t.Errorf("expected copied content, got %q", data)
	}
}
