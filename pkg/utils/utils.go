package utils

import "os"

//TODO : move ths to the GOStrings Module
func GetEnv(defaultString string, envName string) string {
	if value := os.Getenv(envName); value != "" {
		return value
	}
	return defaultString
}
