package kv

import "github.com/unicornultrafoundation/go-hashgraph/u2udb"

// Store defines a wrapper for a key-value store.
type Store struct {
	db u2udb.Store
}

// NewKVStore creates a new instance of Store with the provided u2udb.Store.
func NewKVStore(db u2udb.Store) *Store {
	return &Store{
		db: db,
	}
}

// ClearDB clears the key-value store by closing and dropping it.
func (s *Store) ClearDB() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	s.db.Drop()
	return nil
}

// Close closes the key-value store.
func (s *Store) Close() error {
	return s.db.Close()
}
