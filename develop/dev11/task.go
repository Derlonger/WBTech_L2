package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// Event представляет собой структуру для событий в календаре.
type Event struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
}

// Сериализация объектов в Json
func toJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

var events []Event

func createEventHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		var event Event
		if err := json.NewDecoder(request.Body).Decode(&event); err != nil {
			fmt.Println(err)
			return
		}

		for i := range events {
			if events[i].ID == event.ID {
				writer.Write([]byte(fmt.Sprintf("ID %d alredy exists!", event.ID)))
				return
			}
		}

		events = append(events, event)
		fmt.Println(events)
		writer.Write([]byte("Row success created"))
	default:
		writer.Write([]byte("Only POST HTTP Method"))
		return
	}
}

func updateEventHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		var event Event
		if err := json.NewDecoder(request.Body).Decode(&event); err != nil {
			fmt.Println(err)
			return
		}

		for i := range events {
			if events[i].ID == event.ID {
				events[i] = event
				writer.Write([]byte("Row success update"))
				return
			}
		}

	default:
		writer.Write([]byte("Only PUT HTTP Method"))
		return

	}
}

func deleteEventHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "DELETE":
		id := request.URL.Query().Get("id")

		for i := range events {
			if strconv.Itoa(events[i].ID) == id {
				events = append(events[:i], events[i+1:]...)
				writer.Write([]byte("Row success delete"))
				return
			}

		}
	default:
		writer.Write([]byte("Only DELETE HTTP Method"))
		return
	}
}

func eventsForDayHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		events := getEventsForLastDay()

		// Сериализуем события в формат JSON
		response, err := toJSON(events)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("Error: %v", err)))
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		writer.Write(response)
	default:
		writer.Write([]byte("Only GET HTTP Method"))
		return
	}
}

func getEventsForLastDay() []Event {
	var eventsForLastDay []Event

	// Получение текущей даты и времени
	now := time.Now()

	// Определение временных границ для последнего дня
	dayStart := now.AddDate(0, 0, -1).Truncate(24 * time.Hour)
	dayEnd := now

	// Перебор событий и выбор тех, которые в пределах последнего дня
	for _, event := range events {
		if event.Date.After(dayStart) && event.Date.Before(dayEnd) {
			eventsForLastDay = append(eventsForLastDay, event)
		}
	}

	return eventsForLastDay
}

func eventsForWeekHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		events := getEventsForLastWeek()

		// Сериализуем события в формат JSON
		response, err := toJSON(events)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("Error: %v", err)))
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		writer.Write(response)
	default:
		writer.Write([]byte("Only GET HTTP Method"))
		return
	}
}

func getEventsForLastWeek() []Event {
	var eventsForLastWeek []Event

	// Получение текущей даты и времени
	now := time.Now()

	// Определение временных границ для последней недели
	weekStart := now.AddDate(0, 0, -7).Truncate(24 * time.Hour)
	weekEnd := now

	// Перебор событий и выбор тех, которые в пределах последней недели
	for _, event := range events {
		if event.Date.After(weekStart) && event.Date.Before(weekEnd) {
			eventsForLastWeek = append(eventsForLastWeek, event)
		}
	}

	return eventsForLastWeek
}

func eventsForMonthHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		events := getEventsForLastMonth()

		// Сериализуем события в формат JSON
		response, err := toJSON(events)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("Error: %v", err)))
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		writer.Write(response)
	default:
		writer.Write([]byte("Only GET HTTP Method"))
		return
	}
}

func getEventsForLastMonth() []Event {
	var eventsForLastMonth []Event

	// Получение текущей даты и времени
	now := time.Now()

	// Определение временных границ для последнего месяца
	monthStart := now.AddDate(0, -1, 0).Truncate(24 * time.Hour)
	monthEnd := now

	// Перебор событий и выбор тех, которые в пределах последнего месяца
	for _, event := range events {
		if event.Date.After(monthStart) && event.Date.Before(monthEnd) {
			eventsForLastMonth = append(eventsForLastMonth, event)
		}
	}

	return eventsForLastMonth
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		f(writer, request)
	}
}

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/create_event", logging(createEventHandler))
	http.HandleFunc("/update_event", logging(updateEventHandler))
	http.HandleFunc("/delete_event", logging(deleteEventHandler))
	http.HandleFunc("/events_for_day", logging(eventsForDayHandler))
	http.HandleFunc("/events_for_week", logging(eventsForWeekHandler))
	http.HandleFunc("/events_for_month", logging(eventsForMonthHandler))

	log.Printf("Server is running..")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
