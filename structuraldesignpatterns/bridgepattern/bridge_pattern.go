// Шаблон "Мост" (Bridge)
package structuraldesignpatterns

import (
	"fmt"
	"time"
)

//Мост отделяет реализацию от абстракции. Абстрактный базовый класс может быть
//разделен на подклассы, чтобы предоставлять различные реализации и позволять
//легко изменять детали реализации. Интерфейс, который является мостом, помогает
//сделать функциональность конкретных классов независимой от классов-реализаторов интерфейса.
//Шаблоны моста позволяют изменять детали реализации во время выполнения.
//
//Шаблон bridge демонстрирует принцип, отдавая предпочтение композиции перед наследованием. Это
//помогает в ситуациях, когда нужно многократно создавать подклассы, ортогональные друг другу.
//Привязка приложения к среде выполнения, отображение ортогональных иерархий классов и
//реализация независимости от платформы - вот сценарии, в которых может быть применен шаблон моста.

type IConverterDate interface {
	ConverterDate(strTime string) (time.Time, error)
}

type ConvertDate struct{}

func (cd ConvertDate) ConverterDate(strTime string) (time.Time, error) {
	return time.Parse(time.RFC3339, strTime)
}

type IShowerData interface {
	ConverterDate(strTime string) (time.Time, error)
	GetData() (int, time.Month, int)
	GetTime() (int, int, int)
	SetData(time.Time)
}

type ShowData struct {
	CurrentTime time.Time
	Convert     ConvertDate
}

func (sd ShowData) ConverterDate(strTime string) (time.Time, error) {
	return sd.Convert.ConverterDate(strTime)
}

func (sd *ShowData) SetData(t time.Time) {
	sd.CurrentTime = t
}

func (sd ShowData) GetData() (int, time.Month, int) {
	return sd.CurrentTime.Date()
}

func (sd ShowData) GetTime() (int, int, int) {
	hour := sd.CurrentTime.Hour()
	minute := sd.CurrentTime.Minute()
	second := sd.CurrentTime.Second()

	return hour, minute, second
}

// BridgeExample method
func BridgeExample() {
	var showData IShowerData = &ShowData{Convert: ConvertDate{}}
	t, _ := showData.ConverterDate("2023-01-02T15:04:05+07:00")
	showData.SetData(t)
	year, month, day := showData.GetData()
	h, m, s := showData.GetTime()
	fmt.Printf("Data: %d %s %d", year, month.String(), day)
	fmt.Printf("Time: %d:%d:%d", h, m, s)
}

//Интерфейс IContour позволяет использовать метод drawContour пользовательского
//типа DrawShape через DrawContour
