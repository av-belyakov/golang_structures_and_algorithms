package syncpackage

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//Среди примитивов синхронизации в пакете sync, sync.Cond, вероятно наименее используется и понятен.
//Однако он предоставляет функции, которых мы не можем достичь с помощью каналов. sync.Cond реализует
//переменную условия, точку встречи для подпрограмм, ожидающих или объявляющих о наступлении события.
//Переменная условия - это контейнер потоков (здесь - подпрограмм), ожидающих выполнения определенного
//условия. В нашем примере условием является обновление баланса. Программа обновления
//отправляет уведомление всякий раз, когда обновляется баланс, а программа прослушивания
//ожидает обновления. Кроме того, sync.Cond полагается на синхронизацию.Locker (a *sync.Mutex
//или *sync.RWMutex) для предотвращения гонки данных.

type Donation struct {
	cond    *sync.Cond
	balance int
}

var donation = &Donation{
	cond: sync.NewCond(&sync.Mutex{}),
}

// Listener goroutines
func ListenerDonationUpdate(goal int, ch chan int) {
	donation.cond.L.Lock()
	for donation.balance < goal {
		donation.cond.Wait()
	}

	//fmt.Printf("%d$ goal reached\n", donation.balance)
	ch <- donation.balance
	donation.cond.L.Unlock()
}

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
