package kv

import (
	pState "github.com/unicornultrafoundation/go-hashgraph/proto/u2u/state"
	"github.com/unicornultrafoundation/go-hashgraph/state"
)

// State retrieves the serialized state from the key-value store, deserializes it, and returns a *state.State instance.
func (s *Store) State() (*state.State, error) {
	stateBytes, err := s.db.Get(stateKey)
	if err != nil {
		return nil, err
	}

	protoState := &pState.State{}
	if err := protoState.Unmarshal(stateBytes); err != nil {
		return nil, err
	}
	st := state.FromProto(protoState)
	return st, nil
}

// SaveState serializes the provided *state.State instance to its protocol buffer representation and stores it in the key-value store.
func (s *Store) SaveState(st *state.State) error {
	pbState := st.ToProto()
	stateBytes, err := pbState.Marshal()
	if err != nil {
		return err
	}
	return s.db.Put(stateKey, stateBytes)
}
