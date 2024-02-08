package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import "fmt"

// SubsystemA представляет подсистему A
type SubsystemA struct {
}

func (a *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation A")
}

// SubsystemB представляет подсистему B
type SubsystemB struct {
}

func (b *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation B")
}

// Facade предоставляет унифицированный интерфейс к подсистеме
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

// Operation осуществляет сложную операцию, используя подсистему
func (f *Facade) Operation() {
	fmt.Println("Facade: Operation")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func main() {
	// Использование фасада для упрощения взаимодействия с подсистемой
	facade := NewFacade()
	facade.Operation()
}
