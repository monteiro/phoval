package main

import (
	"phoval/pkg/storage"
	"phoval/service/notification"
)

// app configuration
type App struct {
	Addr     string
	Database storage.Database
	// notifier to send SMS or Emails
	Notifier notification.Notifier
}
