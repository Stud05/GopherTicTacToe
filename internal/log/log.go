// Copyright (c) 2025 Elian Waeber & Valentin Roch
// SPDX-License-Identifier: Apache-2.0

// Package log provides the utility fonction to log messages with different levels of severity.
// The log level can be changed by modifying the CURRENT constant.
package log

import (
	"log"
)

const (
	CURRENT = INFO // Default log level
	DEBUG   = iota // Debug log level
	INFO           // Info log level
	WARN           // Warning log level
	ERROR          // Error log level
	FATAL          // Fatal log level (exits the program without executing defer statements)
)

// Filters log messages based on the log level
func LogMessage(level int, message string) {
	if level >= CURRENT {
		switch level {
		case DEBUG:
			log.Printf("[DEBUG] %s", message)
		case INFO:
			log.Printf("[INFO] %s", message)
		case WARN:
			log.Printf("[WARN] %s", message)
		case ERROR:
			log.Printf("[ERROR] %s", message)
		case FATAL:
			log.Fatalf("[FATAL] %s", message)
		}
	}
}
