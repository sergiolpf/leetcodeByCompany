package designhitcounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitCounter_Hit_EmptyCounter(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(1)

	assert.Equal(t, 1, hitCounter.hits[1%FIVE_MINUES])
	assert.Equal(t, 1, hitCounter.timestamps[1%FIVE_MINUES])
}

func TestHitCounter_Hit_ExistingTimestamp(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(1)
	hitCounter.Hit(1)

	assert.Equal(t, 2, hitCounter.hits[1%FIVE_MINUES])
	assert.Equal(t, 1, hitCounter.timestamps[1%FIVE_MINUES])
}

func TestHitCounter_Hit_NewTimestamp(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(1)
	hitCounter.Hit(2)

	assert.Equal(t, 1, hitCounter.hits[1%FIVE_MINUES])
	assert.Equal(t, 1, hitCounter.hits[2%FIVE_MINUES])
	assert.Equal(t, 1, hitCounter.timestamps[1%FIVE_MINUES])
	assert.Equal(t, 2, hitCounter.timestamps[2%FIVE_MINUES])
}

func TestHitCounter_GetHits_EmptyCounter(t *testing.T) {
	hitCounter := Constructor()
	count := hitCounter.GetHits(1)

	assert.Equal(t, 0, count)
}

func TestHitCounter_GetHits_TimestampOutOfWindow(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(1)
	count := hitCounter.GetHits(302)

	assert.Equal(t, 0, count)
}

func TestHitCounter_GetHits_TimestampWithinWindow(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(1)
	hitCounter.Hit(2)
	hitCounter.Hit(300)
	count := hitCounter.GetHits(301)

	assert.Equal(t, 2, count)
}

func TestHitCounter_GetHits_MultipleHitsWithinWindow(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(2)
	hitCounter.Hit(2)
	hitCounter.Hit(3)
	hitCounter.Hit(300)
	count := hitCounter.GetHits(301)

	assert.Equal(t, 4, count)
}

func TestHitCounter_GetHits_SomeHitsOutOfWindow(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(1)
	hitCounter.Hit(2)
	hitCounter.Hit(300)
	hitCounter.Hit(600)
	count := hitCounter.GetHits(601)

	assert.Equal(t, 1, count)
}

func TestHitCounter_GetHits_OverflowTimestamp(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(300)
	hitCounter.Hit(600)
	hitCounter.Hit(900)
	count := hitCounter.GetHits(901)

	assert.Equal(t, 1, count)
}

func TestHitCounter_GetHits_OverflowAndWithinWindow(t *testing.T) {
	hitCounter := Constructor()
	hitCounter.Hit(300)
	hitCounter.Hit(600)
	hitCounter.Hit(900)
	hitCounter.Hit(901)
	count := hitCounter.GetHits(902)

	assert.Equal(t, 2, count)
}
