package main

import (
	"2fa-api/pkg/storage"
)

// app configuration
type App struct {
	Addr     string
	Database storage.Database
}
