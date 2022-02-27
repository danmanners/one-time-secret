package functions

import "os"

// Get or Set an environment variable, depending on if it exists or not.
// Found here: https://stackoverflow.com/a/45978733
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
