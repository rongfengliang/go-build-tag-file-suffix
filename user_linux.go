package main

import (
	"fmt"
)

// UserApp UserApp
type UserApp struct {
}

const (
	// NAME for name
	NAME string = "dalong linux"
)

// Login login
func (u UserApp) Login() {
	fmt.Println(NAME)
}
