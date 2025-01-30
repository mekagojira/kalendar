package lib

import "os"

func GetEnv(keys ...string) string {
	for _, key := range keys {
		currentEnv := os.Getenv(key)
		if currentEnv != "" {
			return currentEnv
		}
	}

	if len(keys) > 1 {
		return keys[len(keys)-1]
	}
	return ""
}
