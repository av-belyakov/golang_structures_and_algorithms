package structuraldesignpatterns

import "fmt"

//	Шаблон "Составной" (Composite)
//
//Составной объект - это группа похожих объектов в одном объекте. Объекты хранятся в виде ДЕРЕВА
//для сохранения всей иерархии. Составной шаблон используется для изменения иерархической
//коллекции объектов. Составной шаблон моделируется на основе гетерогенной коллекции. Новые
//типы объектов могут быть добавлены без изменения интерфейса и клиентского кода. Вы можете
//использовать составной шаблон, например, для макетов пользовательского интерфейса в Интернете, для деревьев каталогов и
//для управления сотрудниками в разных отделах. Шаблон предоставляет механизм для доступа к
//отдельным объектам и группам аналогичным образом.

// IComposite interface
type IComposite interface {
	perform()
}

// Leaflet struct (лист ветви)
// name - наименование листа
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

// Branch struct (ветвь)
// leafs - листья
// name - наименование ветви
// branches - ветви
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// Branch class method perform
func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}
	for _, branch := range branch.branches {
		branch.perform()
	}
}

// Branch class method add leaflet
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// Branch class method addBranch branch
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// Branch class method getLeaflets
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

// CompositeExample method
func CompositeExample() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}
	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()
}
