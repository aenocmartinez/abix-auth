package dao

import (
	"sync"
)

var lock = &sync.Mutex{}

type connectDB struct {
}

var instance *connectDB

func Instance() *connectDB {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &connectDB{}
		}
	}
	return instance
}
