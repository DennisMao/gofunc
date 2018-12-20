package concurrency

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testDataRawWith10Count = []string{
		"1232132ddeqwerqdsaa",
		"aa1232132ddeqwerqds",
		"12aa32132ddeqwerqds",
		"1232aa132ddeqwerqds",
		"1232132aaddeqwerqds",
	}
)

func TestSingle(t *testing.T) {
	assert.Equal(t, 10, Single(testDataRawWith10Count, "a"))
}

func TestConcurrencyByLock(t *testing.T) {
	assert.Equal(t, 10, ConcurrencyByLock(testDataRawWith10Count, "a"))
}

func TestConcurrencyByChannel(t *testing.T) {
	assert.Equal(t, 10, ConcurrencyByChannel(testDataRawWith10Count, "a"))
}

func TestConcurrencyByChannelWithTimeoutControl(t *testing.T) {
	cnt := ConcurrencyByChannelWithTimeoutControl(testDataRawWith10Count, "a", time.Minute*2) // 限制2分钟
	assert.Equal(t, 10, cnt)
}

func TestConcurrencyByChannelWithCancelControl(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())

	cnt := ConcurrencyByChannelWithCancelControl(ctx, testDataRawWith10Count, "a") // 限制2分钟
	assert.Equal(t, 10, cnt)
}

func TestConcurrencyByChannelWithErrorControl(t *testing.T) {
	cnt, err := ConcurrencyByChannelWithErrorControl(testDataRawWith10Count, "a")
	assert.Equal(t, "test_error", err.Error())
	assert.Equal(t, 0, cnt)
}

func TestSplitConcurrency(t *testing.T) {
	assert.Equal(t, 10, SplitConcurrency(testDataRawWith10Count, "a", 2))
}
