package server

import "sync"

var (
	mu       sync.RWMutex
	services map[string]interface{}
)

func init() {
	services = make(map[string]interface{})
}
