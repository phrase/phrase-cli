package internal

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/antihax/optional"
	"github.com/phrase/phrase-go/v4"
)

const (
	cacheVersion  = 1
	cacheFileName = "download_cache.json"
	cacheDirName  = "phrase"
)

type CacheEntry struct {
	ETag         string `json:"etag,omitempty"`
	LastModified string `json:"last_modified,omitempty"`
}

type DownloadCache struct {
	Version int                   `json:"version"`
	Entries map[string]CacheEntry `json:"entries"`
	path    string
	dirty   bool
}

func LoadDownloadCache() *DownloadCache {
	return loadFromPath(cachePath())
}

func loadFromPath(path string) *DownloadCache {
	dc := &DownloadCache{
		Version: cacheVersion,
		Entries: make(map[string]CacheEntry),
		path:    path,
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Warning: could not read download cache %s: %v\n", path, err)
		}
		return dc
	}

	var loaded DownloadCache
	if err := json.Unmarshal(data, &loaded); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: corrupt download cache %s, starting fresh\n", path)
		return dc
	}
	if loaded.Version != cacheVersion {
		return dc
	}
	loaded.path = path
	if loaded.Entries == nil {
		loaded.Entries = make(map[string]CacheEntry)
	}
	return &loaded
}

func (dc *DownloadCache) Get(key string) (CacheEntry, bool) {
	e, ok := dc.Entries[key]
	return e, ok
}

func (dc *DownloadCache) Set(key string, entry CacheEntry) {
	dc.Entries[key] = entry
	dc.dirty = true
}

func (dc *DownloadCache) Save() error {
	if !dc.dirty {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(dc.path), 0o700); err != nil {
		return err
	}
	data, err := json.Marshal(dc)
	if err != nil {
		return err
	}
	if err := os.WriteFile(dc.path, data, 0o600); err != nil {
		return err
	}
	dc.dirty = false
	return nil
}

// CacheKey builds a deterministic key by hashing the full download parameters.
// It uses reflection to extract actual values from optional fields since
// antihax/optional types don't serialize meaningfully via json.Marshal.
func CacheKey(projectID, localeID string, opts phrase.LocaleDownloadOpts) string {
	// Zero out conditional request fields so they don't affect the key.
	opts.IfNoneMatch = optional.String{}
	opts.IfModifiedSince = optional.String{}

	raw := fmt.Sprintf("%s/%s/%s", projectID, localeID, serializeOpts(opts))
	h := sha256.Sum256([]byte(raw))
	return fmt.Sprintf("%x", h[:12])
}

// serializeOpts extracts set values from optional fields into a deterministic map.
// It assumes all fields in LocaleDownloadOpts are either slices or antihax/optional
// types with IsSet()/Value() methods. Fields with other types are silently excluded.
func serializeOpts(opts phrase.LocaleDownloadOpts) string {
	v := reflect.ValueOf(opts)
	t := v.Type()
	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		name := t.Field(i).Name

		// Handle slices directly
		if field.Kind() == reflect.Slice {
			if field.Len() > 0 {
				m[name] = field.Interface()
			}
			continue
		}

		// For optional types, check IsSet and extract Value
		isSetMethod := field.MethodByName("IsSet")
		valueMethod := field.MethodByName("Value")
		if isSetMethod.IsValid() && valueMethod.IsValid() {
			results := isSetMethod.Call(nil)
			if len(results) > 0 && results[0].Bool() {
				m[name] = valueMethod.Call(nil)[0].Interface()
			}
		}
	}

	data, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("%v", m)
	}
	return string(data)
}

func cachePath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		dir = os.TempDir()
	}
	return filepath.Join(dir, cacheDirName, cacheFileName)
}
