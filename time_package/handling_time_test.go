package timepackage

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//Локальные часы используются для определения текущего времени суток. Эти часы могут
//отличаться. Например, если часы синхронизированы по NTP, они могут перемещаться назад
// или вперед по времени когда NTP сервер выполняет корректировку. По этому не стоит
//измерять длительность с помощью локальных часов. Может возникнуть странное поведение,
//например отрицательная длительность.
//Вот почему в OS есть второй тип часов: монотонные часы. Монотонные часы гарантируют,
//что время всегда движется вперед и на него не влияют скачки во времени. На это
//могут повлиять настройки частоты (например, если сервер обнаружит, что
//локальные кварцевые часы движутся с другой скоростью, чем NTP-сервер), но никогда
//скачками во времени.
//Монотонные часы не сохраняются например при преобразовании к формату JSON

func TestEqualTime(t *testing.T) {
	type Event struct {
		time.Time
	}

	var (
		event1 Event
		event2 Event
	)

	timeNow := time.Now()

	event1 = Event{Time: timeNow}
	b, err := json.Marshal(event1)
	assert.NoError(t, err)

	event2 = Event{}
	err = json.Unmarshal(b, &event2)
	assert.NoError(t, err)

	//несмотра на то что для сравнения использовалось одно и тоже время,
	//результат не равен. Так как до преобразования в JSON переменная
	//event1 в анонимном поле time содержала локальное время + монотонное время в виде:
	//2021-01-10 17:13:08.852061 +0100 CET m=+0.000338660, где m=+0.000338660
	//это монотонное время, а после преобразования из JSON переменная event2
	//содержит только локальное время 2021-01-10 17:13:08.852061 +0100 CET
	isNotEqual := event1.Time == event2.Time
	assert.False(t, isNotEqual)

	//этого можно избежать двумя способами, первый это использовать для сравнения
	//специальный метод time.Time.Equal
	isEqual := event1.Time.Equal(event2.Time)
	assert.True(t, isEqual)

	//второй вариант, при формировании пользовательского типа использовать time.Time.Truncate
	event1 = Event{Time: timeNow.Truncate(0)}
	b, err = json.Marshal(event1)
	assert.NoError(t, err)

	event2 = Event{}
	err = json.Unmarshal(b, &event2)
	assert.NoError(t, err)

	isEqual = event1.Time == event2.Time
	assert.True(t, isEqual)
}

//Кроме того можно использовать time.LoadLocation для задания местоположения. Например:
//location, err := time.LoadLocation("Russia/Moscow")
//if err != nil {
// return err
//}
//t := time.Now().In(location) // 2021-05-18 22:47:04.155755 -0500 EST
//Как альтернативу мы можем использовать UTC формат
//t := time.Now().UTC() // 2021-05-18 22:47:04.155755 +0000 UTC
