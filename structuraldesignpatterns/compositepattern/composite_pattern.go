// Шаблон "Составной" (Composite)
package structuraldesignpatterns

//Составной объект - это группа похожих объектов в одном объекте. Объекты хранятся в виде ДЕРЕВА
//для сохранения всей иерархии. Составной шаблон используется для изменения иерархической
//коллекции объектов. Составной шаблон моделируется на основе гетерогенной коллекции. Новые
//типы объектов могут быть добавлены без изменения интерфейса и клиентского кода. Вы можете
//использовать составной шаблон, например, для макетов пользовательского интерфейса в Интернете, для деревьев каталогов и
//для управления сотрудниками в разных отделах. Шаблон предоставляет механизм для доступа к
//отдельным объектам и группам аналогичным образом.

// IComposite interface
type IComposite interface {
	getName() string
}

// Leaflet struct (лист ветви)
// name - наименование листа
type Leaflet struct {
	name string
}

// Leaflet class method perform
func (leaf *Leaflet) getName() string {
	return "Leaflet " + leaf.name
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
func (branch *Branch) getName() string {
	return "Branch: " + branch.name
}

func (branch *Branch) perform() []string {
	result := []string{}
	result = append(result, branch.getName())

	for _, leaf := range branch.leafs {
		result = append(result, leaf.getName())
	}

	for _, branch := range branch.branches {
		result = append(result, branch.perform()...)
	}

	return result
}

// Branch class method add leaflet
func (branch *Branch) addLeaf(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// Branch class method addBranch branch
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// CompositeExample method
func CompositeExample() {
	var branch = &Branch{name: "branch 1"}
	var leaf1 = Leaflet{name: "leaf 1"}
	var leaf2 = Leaflet{name: "leaf 2"}
	var branch2 = Branch{name: "branch 2"}
	branch.addLeaf(leaf1)
	branch.addLeaf(leaf2)
	branch.addBranch(branch2)
	branch.perform()
}
