package functions

import "os"

var (
	// Check if an AES string was provided or not
	key string = ""
)

// Get or Set an environment variable, depending on if it exists or not.
// Found here: https://stackoverflow.com/a/45978733
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func initialization() {
	// Check if an AES string was provided or not
	CheckForAesKey(
		GetEnv("ots_aeskeyfile", "rand"),
	)

}
