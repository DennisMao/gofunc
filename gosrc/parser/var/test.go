package main

import (
	"fmt"
)

type Hello struct {
	Content string `json:"content"`
}

// Main
func main() {
	_ = Hello{}
}
