package main

import (
	"crypto/md5"
	"fmt"
	"sync"
)


var (
	storage map[string]string
	lock    sync.Mutex
)

func init() {
	storage = make(map[string]string)
}

func Get(key string) string {
	return storage[key]
}

func Save(data string) string {
	lock.Lock()
	defer lock.Unlock()

	key := checksum(data)
	storage[key] = data
	return key
}

func checksum(data string) string {
	barr := []byte(data)
	return fmt.Sprintf("%x", md5.Sum(barr))
}

