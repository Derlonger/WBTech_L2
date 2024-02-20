package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Builder - это порождающий паттерн проектирования, который позволяет создавать сложные объекты поэтапно.
Так же имеется Директор (или Director) — это дополнительный компонент, связанный с паттерном Builder.

Назначение директора:
Директор — это структура, который управляет процессом пошагового создания сложного объекта с помощью Builder.
Его задача — вызывать методы Builder в правильной последовательности, чтобы построить объект.
Директор отделяет клиентский код от деталей создания объекта.


Плюс:
	1. Разделение сложного процесса создание объекта. Builder разбивает процесс создания на несколько шагов,
что упрощает работу с большим кол-ом параметров
	2. Гибкость и расширяемость. Мы можем добавлять новые параметры и методы в Builder без изменения сущ. кода.
	3. Читаемость и поддерживаемость. Код, использующий Builder, более читаем и понятен.
Мы можем легко настроить объект, добавив или изменив параметры.

Минусы:
	1. Дополнительный код: Нам нужно создать отдельный класс Builder, что может увеличить количество кода.
	2. Сложность для простых объектов: Для простых объектов Builder может показаться избыточным.
Если у нас всего несколько параметров, то использование Builder может быть излишним.
	3. Возможность ошибок: При создании объекта через Builder мы все равно должны правильно вызвать методы в правильной последовательности.
Ошибки могут возникнуть, если мы забудем вызвать какой-то метод или вызовем его не в той последовательности.
*/

// Структура персонажа компьютерной игры
type Character struct {
	Name     string
	Level    int
	Strength int
	Health   int
}

// Интерфейс для определения методов для настройки персонажа
type CharacterBuilder interface {
	SetName(name string)
	SetLevel(level int)
	SetStrength(strength int)
	SetHealth(health int)
	Build() *Character
}

// Структура для реализации интерфейса CharacterBuilder
type SimpleCharacterBuilder struct {
	character *Character
}

func NewSimpleCharacterBuilder() *SimpleCharacterBuilder {
	return &SimpleCharacterBuilder{
		character: &Character{},
	}
}

func (b *SimpleCharacterBuilder) SetName(name string) {
	b.character.Name = name
}

func (b *SimpleCharacterBuilder) SetLevel(level int) {
	b.character.Level = level
}

func (b *SimpleCharacterBuilder) SetStrength(strength int) {
	b.character.Strength = strength
}

func (b *SimpleCharacterBuilder) SetHealth(health int) {
	b.character.Health = health
}

func (b *SimpleCharacterBuilder) Build() *Character {
	return b.character
}

// Директор для создания персонажа
type CharacterDirector struct {
	builder CharacterBuilder
}

func NewCharacterDirector(builder CharacterBuilder) *CharacterDirector {
	return &CharacterDirector{builder: builder}
}

func (d *CharacterDirector) CreateHero() *Character {
	d.builder.SetName("Герой")
	d.builder.SetLevel(1)
	d.builder.SetStrength(10)
	d.builder.SetHealth(100)
	return d.builder.Build()
}

func main() {
	builder := NewSimpleCharacterBuilder()
	director := NewCharacterDirector(builder)

	character := director.CreateHero()

	// Теперь у нас есть готовый персонаж!
	fmt.Printf("Имя: %s, Уровень: %d, Сила: %d, Здоровье: %d\n",
		character.Name, character.Level, character.Strength, character.Health)
}
