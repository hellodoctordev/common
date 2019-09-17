package logging

import "log"

func Debug(message string) {
	log.Printf("[DEBUG] %s", message)
}

func Info(message string) {
	log.Printf("[INFO] %s", message)
}

func Warn(message string) {
	log.Printf("[WARN] %s", message)
}

func Error(message string) {
	log.Printf("[ERROR] %s", message)
}