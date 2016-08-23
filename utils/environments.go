package utils

import "os"

func EnvOrDefault (osEnv string, defaultVal string) string {
	//** Returns the environment variable or the specified **//
	//** Returns String, if different value desired, perform a cast **//
	envVariable := os.Getenv(osEnv)
	if envVariable == "" {
		// Environment variable not defined
		envVariable = defaultVal
	}
	return envVariable
}

