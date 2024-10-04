// Структура данных типа "Множество" или "Набор" (Set)
package lineardatastructures

import "fmt"

//Множество (Набор) - это линейная структура данных, содержащая набор значений, которые не повторяются. Набор может
//хранить уникальные значения без какого-либо определенного порядка. В реальном мире наборы можно использовать для
//сбора всех тегов для записей в блоге и участников беседы в чате. Данные могут быть логического, целочисленного,
//плавающего, символьного и других типов. Статические наборы допускают только операции запроса, что означает
//операции, связанные с запросом элементов. Динамические и изменяемые наборы допускают вставку и удаление элементов.
//Алгебраические операции, такие как объединение, пересечение, разница и подмножество могут быть определены в наборах.

// Set class
type Set struct {
	integerMap map[int]bool
}

// New создает новое множество
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// AddElement добавляет элемент во множество
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// DeleteElement удаляет элемент из множества
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// ContainsElement проверяет наличие элемента во множестве
func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]

	return exists
}

// InterSect метод возвращает набор пересечений, состоящий из пересечения set и другого множества (набора)
func (set *Set) InterSect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	for value := range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}

	return intersectSet
}

// Union метод выполняющий обединение нескольких множеств (наборов)
func (set *Set) Union(anotherSet *Set) *Set {
	var (
		unionSet = &Set{}
		value    int
	)

	unionSet.New()
	for value = range set.integerMap {
		unionSet.AddElement(value)
	}

	for value = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}

	return unionSet
}

func SetsExample() {
	var set *Set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
}
