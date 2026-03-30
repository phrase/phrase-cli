package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-go/v4"
)

func TestLoadFromPath_Missing(t *testing.T) {
	dc := loadFromPath("/nonexistent/path/cache.json")
	if dc == nil {
		t.Fatal("expected non-nil cache")
	}
	if dc.Version != cacheVersion {
		t.Errorf("expected version %d, got %d", cacheVersion, dc.Version)
	}
	if len(dc.Entries) != 0 {
		t.Errorf("expected empty entries, got %d", len(dc.Entries))
	}
}

func TestLoadFromPath_Corrupt(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, cacheFileName)
	os.WriteFile(path, []byte("not json"), 0o600)

	dc := loadFromPath(path)
	if dc.Version != cacheVersion {
		t.Errorf("expected version %d, got %d", cacheVersion, dc.Version)
	}
	if len(dc.Entries) != 0 {
		t.Errorf("expected empty entries on corrupt file, got %d", len(dc.Entries))
	}
}

func TestLoadFromPath_VersionMismatch(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, cacheFileName)
	data, _ := json.Marshal(DownloadCache{
		Version: cacheVersion + 99,
		Entries: map[string]CacheEntry{"key": {ETag: "old"}},
	})
	os.WriteFile(path, data, 0o600)

	dc := loadFromPath(path)
	if dc.Version != cacheVersion {
		t.Errorf("expected fresh cache with version %d, got %d", cacheVersion, dc.Version)
	}
	if len(dc.Entries) != 0 {
		t.Errorf("expected empty entries on version mismatch, got %d", len(dc.Entries))
	}
}

func TestLoadFromPath_Valid(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, cacheFileName)
	data, _ := json.Marshal(DownloadCache{
		Version: cacheVersion,
		Entries: map[string]CacheEntry{
			"k1": {ETag: `"e1"`, LastModified: "mod1"},
		},
	})
	os.WriteFile(path, data, 0o600)

	dc := loadFromPath(path)
	if dc.Version != cacheVersion {
		t.Errorf("expected version %d, got %d", cacheVersion, dc.Version)
	}
	entry, ok := dc.Get("k1")
	if !ok {
		t.Fatal("expected entry k1 to exist")
	}
	if entry.ETag != `"e1"` || entry.LastModified != "mod1" {
		t.Errorf("unexpected entry: %+v", entry)
	}
}

func TestDownloadCache_RoundTrip(t *testing.T) {
	dc := &DownloadCache{
		Version: cacheVersion,
		Entries: make(map[string]CacheEntry),
	}

	dc.Set("abc", CacheEntry{ETag: `"etag123"`, LastModified: "Thu, 01 Jan 2025 00:00:00 GMT"})

	entry, ok := dc.Get("abc")
	if !ok {
		t.Fatal("expected entry to exist")
	}
	if entry.ETag != `"etag123"` {
		t.Errorf("expected ETag %q, got %q", `"etag123"`, entry.ETag)
	}
	if entry.LastModified != "Thu, 01 Jan 2025 00:00:00 GMT" {
		t.Errorf("unexpected LastModified: %s", entry.LastModified)
	}

	_, ok = dc.Get("missing")
	if ok {
		t.Error("expected missing key to return false")
	}
}

func TestDownloadCache_SaveLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", cacheFileName)

	dc := &DownloadCache{
		Version: cacheVersion,
		Entries: make(map[string]CacheEntry),
		path:    path,
	}
	dc.Set("key1", CacheEntry{ETag: `"e1"`, LastModified: "mod1"})
	dc.Set("key2", CacheEntry{ETag: `"e2"`})

	if err := dc.Save(); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded := loadFromPath(path)
	if loaded.Version != cacheVersion {
		t.Errorf("expected version %d, got %d", cacheVersion, loaded.Version)
	}
	if len(loaded.Entries) != 2 {
		t.Errorf("expected 2 entries, got %d", len(loaded.Entries))
	}
	if loaded.Entries["key1"].ETag != `"e1"` {
		t.Errorf("unexpected ETag for key1: %s", loaded.Entries["key1"].ETag)
	}
	if loaded.Entries["key1"].LastModified != "mod1" {
		t.Errorf("unexpected LastModified for key1: %s", loaded.Entries["key1"].LastModified)
	}
	if loaded.Entries["key2"].ETag != `"e2"` {
		t.Errorf("unexpected ETag for key2: %s", loaded.Entries["key2"].ETag)
	}
}

