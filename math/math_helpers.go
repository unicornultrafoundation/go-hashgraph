package math

import "sync"

var (
	// Sensible guess for 500 000 validators
	cachedSquareRoot = struct {
		sync.Mutex
		squareRoot, balance uint64
	}{squareRoot: 126491106, balance: 15999999897103236}
)

func CachedSquareRoot(balance uint64) uint64 {
	if balance == 0 {
		return 0
	}
	cachedSquareRoot.Lock()
	defer cachedSquareRoot.Unlock()
	if balance == cachedSquareRoot.balance {
		return cachedSquareRoot.squareRoot
	}
	cachedSquareRoot.balance = balance
	val := balance / cachedSquareRoot.squareRoot
	for {
		cachedSquareRoot.squareRoot = (cachedSquareRoot.squareRoot + val) / 2
		val = balance / cachedSquareRoot.squareRoot
		if cachedSquareRoot.squareRoot <= val {
			return cachedSquareRoot.squareRoot
		}
	}
}
