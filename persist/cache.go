package persist

type Cache interface {
	// Set puts key and value into cache.
	// First parameter for extra should be uint denoting expected survival time.
	// If survival time equals 0 or less, the key will always be survival.
	Set(key string, value bool, extra ...interface{}) error

	// Get returns result for key,
	// If there's no such key existing in cache,
	// ErrNoSuchKey will be returned.
	Get(key string) (bool, error)

	// Delete will remove the specific key in cache.
	// If there's no such key existing in cache,
	// ErrNoSuchKey will be returned.
	Delete(key string) error

	// Clear deletes all the items stored in cache.
	Clear() error
}