func TestDownloadCache_SaveSkipsWhenClean(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, cacheFileName)

	dc := &DownloadCache{
		Version: cacheVersion,
		Entries: make(map[string]CacheEntry),
		path:    path,
	}

	if err := dc.Save(); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	if _, err := os.Stat(path); err == nil {
		t.Error("expected no file to be written when cache is clean")
	}
}

func TestDownloadCache_ETagOnly(t *testing.T) {
	dc := &DownloadCache{
		Version: cacheVersion,
		Entries: make(map[string]CacheEntry),
	}
	dc.Set("k", CacheEntry{ETag: `"abc"`})

	entry, ok := dc.Get("k")
	if !ok {
		t.Fatal("expected entry")
	}
	if entry.ETag != `"abc"` {
		t.Errorf("unexpected ETag: %s", entry.ETag)
	}
	if entry.LastModified != "" {
		t.Errorf("expected empty LastModified, got: %s", entry.LastModified)
	}
}

func TestCacheKey_Deterministic(t *testing.T) {
	opts := phrase.LocaleDownloadOpts{
		FileFormat: optional.NewString("json"),
		Tags:       optional.NewString("web"),
	}

	k1 := CacheKey("proj1", "locale1", opts)
	k2 := CacheKey("proj1", "locale1", opts)

	if k1 != k2 {
		t.Errorf("expected deterministic keys, got %s and %s", k1, k2)
	}
}

func TestCacheKey_ZeroOpts(t *testing.T) {
	k := CacheKey("proj", "locale", phrase.LocaleDownloadOpts{})
	if k == "" {
		t.Error("expected non-empty key for zero opts")
	}
	if len(k) != 24 {
		t.Errorf("expected 24-char hex key, got %d chars: %s", len(k), k)
	}
}

func TestCacheKey_DifferentParams(t *testing.T) {
	opts1 := phrase.LocaleDownloadOpts{
		FileFormat: optional.NewString("json"),
	}
	opts2 := phrase.LocaleDownloadOpts{
		FileFormat: optional.NewString("yml"),
	}

	k1 := CacheKey("proj1", "locale1", opts1)
	k2 := CacheKey("proj1", "locale1", opts2)

	if k1 == k2 {
		t.Error("expected different keys for different file formats")
	}

	// Different project
	k3 := CacheKey("proj2", "locale1", opts1)
	if k1 == k3 {
		t.Error("expected different keys for different projects")
	}

	// Different locale
	k4 := CacheKey("proj1", "locale2", opts1)
	if k1 == k4 {
		t.Error("expected different keys for different locales")
	}
}

func TestCacheKey_IgnoresConditionalHeaders(t *testing.T) {
	base := phrase.LocaleDownloadOpts{
		FileFormat: optional.NewString("json"),
	}

	withETag := phrase.LocaleDownloadOpts{
		FileFormat:  optional.NewString("json"),
		IfNoneMatch: optional.NewString(`"abc123"`),
	}

	withLastMod := phrase.LocaleDownloadOpts{
		FileFormat:      optional.NewString("json"),
		IfModifiedSince: optional.NewString("Thu, 01 Jan 2025 00:00:00 GMT"),
	}

	kBase := CacheKey("p", "l", base)
	kETag := CacheKey("p", "l", withETag)
	kLastMod := CacheKey("p", "l", withLastMod)

	if kBase != kETag {
		t.Errorf("IfNoneMatch should not affect cache key: %s vs %s", kBase, kETag)
	}
	if kBase != kLastMod {
		t.Errorf("IfModifiedSince should not affect cache key: %s vs %s", kBase, kLastMod)
	}
}
