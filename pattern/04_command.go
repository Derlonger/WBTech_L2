package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн Command позволяет инкапсулировать запрос на выполнение определенного действия в виде отдельного объекта.
Этот объект запроса на действие и называется 'командой'

Плюсы:

	1. Инкапсуляция действий:
		- Команда позволяет инкапсулировать запрос на выполнение действия в виде отдельного объекта.
		- Это упрощает добавление новых операций без изменения самих объектов.
	2. Отделение инициатора и получателя:
		- Объекты, инициирующие запросы на выполнение действия (инициаторы), отделяются от объектов,
	которые выполняют это действие (получатели).
	3. Поддержка отмены операций:
		-Команды могут поддерживать отмену выполненных действий (например, через метод Undo()).

Минусы:

	1.Сложность:
		- Внедрение паттерна может усложнить код, особенно если у нас большое количество команд и получателей.
	2. Нарушение инкапсуляции:
		- Получатель команды (например, класс Receiver) получает доступ к приватным членам объектов, что может нарушить инкапсуляцию.
	3. Увеличение числа классов:
		- Для каждой операции требуется создание отдельной команды, что может привести к увеличению числа классов в системе.

*/

// Интерфейс команды
type Command interface {
	Execute() // Выполнение команды
	Undo()    // Отмена команды
}

// Конкретная команда для добавления книги
type AddBookCommand struct {
	receiver *BookReceiver
	book     Book
}

func (c *AddBookCommand) Execute() {
	c.receiver.AddBook(c.book)
}

func (c *AddBookCommand) Undo() {
	c.receiver.RemoveBook(c.book)
}

// Получатель команды
type BookReceiver struct {
	books []Book
}

func (r *BookReceiver) AddBook(book Book) {
	r.books = append(r.books, book)
	fmt.Println("Книга добавлена:", book.Title)
}

func (r *BookReceiver) RemoveBook(book Book) {
	index := -1
	for i, b := range r.books {
		if b.Title == book.Title {
			index = i
			break
		}
	}

	if index != -1 {
		r.books = append(r.books[:index], r.books[index+1:]...)
		fmt.Println("Книга удалена:", book.Title)
	} else {
		fmt.Println("Книга не найдена:", book.Title)
	}
}

// Структура книги
type Book struct {
	Title string
}

// Инициатор команды
type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) Run() {
	i.command.Execute()
}

func main() {
	receiver := &BookReceiver{}
	addCommand := &AddBookCommand{
		receiver: receiver,
		book:     Book{Title: "Гарри Поттер и философский камень"},
	}

	// Инициатор команды
	invoker := &Invoker{}
	invoker.SetCommand(addCommand)
	invoker.Run()
}
