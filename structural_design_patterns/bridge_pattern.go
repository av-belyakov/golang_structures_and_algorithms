package structuraldesignpatterns

import "fmt"

// Шаблон "Мост" (Bridge)
//
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

// IDrawShape interface
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// DrawShape struct
type DrawShape struct{}

// Метод DrawShape рисует фигуру с заданными координатами
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

// IContour interface
type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

// DrawContour struct
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

// DrawContour метод drawContour возвращает координаты
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

// DrawContour метод resizeByFactor возвращает множитель factor
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

// BridgeExample method
func BridgeExample() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}

//Интерфейс IContour позволяет использовать метод drawContour пользовательского
//типа DrawShape через DrawContour
