package internal

import (
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
