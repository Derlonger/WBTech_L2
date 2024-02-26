package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Strategy pattern (Стратегия) - это поведенческий шаблон проектирования, который предназначен для определения семейства
алгоритмов, инкапсуляции каждого из них и обеспечения их взаимозаменяемости.

Плюсы:
	- Инкапсуляция реализации различных алгоритмов.
	- Система становится независимой от возможных изменений бизнес-правил.
	- Вызов всех алгоритмов одним стандартным образом.
	- Отказ от использования переключателей и/или условных операторов.
Минусы:
	- Создание дополнительных классов.
*/

import "fmt"

// Интерфейс для стратегий оплаты
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Конкретная стратегия: оплата через банковскую карту
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено %.2f с помощью банковской карты", amount)
}

// Конкретная стратегия: оплата через PayPal
type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено %.2f через PayPal", amount)
}

// Контекст (клиентский код)
type PaymentContext struct {
	strategy PaymentStrategy
}

func (ctx *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	ctx.strategy = strategy
}

func (ctx *PaymentContext) MakePayment(amount float64) string {
	return ctx.strategy.Pay(amount)
}

func main() {
	context := &PaymentContext{}

	// Выбираем стратегию оплаты
	context.SetStrategy(&CreditCardPayment{})
	fmt.Println(context.MakePayment(100.50))

	context.SetStrategy(&PayPalPayment{})
	fmt.Println(context.MakePayment(50.25))
}
