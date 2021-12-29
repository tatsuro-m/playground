package util

import "os"

func IsProd() bool {
	e := os.Getenv("APP_ENV")

	if e == "prod" {
		return true
	} else {
		return false
	}
}

func IsDev() bool {
	e := os.Getenv("APP_ENV")

	if e == "dev" {
		return true
	} else {
		return false
	}
}

func IsTest() bool {
	e := os.Getenv("APP_ENV")

	if e == "test" {
		return true
	} else {
		return false
	}
}
