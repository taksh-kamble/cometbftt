package adaptive

// PolicyConfig defines the parameters used by the adaptive
// retention policy.
type PolicyConfig struct {
	MinRetainBlocks uint64
	MaxRetainBlocks uint64

	// Below this free-space percentage, retention is reduced.
	LowFreePercent float64

	// Above this free-space percentage, retention is increased.
	HighFreePercent float64

	// Number of blocks by which X changes at a time.
	AdjustmentStep uint64
}

// Policy decides how many recent blocks should remain in hot storage.
type Policy struct {
	config PolicyConfig
}

// NewPolicy creates an adaptive retention policy.
func NewPolicy(config PolicyConfig) *Policy {
	return &Policy{
		config: config,
	}
}

// Calculate determines the next retention window X.
//
// currentX is the current number of retained blocks.
// freePercent is the currently available filesystem percentage.
func (p *Policy) Calculate(currentX uint64, freePercent float64) uint64 {
	x := currentX

	if freePercent < p.config.LowFreePercent {
		if x > p.config.MinRetainBlocks {
			if x > p.config.MinRetainBlocks+p.config.AdjustmentStep {
				x -= p.config.AdjustmentStep
			} else {
				x = p.config.MinRetainBlocks
			}
		}
	}

	if freePercent > p.config.HighFreePercent {
		if x < p.config.MaxRetainBlocks {
			x += p.config.AdjustmentStep

			if x > p.config.MaxRetainBlocks {
				x = p.config.MaxRetainBlocks
			}
		}
	}

	return x
}
