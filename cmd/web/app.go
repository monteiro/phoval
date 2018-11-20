package main

import (
	"phoval/pkg/storage"
)

// app configuration
type App struct {
	Addr     string
	Database storage.Database
}
