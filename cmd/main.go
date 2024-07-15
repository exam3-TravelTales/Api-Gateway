package main

import (
	"api/api/handler"
	"api/config"
	"fmt"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)
	// hand := NewHandler()

}

func NewHandler() *handler.Handler {
	return &handler.Handler{}
}
