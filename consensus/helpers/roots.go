package helpers

import (
	"github.com/unicornultrafoundation/go-hashgraph/native/idx"
	"github.com/unicornultrafoundation/go-hashgraph/state"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

func FrameRoots(st *state.State, frame idx.Frame) []types.RootAndSlot {
	frameRoots := st.Roots()
	filtered := make([]types.RootAndSlot, 0)
	for _, frameRoot := range frameRoots {
		if frameRoot.Slot.Frame == frame {
			filtered = append(filtered, *frameRoot)
		}
	}
	return filtered
}
