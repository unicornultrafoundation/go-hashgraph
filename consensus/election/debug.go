package election

import (
	"crypto/sha256"
	"fmt"

	"github.com/unicornultrafoundation/go-hashgraph/consensus/helpers"
	"github.com/unicornultrafoundation/go-hashgraph/hash"
	"github.com/unicornultrafoundation/go-hashgraph/types"
)

// DebugStateHash may be used in tests to match election state
func (el *Election) DebugStateHash() hash.Hash {
	hasher := sha256.New()
	write := func(bb []byte) {
		if _, err := hasher.Write(bb); err != nil {
			panic(err)
		}
	}

	for vid, vote := range el.votes {
		write(vid.fromRoot.Hash[:])
		write(vid.fromRoot.Slot.Frame.Bytes())
		write(vid.fromRoot.Slot.Validator.Bytes())
		write(vote.observedRoot.Bytes())
	}
	for validator, vote := range el.decidedRoots {
		write(validator.Bytes())
		write(vote.observedRoot.Bytes())
	}
	return hash.FromBytes(hasher.Sum(nil))
}

// @param (optional) voters is roots to print votes for. May be nil
// @return election summary in a human readable format
func (el *Election) String(voters []types.RootAndSlot) string {
	if voters == nil {
		votersM := make(map[types.RootAndSlot]bool)
		for vid := range el.votes {
			votersM[vid.fromRoot] = true
		}
		for voter := range votersM {
			voters = append(voters, voter)
		}
	}
	activeIndices, _ := helpers.ActiveValidatorIndices(el.state, el.state.Epoch())
	info := "Every line contains votes from a root, for each subject. y is yes, n is no. Upper case means 'decided'. '-' means that subject was already decided when root was processed.\n"
	for _, root := range voters { // voter
		info += fmt.Sprintf("%s-%d: ", root.Hash.String(), root.Slot.Frame)
		for _, forV := range activeIndices { // subject
			vid := voteID{
				fromRoot:     root,
				forValidator: forV,
			}
			vote, ok := el.votes[vid]
			if !ok { // i.e. subject was decided when root processed
				info += "-"
				continue
			}
			if vote.yes {
				if vote.decided {
					info += "Y"
				} else {
					info += "y"
				}
			} else {
				if vote.decided {
					info += "N"
				} else {
					info += "n"
				}
			}
		}
		info += "\n"
	}
	return info
}
