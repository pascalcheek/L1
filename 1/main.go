package main

import "fmt"

type Human struct {
	name string
	age  int
}

// Вызов этого метода доступен как от типа Human, так и от типа Action
func (h Human) PrintName() {
	fmt.Printf(h.name)
}

type Action struct {
	Human  // Встраивание структуры Human в структуру Action.
	action string
}

func (a Action) DoAction() {
	fmt.Printf("%s, %s!\n", a.name, a.action) // Тут демонстрируется использование наследованного поля Human
}

func (a Action) AnotherDoAction() {
	a.PrintName() // Тут демонстрируется наследование метода
	fmt.Printf(", %s!\n", a.action)
}

func main() {
	a := Action{
		Human:  Human{"Pascal", 20},
		action: "playing football",
	}
	a.PrintName()
	fmt.Println()
	a.DoAction()
	a.AnotherDoAction()
}
