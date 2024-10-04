// Некоторые примеры работы с пакетом sync
package syncpackage

import "sync"

type Donation struct {
	cond    *sync.Cond
	balance int
}

var donation = &Donation{
	cond: sync.NewCond(&sync.Mutex{}),
}

// ListenerDonationUpdate функция показывающая примеры работы с sync.Cond
// Среди примитивов синхронизации в пакете sync, sync.Cond, вероятно наименее используется и понятен.
// Однако он предоставляет функции, которых мы не можем достичь с помощью каналов. sync.Cond реализует
// переменную условия, точку встречи для подпрограмм, ожидающих или объявляющих о наступлении события.
// Переменная условия - это контейнер потоков (здесь - подпрограмм), ожидающих выполнения определенного
// условия. В нашем примере условием является обновление баланса. Программа обновления
// отправляет уведомление всякий раз, когда обновляется баланс, а программа прослушивания
// ожидает обновления. Кроме того, sync.Cond полагается на синхронизацию.Locker (a *sync.Mutex
// или *sync.RWMutex) для предотвращения гонки данных.
func ListenerDonationUpdate(goal int, ch chan int) {
	donation.cond.L.Lock()
	for donation.balance < goal {
		donation.cond.Wait()
	}

	//fmt.Printf("%d$ goal reached\n", donation.balance)
	ch <- donation.balance
	donation.cond.L.Unlock()
}
