package code_challenge

type Database interface {
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
	return &db{}
}

type db struct {
	// TODO: Implement
	data map[string]int
	tx   map[string]int
}

// Set the key to given value
func (d *db) Set(key string, value int) {
	d.data[key] = value
}

// Get the value for the given key, set 'ok' to true if key exists
func (d *db) Get(key string) (value int, ok bool) {
	value, ok = d.data[key]
	return value, ok
}

// Unset the key, making it just like that key was never set
func (d *db) Unset(key string) {
	delete(d.data, key)
}

// Begin opens a new transaction
func (d *db) Begin() {
	// Copy current data to transaction map
	for k, v := range d.data {
		d.tx[k] = v
	}
}

// Commit closes all open transaction blocks, permanently apply the
// changes made in them.
func (d *db) Commit() error {
	for k, v := range d.tx {
		d.data[k] = v
	}
	d.tx = nil
	return nil
}

// Rollback undoes all of the commands issued in the most recent
// transaction block, and closes the block.
func (d *db) Rollback() error {
	for k := range d.tx {
		delete(d.data, k)
	}
	d.tx = nil
	return nil
}
