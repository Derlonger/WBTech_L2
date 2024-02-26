package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Factory Method(Фабричный метод) - это порождающий паттерн проектирования, который решает проблему создания различных продуктов
без указания конкретных классов продуктов. Он создает метод, который следует использовать вместо вызова оператора new
для создания объектов продуктов.

Плюс:

	1. Расширяемость и гибкость:
		- Паттерн Фабричный метод позволяет легко добавлять новые продукты (классы), не изменяя существующий код.
		- Это дает возможность гибко справляться с изменениями требований и добавлением нового функционала в систему.
	2. Отделение клиентского кода от конкретных классов:
		- Клиенту не нужно уделять значительное внимание конфигурированию экземпляра конкретного
		создателя (ConcreteCreator) специфическим продуктом.
		- Каждый продукт (ConcreteProduct) может конфигурироваться параметрами, о которых знает только
		конкретный создатель (ConcreteCreator)2.

Минусы:

	1. Дополнительная абстракция:
		- Внедрение фабричного метода может добавить дополнительный уровень абстракции,
		что может быть избыточным для простых систем.
		- В некоторых случаях прямое создание объектов без фабрик может быть более простым и понятным.
	2. Сложность для небольших систем:
		- В небольших системах, где нет сложной иерархии продуктов, использование фабричного метода может показаться излишним.
		- В таких случаях простое создание объектов напрямую может быть более эффективным.

*/

// Product - Интерфейс для продуктов
type Product interface {
	Use()
}

// ConcreteProductA - Конкретный продукт A
type ConcreteProductA struct{}

func (c *ConcreteProductA) Use() {
	fmt.Println("Using Concrete Product A")
}

// ConcreteProductB - Конкретный продукт B
type ConcreteProductB struct{}

func (c *ConcreteProductB) Use() {
	fmt.Println("Using Concrete Product B")
}

// Creator - интерфейс для фабрики
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA - Конкретная фабрика A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB - Конкретная фабрика B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Создаем фабрику A
	creatorA := ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	productA.Use()

	// Создаем фабрику B
	creatorB := ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	productB.Use()
}
