package kv

import "github.com/unicornultrafoundation/go-hashgraph/state"

// SaveGenesis saves the initial state (genesis) by calling the SaveState method.
func (s *Store) SaveGenesis(st *state.State) error {
	return s.SaveState(st)
}
