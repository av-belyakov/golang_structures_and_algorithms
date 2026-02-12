package supportfunctions

import "sync"

// ElementsFromJSON элементы полученные при обработки JSON объекта
type ElementsFromJSON struct {
	mutex sync.RWMutex
	Data  map[string]Element
}

// CommonValues общие, для типов, значения
type CommonValues struct {
	Value     any    //любые передаваемые данные
	FieldName string //наименование поля
	ValueType string //тип передаваемого значения (string, int и т.д.)
}

// Element описание объекта
type Element struct {
	CommonValues
}

type chResult struct {
	CommonValues
	FieldBranch string //'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
}
