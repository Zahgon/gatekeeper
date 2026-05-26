package catalog

import (
	"context"
)

const (
	// CatalogFileName is the name of the cached catalog file.
	CatalogFileName = "catalog.yaml"
	// CatalogSourceFileName is the name of the cached catalog source URL file.
	CatalogSourceFileName = "catalog.source"
	// ConfigFileName is the name of the user config file (future use).
	ConfigFileName = "config.yaml"
)

// Cache manages local caching of the policy catalog.
type Cache struct {
	dir string
}

// NewCache creates a new Cache instance, creating the cache directory if needed.
func NewCache() (*Cache, error) { _ = "STUB: not implemented"; return nil, nil }

// Use os.UserConfigDir() for cross-platform support

// Fall back to home directory

// Dir returns the cache directory path.
func (c *Cache) Dir() string {
	_ = "STUB: not implemented"

	// CatalogPath returns the path to the cached catalog file.
	return ""
}

func (c *Cache) CatalogPath() string { _ = "STUB: not implemented"; return "" }

// CatalogSourcePath returns the path to the cached catalog source URL file.
func (c *Cache) CatalogSourcePath() string { _ = "STUB: not implemented"; return "" }

// SaveCatalog saves catalog data to the cache.
func (c *Cache) SaveCatalog(data []byte, sourceURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadCatalogData reads the cached catalog data.
func (c *Cache) LoadCatalogData() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadCatalog reads and parses the cached catalog.
func (c *Cache) LoadCatalog() (*PolicyCatalog, error) { _ = "STUB: not implemented"; return nil, nil }

// LoadCatalogSource reads the source URL used to cache the catalog.
func (c *Cache) LoadCatalogSource() (string, error) { _ = "STUB: not implemented"; return "", nil }

// LoadCatalogWithSource reads and parses the cached catalog and its source URL.
func (c *Cache) LoadCatalogWithSource() (*PolicyCatalog, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// CatalogExists checks if a cached catalog exists.
func (c *Cache) CatalogExists() bool { _ = "STUB: not implemented"; return false }

// GetCatalogURL returns the catalog URL from environment or the default.
func GetCatalogURL() string { _ = "STUB: not implemented"; return "" }

// EnsureCatalog loads the catalog from cache, or fetches it if not cached.
func EnsureCatalog(ctx context.Context, cache *Cache, fetcher Fetcher) (*PolicyCatalog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fetch and cache

// Save to cache (warn on failure, non-fatal)

// CatalogNotCachedError is returned when no cached catalog exists.
type CatalogNotCachedError struct{}

func (e *CatalogNotCachedError) Error() string { _ = "STUB: not implemented"; return "" }
