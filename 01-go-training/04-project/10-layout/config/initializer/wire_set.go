package initializer

import "github.com/google/wire"

// Set Set
var Set = wire.NewSet(
	NewGorm,
)
