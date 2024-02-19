package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад (Facade) - структурный паттерн.
Представляет собой простой доступ (простой интерфейс) к сложной системе.

Используем его, когда у нас есть много разных подсистем, которые используют свои интерфейсы
и реализуют какой-то свой функционал поведения.

Реализует принципы SOLID:
- Открытости/закрытости (интерфейс в который скрывает код подсистемы)
- Инверсия зависимостей (уменьшает связанность)

Преимущества:
- Изолирует клиентов от поведения сложной системы
- Сам интерфейс фасада простой

Минусы:
- Является супер-объектом и все последующие вызовы в системе будут проходить
через этот объект
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
