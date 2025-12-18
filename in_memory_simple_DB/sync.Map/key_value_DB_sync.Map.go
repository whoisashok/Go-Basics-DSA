package main

import "sync"

type DB struct {
	data sync.Map
}

// New creates a new in-memory key-value database
func New() *DB {
	return &DB{}
}

// Set stores a value for a key
func (db *DB) Set(key string, value []byte) {
	db.data.Store(key, value)
}

// Get retrieves a value by key
// Returns nil if the key does not exist
func (db *DB) Get(key string) []byte {
	val, ok := db.data.Load(key)
	if !ok {
		return nil
	}
	return val.([]byte)
}

// Delete removes a key from the database
func (db *DB) Delete(key string) {
	db.data.Delete(key)
}

// Exists checks whether a key exists
func (db *DB) Exists(key string) bool {
	_, ok := db.data.Load(key)
	return ok
}

// Len returns number of key–value pairs
func (db *DB) Len() int {
	count := 0
	db.data.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}

// Keys returns all keys
func (db *DB) Keys() []string {
	keys := []string{}
	db.data.Range(func(k, _ any) bool {
		keys = append(keys, k.(string))
		return true
	})
	return keys
}

// Clear removes all entries
func (db *DB) Clear() {
	db.data.Range(func(k, _ any) bool {
		db.data.Delete(k)
		return true
	})
}

func main() {
	db := New()

	db.Set("name", []byte("golang"))
	db.Set("version", []byte("1.22"))

	println(string(db.Get("name"))) // golang
	println(db.Exists("version"))   // true
	println(db.Len())               // 2

	db.Delete("version")
	println(db.Exists("version")) // false
}
