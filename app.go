// +build darwin

package main

import (
	"log"
)

// App App
type App struct {
}

// Log log
func (a App) Log() {
	log.Println("log darwin")
}
