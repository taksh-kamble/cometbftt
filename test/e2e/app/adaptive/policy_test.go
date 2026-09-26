package adaptive

import "testing"

func TestPolicyReducesRetentionWhenDiskIsLow(t *testing.T) {
	policy := NewPolicy(PolicyConfig{
		MinRetainBlocks: 1000,
		MaxRetainBlocks: 10000,
		LowFreePercent:  20,
		HighFreePercent: 40,
		AdjustmentStep:  1000,
	})

	got := policy.Calculate(5000, 15)

	if got != 4000 {
		t.Fatalf("expected X=4000, got X=%d", got)
	}
}

func TestPolicyIncreasesRetentionWhenDiskIsAbundant(t *testing.T) {
	policy := NewPolicy(PolicyConfig{
		MinRetainBlocks: 1000,
		MaxRetainBlocks: 10000,
		LowFreePercent:  20,
		HighFreePercent: 40,
		AdjustmentStep:  1000,
	})

	got := policy.Calculate(5000, 50)

	if got != 6000 {
		t.Fatalf("expected X=6000, got X=%d", got)
	}
}

func TestPolicyDoesNotGoBelowMinimum(t *testing.T) {
	policy := NewPolicy(PolicyConfig{
		MinRetainBlocks: 1000,
		MaxRetainBlocks: 10000,
		LowFreePercent:  20,
		HighFreePercent: 40,
		AdjustmentStep:  1000,
	})

	got := policy.Calculate(1200, 10)

	if got != 1000 {
		t.Fatalf("expected X=1000, got X=%d", got)
	}
}

func TestPolicyDoesNotExceedMaximum(t *testing.T) {
	policy := NewPolicy(PolicyConfig{
		MinRetainBlocks: 1000,
		MaxRetainBlocks: 10000,
		LowFreePercent:  20,
		HighFreePercent: 40,
		AdjustmentStep:  1000,
	})

	got := policy.Calculate(9500, 50)

	if got != 10000 {
		t.Fatalf("expected X=10000, got X=%d", got)
	}
}

func TestPolicyKeepsCurrentRetentionInMiddleRange(t *testing.T) {
	policy := NewPolicy(PolicyConfig{
		MinRetainBlocks: 1000,
		MaxRetainBlocks: 10000,
		LowFreePercent:  20,
		HighFreePercent: 40,
		AdjustmentStep:  1000,
	})

	got := policy.Calculate(5000, 30)

	if got != 5000 {
		t.Fatalf("expected X=5000, got X=%d", got)
	}
}
