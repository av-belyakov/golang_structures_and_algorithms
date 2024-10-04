// Примеры работы с пакетом time
package timepackage

import "time"

// TimeEvent тип хранящий некое время
type TimeEvent struct {
	time.Time
}

// IsEqual функция выполняющая сравнение времени
func (t *TimeEvent) IsEqual(myTime time.Time) bool {
	return t.Time.Equal(myTime)
}

// NewTimeWithTruncate функция формирующая новый пользовательский тип с использоанием метода Truncate
func NewTimeWithTruncate(myTime time.Time) TimeEvent {
	return TimeEvent{Time: myTime.Truncate(0)}
}
