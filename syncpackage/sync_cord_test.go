package syncpackage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDonationUpdate(t *testing.T) {
	chanData := make(chan int)

	go ListenerDonationUpdate(10, chanData)
	go ListenerDonationUpdate(15, chanData)
	// Updater goroutine
	go func() {
		defer close(chanData)

		for i := 0; i < 20; i++ {
			time.Sleep(time.Second)
			donation.cond.L.Lock()
			donation.balance++
			donation.cond.L.Unlock()
			donation.cond.Broadcast()
		}
	}()

	var sum int
	res := []int(nil)
	for data := range chanData {
		res = append(res, data)
		sum += data
	}

	assert.Equal(t, len(res), 2)
	assert.Equal(t, sum, 25)
}
