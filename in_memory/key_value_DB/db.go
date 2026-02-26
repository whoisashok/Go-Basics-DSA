package in_memory_simple_DB

import "errors"

type DBInterface interface {
	// Set the key to given value
	Set(key string, value int)

	// Get the value for the given key, set 'ok' to true if key exists
	Get(key string) (value int, ok bool)

	// Unset the key, making it just like that key was never set
	Unset(key string)

	// Begin opens a new transaction
	Begin()

	// Commit closes all open transaction blocks, permanently apply the
	// changes made in them.
	Commit() error

	// Rollback undoes all of the commands issued in the most recent
	// transaction block, and closes the block.
	Rollback() error
}

// NewDatabase creates a new in-memory simple database
func NewDatabase() Database {
	return Database{
		data:         make(map[string]int),
		transactions: []map[string]*int{},
	}
}

// Database represents our in-memory store with transaction support.
type Database struct {
	// data represents the committed, global state.
	data map[string]int
	// transactions is a stack of maps. Each map stores changes for that transaction level.
	// We use *int to distinguish between "value is 0" and "key was UNSET" (nil).
	transactions []map[string]*int
}

// Set the key to given value
func (db *Database) Set(key string, value int) {
	if len(db.transactions) > 0 {
		// Because transactions are stored as a stack. The most recent transaction is always at the top,
		// which is the last element in the slice.
		// So len(db.transactions)-1 gives us the active transaction layer.
		top := db.transactions[len(db.transactions)-1]
		top[key] = &value
	} else {
		// Apply directly to base data
		db.data[key] = value
	}
}

// Get the value for the given key, set 'ok' to true if key exists
func (db *Database) Get(key string) (value int, ok bool) {
	// Search through transaction layers from top to bottom
	for i := len(db.transactions) - 1; i >= 0; i-- {
		if valPtr, exists := db.transactions[i][key]; exists {
			if valPtr == nil {
				return 0, false // Explicitly UNSET in this transaction
			}
			return *valPtr, true
		}
	}

	// Finally, check the base data
	val, exists := db.data[key]
	return val, exists
}

// Unset the key, making it just like that key was never set
func (db *Database) Unset(key string) {
	if len(db.transactions) > 0 {
		// Record a nil pointer to signify the key is deleted in this transaction
		db.transactions[len(db.transactions)-1][key] = nil
	} else {
		delete(db.data, key)
	}
}

// Begin opens a new transaction
func (db *Database) Begin() {
	// Copy current data to transaction map
	db.transactions = append(db.transactions, make(map[string]*int))
}

// Commit closes all open transaction blocks, permanently apply the
// changes made in them.
func (db *Database) Commit() error {
	if len(db.transactions) == 0 {
		return errors.New("NO TRANSACTION")
	}

	// Merge all layers from oldest to newest into the base data
	for _, layer := range db.transactions {
		for k, v := range layer {
			if v == nil {
				delete(db.data, k)
			} else {
				db.data[k] = *v
			}
		}
	}

	// Clear the transaction stack
	db.transactions = []map[string]*int{}
	return nil
}

// Rollback undoes all of the commands issued in the most recent
// transaction block, and closes the block.
func (db *Database) Rollback() error {
	if len(db.transactions) == 0 {
		return errors.New("NO TRANSACTION")
	}
	db.transactions = db.transactions[:len(db.transactions)-1]
	return nil
}
