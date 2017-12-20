// Package env provides helper functions that simplify loading values from
// environment variables.
package env

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Csv returns the comma-separated values from the supplied environment
// variable.
func Csv(name string) []string {
	v := os.Getenv(name)
	return strings.Split(v, ",")
}

// CsvInt64s returns the int64 representations of the comma separate values that
// are extracted from the supplied environment variable.
func CsvInt64s(name string) []int64 {
	ints := []int64{}
	ss := Csv(name)
	for _, s := range ss {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Println("error while parsing", name)
			log.Panic(err)
		}
		ints = append(ints, v)
	}
	return ints
}

// Float64 returns the float64 representation of the supplied environment
// variable.
func Float64(name string) float64 {
	v, err := strconv.ParseFloat(os.Getenv(name), 64)
	if err != nil {
		log.Println("error while parsing", name)
		log.Panic(err)
	}
	return v
}

// Int64 returns the int64 representation of the supplied environment variable.
func Int64(name string) int64 {
	v, err := strconv.ParseInt(os.Getenv(name), 10, 64)
	if err != nil {
		log.Println("error while parsing", name)
		log.Panic(err)
	}
	return v
}

// Duration returns the time.Duration representation of th esupplied environment
// variable.
func Duration(name string) time.Duration {
	v, err := time.ParseDuration(os.Getenv(name))
	if err != nil {
		log.Println("error while parsing", name)
		log.Panic(err)
	}
	return v
}
