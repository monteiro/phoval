package main

import "phoval/service/notification"

// App configuration
type App struct {
	Addr string
	// notifier to send SMS or Emails
	Notifier notification.Notifier
}
