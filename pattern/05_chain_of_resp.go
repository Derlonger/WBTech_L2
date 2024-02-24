package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Chain of responsibility(Цепочка вызовов) - это поведенческий паттерн проектирования, который позволяет передавать задачи
по цепочке объектов. Каждый объект в цепочке может обработать запрос или передать его дальше.

Плюсы:

	1. Разделение обязанностей:
		- Позволяет разделить обязанности между объектами, каждый из которых выполняет свою часть работы.
	2. Гибкость и расширяемость:
		- Мы можем легко добавлять новые обработчики в цепочку без изменения существующего кода.
	3. Избегание жестких зависимостей:
		- Объекты в цепочке не зависят друг от друга напрямую, что делает систему более гибкой.

Минусы:
	1. Нет гарантии обработки:
		- Если ни один из обработчиков не обработал запрос, он может остаться без ответа. Необходимо обеспечить,
		чтобы в цепочке всегда был обработчик, способный обработать запрос.
	2. Потенциальная перегрузка:
		- Если цепочка слишком длинная или обработчики выполняют сложные операции,
		это может привести к нежелательной нагрузке на систему.
	3. Сложность отладки:
		- При наличии множества обработчиков может быть сложно определить, где именно возникла ошибка.

*/

// OrderHandler - Интерфейс для обработки заказов
type OrderHandler interface {
	SetNext(handler OrderHandler)
	Handle(order Order)
}

// Order - структура заказа
type Order struct {
	ID       int
	Products []string
}

// ProductAvailabilityHandler - обработчик проверки наличия товара
type ProductAvailabilityHandler struct {
	next OrderHandler
}

func (h *ProductAvailabilityHandler) SetNext(handler OrderHandler) {
	h.next = handler
}

func (h *ProductAvailabilityHandler) Handle(order Order) {
	// Проверка наличия товара
	fmt.Println("Проверка наличия товара для заказа", order.ID)
	// Если товар есть, передаем обработку следующему обработчику
	if h.next != nil {
		h.next.Handle(order)
	}
}

// CostCalculationHandler - обработчик расчета стоимости
type CostCalculationHandler struct {
	next OrderHandler
}

func (h *CostCalculationHandler) SetNext(handler OrderHandler) {
	h.next = handler
}

func (h *CostCalculationHandler) Handle(order Order) {
	// Рассчитываем стоимость заказа
	fmt.Println("Расчет стоимости заказа", order.ID)
	// Передаем обработку следующему обработчику
	if h.next != nil {
		h.next.Handle(order)
	}
}

// OrderConfirmationHandler - обработчик оформления заказа
type OrderConfirmationHandler struct {
	next OrderHandler
}

func (h *OrderConfirmationHandler) SetNext(handler OrderHandler) {
	h.next = handler
}

func (h *OrderConfirmationHandler) Handle(order Order) {
	// Оформляем заказ
	fmt.Println("Оформление заказа", order.ID)
	// Передаем обработку следующему обработчику
	if h.next != nil {
		h.next.Handle(order)
	}
}

// NotificationHandler - обработчик отправки уведомления клиенту
type NotificationHandler struct {
	next OrderHandler
}

func (h *NotificationHandler) SetNext(handler OrderHandler) {
	h.next = handler
}

func (h *NotificationHandler) Handle(order Order) {
	// Отправляем уведомление клиенту
	fmt.Println("Отправка уведомления клиенту о заказе", order.ID)
}

func main() {
	order := Order{
		ID:       123,
		Products: []string{"item1", "item2"},
	}

	// Создаем цепочку обработчиков
	productAvailabilityHandler := &ProductAvailabilityHandler{} // Проверка наличия товара
	costCalculationHandler := &CostCalculationHandler{}         // расчет стоимости
	orderConfirmationHandler := &OrderConfirmationHandler{}     // оформление заказа
	notificationHandler := &NotificationHandler{}               // Отправка уведомления пользователю

	productAvailabilityHandler.SetNext(costCalculationHandler)
	costCalculationHandler.SetNext(orderConfirmationHandler)
	orderConfirmationHandler.SetNext(notificationHandler)

	// Запускаем обработку заказа
	productAvailabilityHandler.Handle(order)
}
